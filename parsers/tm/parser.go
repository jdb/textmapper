// generated by Textmapper; DO NOT EDIT

package tm

import (
	"context"
	"fmt"

	"github.com/inspirer/textmapper/parsers/tm/token"
)

// ErrorHandler is called every time a parser is unable to process some part of the input.
// This handler can return false to abort the parser.
type ErrorHandler func(err SyntaxError) bool

// StopOnFirstError is an error handler that forces the parser to stop on and return the first
// error.
func StopOnFirstError(_ SyntaxError) bool { return false }

// Parser is a table-driven LALR parser for tm.
type Parser struct {
	eh       ErrorHandler
	listener Listener

	next       symbol
	endState   int16
	recovering int

	// Tokens to be reported with the next shift. Only non-empty when next.symbol != noToken.
	pending []symbol
}

type SyntaxError struct {
	Line      int
	Offset    int
	Endoffset int
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf("syntax error at line %v", e.Line)
}

type symbol struct {
	symbol    int32
	offset    int
	endoffset int
}

type stackEntry struct {
	sym   symbol
	state int16
}

func (p *Parser) Init(eh ErrorHandler, l Listener) {
	p.eh = eh
	p.listener = l
	if cap(p.pending) < startTokenBufferSize {
		p.pending = make([]symbol, 0, startTokenBufferSize)
	}
}

const (
	startStackSize       = 256
	startTokenBufferSize = 16
	noToken              = int32(token.UNAVAILABLE)
	eoiToken             = int32(token.EOI)
	debugSyntax          = false
)

func (p *Parser) ParseFile(ctx context.Context, lexer *Lexer) error {
	return p.parse(ctx, 0, 509, lexer)
}

func (p *Parser) ParseNonterm(ctx context.Context, lexer *Lexer) error {
	return p.parse(ctx, 1, 510, lexer)
}

func (p *Parser) parse(ctx context.Context, start, end int16, lexer *Lexer) error {
	p.pending = p.pending[:0]
	var shiftCounter int
	state := start
	var lastErr SyntaxError
	p.recovering = 0

	var alloc [startStackSize]stackEntry
	stack := append(alloc[:0], stackEntry{state: state})
	p.endState = end
	p.fetchNext(lexer, stack)

	for state != end {
		action := tmAction[state]
		if action < -2 {
			// Lookahead is needed.
			if p.next.symbol == noToken {
				p.fetchNext(lexer, stack)
			}
			action = lalr(action, p.next.symbol)
		}

		if action >= 0 {
			// Reduce.
			rule := action
			ln := int(tmRuleLen[rule])

			var entry stackEntry
			entry.sym.symbol = tmRuleSymbol[rule]
			rhs := stack[len(stack)-ln:]
			stack = stack[:len(stack)-ln]
			if ln == 0 {
				if p.next.symbol == noToken {
					p.fetchNext(lexer, stack)
				}
				entry.sym.offset, entry.sym.endoffset = p.next.offset, p.next.offset
			} else {
				entry.sym.offset = rhs[0].sym.offset
				entry.sym.endoffset = rhs[ln-1].sym.endoffset
			}
			if err := p.applyRule(ctx, rule, &entry, rhs, lexer); err != nil {
				return err
			}
			if debugSyntax {
				fmt.Printf("reduced to: %v\n", symbolName(entry.sym.symbol))
			}
			state = gotoState(stack[len(stack)-1].state, entry.sym.symbol)
			entry.state = state
			stack = append(stack, entry)

		} else if action == -1 {
			if shiftCounter++; shiftCounter&0x1ff == 0 {
				// Note: checking for context cancellation is expensive so we do it from time to time.
				select {
				case <-ctx.Done():
					return ctx.Err()
				default:
				}
			}

			// Shift.
			if p.next.symbol == noToken {
				p.fetchNext(lexer, stack)
			}
			state = gotoState(state, p.next.symbol)
			if state >= 0 {
				stack = append(stack, stackEntry{
					sym:   p.next,
					state: state,
				})
				if debugSyntax {
					fmt.Printf("shift: %v (%s)\n", symbolName(p.next.symbol), lexer.Text())
				}
				if len(p.pending) > 0 {
					for _, tok := range p.pending {
						p.reportIgnoredToken(ctx, tok)
					}
					p.pending = p.pending[:0]
				}
				if p.next.symbol != eoiToken {
					p.next.symbol = noToken
				}
				if p.recovering > 0 {
					p.recovering--
				}
			}
		}

		if action == -2 || state == -1 {
			if p.recovering == 0 {
				if p.next.symbol == noToken {
					p.fetchNext(lexer, stack)
				}
				lastErr = SyntaxError{
					Line:      lexer.Line(),
					Offset:    p.next.offset,
					Endoffset: p.next.endoffset,
				}
				if !p.eh(lastErr) {
					if len(p.pending) > 0 {
						for _, tok := range p.pending {
							p.reportIgnoredToken(ctx, tok)
						}
						p.pending = p.pending[:0]
					}
					return lastErr
				}
			}

			p.recovering = 4
			if stack = p.recoverFromError(ctx, lexer, stack); stack == nil {
				if len(p.pending) > 0 {
					for _, tok := range p.pending {
						p.reportIgnoredToken(ctx, tok)
					}
					p.pending = p.pending[:0]
				}
				return lastErr
			}
			state = stack[len(stack)-1].state
		}
	}

	return nil
}

