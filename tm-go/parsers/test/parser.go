// generated by Textmapper; DO NOT EDIT

package test

import (
	"fmt"
)

// ErrorHandler is called every time a parser is unable to process some part of the input.
// This handler can return false to abort the parser.
type ErrorHandler func(err SyntaxError) bool

// Parser is a table-driven LALR parser for test.
type Parser struct {
	eh       ErrorHandler
	listener Listener

	stack         []stackEntry
	lexer         *Lexer
	next          symbol
	ignoredTokens []symbol // to be reported with the next shift
}

type SyntaxError struct {
	Offset    int
	Endoffset int
}

func (e SyntaxError) Error() string {
	return "syntax error"
}

type symbol struct {
	symbol    int32
	offset    int
	endoffset int
}

type stackEntry struct {
	sym   symbol
	state int8
}

func (p *Parser) Init(eh ErrorHandler, l Listener) {
	p.eh = eh
	p.listener = l
}

const (
	startStackSize       = 512
	startTokenBufferSize = 16
	noToken              = int32(UNAVAILABLE)
	eoiToken             = int32(EOI)
	debugSyntax          = false
)

func (p *Parser) Parse(lexer *Lexer) error {
	return p.parse(0, 18, lexer)
}

func (p *Parser) parse(start, end int8, lexer *Lexer) error {
	if cap(p.stack) < startStackSize {
		p.stack = make([]stackEntry, 0, startStackSize)
	}
	if cap(p.ignoredTokens) < startTokenBufferSize {
		p.ignoredTokens = make([]symbol, 0, startTokenBufferSize)
	} else {
		p.ignoredTokens = p.ignoredTokens[:0]
	}
	state := start
	var lastErr SyntaxError
	recovering := 0

	p.stack = append(p.stack[:0], stackEntry{state: state})
	p.lexer = lexer
	p.fetchNext()

	for state != end {
		action := tmAction[state]
		if action < -2 {
			// Lookahead is needed.
			if p.next.symbol == noToken {
				p.fetchNext()
			}
			action = lalr(action, p.next.symbol)
		}

		if action >= 0 {
			// Reduce.
			rule := action
			ln := int(tmRuleLen[rule])

			var entry stackEntry
			entry.sym.symbol = tmRuleSymbol[rule]
			if ln == 0 {
				entry.sym.offset, _ = lexer.Pos()
				entry.sym.endoffset = entry.sym.offset
			} else {
				entry.sym.offset = p.stack[len(p.stack)-ln].sym.offset
				entry.sym.endoffset = p.stack[len(p.stack)-1].sym.endoffset
			}
			p.applyRule(rule, &entry, p.stack[len(p.stack)-ln:])
			if debugSyntax {
				fmt.Printf("reduced to: %v\n", Symbol(entry.sym.symbol))
			}
			p.stack = p.stack[:len(p.stack)-ln]
			state = gotoState(p.stack[len(p.stack)-1].state, entry.sym.symbol)
			entry.state = state
			p.stack = append(p.stack, entry)

		} else if action == -1 {
			// Shift.
			if p.next.symbol == noToken {
				p.fetchNext()
			}
			state = gotoState(state, p.next.symbol)
			p.stack = append(p.stack, stackEntry{
				sym:   p.next,
				state: state,
			})
			if debugSyntax {
				fmt.Printf("shift: %v (%s)\n", Symbol(p.next.symbol), lexer.Text())
			}
			if len(p.ignoredTokens) > 0 {
				p.reportIgnoredTokens()
			}
			switch Token(p.next.symbol) {
			case IDENTIFIER:
				p.listener(Identifier, p.next.offset, p.next.endoffset)
			}
			if state != -1 && p.next.symbol != eoiToken {
				p.next.symbol = noToken
			}
			if recovering > 0 {
				recovering--
			}
		}

		if action == -2 || state == -1 {
			if p.recover() {
				state = p.stack[len(p.stack)-1].state
				if recovering == 0 {
					offset, endoffset := lexer.Pos()
					lastErr = SyntaxError{
						Offset:    offset,
						Endoffset: endoffset,
					}
					if !p.eh(lastErr) {
						return lastErr
					}
				}
				if recovering >= 3 {
					p.fetchNext()
				}
				recovering = 4
				continue
			}
			if len(p.stack) == 0 {
				state = start
				p.stack = append(p.stack, stackEntry{state: state})
			}
			break
		}
	}

	if state != end {
		if recovering > 0 {
			return lastErr
		}
		offset, endoffset := lexer.Pos()
		err := SyntaxError{
			Offset:    offset,
			Endoffset: endoffset,
		}
		return err
	}

	return nil
}

const errSymbol = 21

func (p *Parser) recover() bool {
	if p.next.symbol == noToken {
		p.fetchNext()
	}
	if p.next.symbol == eoiToken {
		return false
	}
	e, _ := p.lexer.Pos()
	s := e
	for len(p.stack) > 0 && gotoState(p.stack[len(p.stack)-1].state, errSymbol) == -1 {
		// TODO cleanup
		p.stack = p.stack[:len(p.stack)-1]
		if len(p.stack) > 0 {
			s = p.stack[len(p.stack)-1].sym.offset
		}
	}
	if len(p.stack) > 0 {
		state := gotoState(p.stack[len(p.stack)-1].state, errSymbol)
		p.stack = append(p.stack, stackEntry{
			sym:   symbol{errSymbol, s, e},
			state: state,
		})
		return true
	}
	return false
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

func gotoState(state int8, symbol int32) int8 {
	min := tmGoto[symbol]
	max := tmGoto[symbol+1] - 1

	for min <= max {
		e := (min + max) >> 1
		i := tmFrom[e]
		if i == state {
			return tmTo[e]
		} else if i < state {
			min = e + 1
		} else {
			max = e - 1
		}
	}
	return -1
}

func (p *Parser) fetchNext() {
restart:
	tok := p.lexer.Next()
	switch tok {
	case MULTILINECOMMENT, SINGLELINECOMMENT, INVALID_TOKEN:
		s, e := p.lexer.Pos()
		p.ignoredTokens = append(p.ignoredTokens, symbol{int32(tok), s, e})
		goto restart
	}
	p.next.symbol = int32(tok)
	p.next.offset, p.next.endoffset = p.lexer.Pos()
}

func (p *Parser) applyRule(rule int32, lhs *stackEntry, rhs []stackEntry) {
	nt := ruleNodeType[rule]
	if nt == 0 {
		return
	}
	p.listener(nt, lhs.sym.offset, lhs.sym.endoffset)
}

func (p *Parser) reportIgnoredTokens() {
	for _, c := range p.ignoredTokens {
		var t NodeType
		switch Token(c.symbol) {
		case MULTILINECOMMENT:
			t = MultiLineComment
		case SINGLELINECOMMENT:
			t = SingleLineComment
		case INVALID_TOKEN:
			t = InvalidToken
		default:
			continue
		}
		p.listener(t, c.offset, c.endoffset)
	}
	p.ignoredTokens = p.ignoredTokens[:0]
}