const errSymbol = 38

// willShift checks if "symbol" is going to be shifted in the given state.
// This function does not support empty productions and returns false if they occur before "symbol".
func (p *Parser) willShift(stackPos int, state int16, symbol int32, stack []stackEntry) bool {
	if state == -1 {
		return false
	}

	for state != p.endState {
		action := tmAction[state]
		if action < -2 {
			action = lalr(action, symbol)
		}

		if action >= 0 {
			// Reduce.
			rule := action
			ln := int(tmRuleLen[rule])
			if ln == 0 {
				// we do not support empty productions
				return false
			}
			stackPos -= ln - 1
			state = gotoState(stack[stackPos-1].state, tmRuleSymbol[rule])
		} else {
			return action == -1 && gotoState(state, symbol) >= 0
		}
	}
	return symbol == eoiToken
}

func (p *Parser) skipBrokenCode(ctx context.Context, lexer *Lexer, stack []stackEntry, canRecover func(symbol int32) bool) int {
	var e int
	for p.next.symbol != eoiToken && !canRecover(p.next.symbol) {
		if debugSyntax {
			fmt.Printf("skipped while recovering: %v (%s)\n", symbolName(p.next.symbol), lexer.Text())
		}
		if len(p.pending) > 0 {
			for _, tok := range p.pending {
				p.reportIgnoredToken(ctx, tok)
			}
			p.pending = p.pending[:0]
		}
		e = p.next.endoffset
		p.fetchNext(lexer, stack)
	}
	return e
}

func (p *Parser) recoverFromError(ctx context.Context, lexer *Lexer, stack []stackEntry) []stackEntry {
	var recoverSyms [1 + token.NumTokens/8]uint8
	var recoverPos []int

	if debugSyntax {
		fmt.Printf("broke at %v\n", symbolName(p.next.symbol))
	}
	for size := len(stack); size > 0; size-- {
		if gotoState(stack[size-1].state, errSymbol) == -1 {
			continue
		}
		recoverPos = append(recoverPos, size)
		if recoveryScopeStates[int(stack[size-1].state)] {
			break
		}
	}
	if len(recoverPos) == 0 {
		return nil
	}

	for _, v := range afterErr {
		recoverSyms[v/8] |= 1 << uint32(v%8)
	}
	canRecover := func(symbol int32) bool {
		return recoverSyms[symbol/8]&(1<<uint32(symbol%8)) != 0
	}
	if p.next.symbol == noToken {
		p.fetchNext(lexer, stack)
	}
	// By default, insert 'error' in front of the next token.
	s := p.next.offset
	e := s
	for _, tok := range p.pending {
		// Try to cover all nearby invalid tokens.
		if token.Token(tok.symbol) == token.INVALID_TOKEN {
			if s > tok.offset {
				s = tok.offset
			}
			e = tok.endoffset
		}
	}
	for {
		if endoffset := p.skipBrokenCode(ctx, lexer, stack, canRecover); endoffset > e {
			e = endoffset
		}

		var matchingPos int
		if debugSyntax {
			fmt.Printf("trying to recover on %v\n", symbolName(p.next.symbol))
		}
		for _, pos := range recoverPos {
			if p.willShift(pos, gotoState(stack[pos-1].state, errSymbol), p.next.symbol, stack) {
				matchingPos = pos
				break
			}
		}
		if matchingPos == 0 {
			if p.next.symbol == eoiToken {
				return nil
			}
			recoverSyms[p.next.symbol/8] &^= 1 << uint32(p.next.symbol%8)
			continue
		}

		if matchingPos < len(stack) {
			if s == e {
				// Avoid producing syntax problems covering trailing whitespace.
				e = stack[len(stack)-1].sym.endoffset
			}
			s = stack[matchingPos].sym.offset
		} else if s == e && len(p.pending) > 0 {
			// This means pending tokens don't contain InvalidTokens.
			for _, tok := range p.pending {
				p.reportIgnoredToken(ctx, tok)
			}
			p.pending = p.pending[:0]
		}
		if s != e {
			// Consume trailing invalid tokens.
			for _, tok := range p.pending {
				if token.Token(tok.symbol) == token.INVALID_TOKEN && tok.endoffset > e {
					e = tok.endoffset
				}
			}
			var consumed int
			for ; consumed < len(p.pending); consumed++ {
				tok := p.pending[consumed]
				if tok.offset >= e {
					break
				}
				p.reportIgnoredToken(ctx, tok)
			}
			newSize := len(p.pending) - consumed
			copy(p.pending[:newSize], p.pending[consumed:])
			p.pending = p.pending[:newSize]
		}
		if debugSyntax {
			for i := len(stack) - 1; i >= matchingPos; i-- {
				fmt.Printf("dropped from stack: %v\n", symbolName(stack[i].sym.symbol))
			}
			fmt.Println("recovered")
		}
		stack = append(stack[:matchingPos], stackEntry{
			sym:   symbol{errSymbol, s, e},
			state: gotoState(stack[matchingPos-1].state, errSymbol),
		})
		return stack
	}
}

func lalr(action, next int32) int32 {
	a := -action - 3
	for ; tmLalr[a] >= 0; a += 2 {
		if tmLalr[a] == next {
			break
		}
	}
	return tmLalr[a+1]
}

func gotoState(state int16, symbol int32) int16 {
	min := tmGoto[symbol]
	max := tmGoto[symbol+1]

	if max-min < 32 {
		for i := min; i < max; i += 2 {
			if tmFromTo[i] == state {
				return tmFromTo[i+1]
			}
		}
	} else {
		for min < max {
			e := (min + max) >> 1 &^ int32(1)
			i := tmFromTo[e]
			if i == state {
				return tmFromTo[e+1]
			} else if i < state {
				min = e + 2
			} else {
				max = e
			}
		}
	}
	return -1
}

func (p *Parser) fetchNext(lexer *Lexer, stack []stackEntry) {
restart:
	tok := lexer.Next()
	switch tok {
	case token.INVALID_TOKEN, token.MULTILINECOMMENT, token.COMMENT, token.TEMPLATES:
		s, e := lexer.Pos()
		tok := symbol{int32(tok), s, e}
		p.pending = append(p.pending, tok)
		goto restart
	}
	p.next.symbol = int32(tok)
	p.next.offset, p.next.endoffset = lexer.Pos()
}

func (p *Parser) applyRule(ctx context.Context, rule int32, lhs *stackEntry, rhs []stackEntry, lexer *Lexer) (err error) {
	switch rule {
	case 198: // directive : '%' 'assert' 'empty' rhsSet ';'
		p.listener(Empty, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 199: // directive : '%' 'assert' 'nonempty' rhsSet ';'
		p.listener(NonEmpty, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 207: // inputref : symref 'no-eoi'
		p.listener(NoEoi, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 233: // rhsSuffix : '%' 'prec' symref
		p.listener(Name, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 234: // rhsSuffix : '%' 'shift' symref
		p.listener(Name, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 254: // lookahead_predicate : '!' symref
		p.listener(Not, rhs[0].sym.offset, rhs[0].sym.endoffset)
	}
	if nt := tmRuleType[rule]; nt != 0 {
		p.listener(nt, lhs.sym.offset, lhs.sym.endoffset)
	}
	return
}

func (p *Parser) reportIgnoredToken(ctx context.Context, tok symbol) {
	var t NodeType
	switch token.Token(tok.symbol) {
	case token.INVALID_TOKEN:
		t = InvalidToken
	case token.MULTILINECOMMENT:
		t = MultilineComment
	case token.COMMENT:
		t = Comment
	case token.TEMPLATES:
		t = Templates
	default:
		return
	}
	if debugSyntax {
		fmt.Printf("ignored: %v as %v\n", token.Token(tok.symbol), t)
	}
	p.listener(t, tok.offset, tok.endoffset)
}
