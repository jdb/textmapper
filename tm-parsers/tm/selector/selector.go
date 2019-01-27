// generated by Textmapper; DO NOT EDIT

package selector

import (
	"github.com/inspirer/textmapper/tm-parsers/tm"
)

type Selector func(nt tm.NodeType) bool

var (
	Any                  = func(t tm.NodeType) bool { return true }
	AnnotationImpl       = func(t tm.NodeType) bool { return t == tm.AnnotationImpl }
	Annotations          = func(t tm.NodeType) bool { return t == tm.Annotations }
	ArgumentFalse        = func(t tm.NodeType) bool { return t == tm.ArgumentFalse }
	ArgumentImpl         = func(t tm.NodeType) bool { return t == tm.ArgumentImpl }
	ArgumentTrue         = func(t tm.NodeType) bool { return t == tm.ArgumentTrue }
	Array                = func(t tm.NodeType) bool { return t == tm.Array }
	Assoc                = func(t tm.NodeType) bool { return t == tm.Assoc }
	BooleanLiteral       = func(t tm.NodeType) bool { return t == tm.BooleanLiteral }
	ClassType            = func(t tm.NodeType) bool { return t == tm.ClassType }
	Command              = func(t tm.NodeType) bool { return t == tm.Command }
	DirectiveAssert      = func(t tm.NodeType) bool { return t == tm.DirectiveAssert }
	DirectiveBrackets    = func(t tm.NodeType) bool { return t == tm.DirectiveBrackets }
	DirectiveInput       = func(t tm.NodeType) bool { return t == tm.DirectiveInput }
	DirectiveInterface   = func(t tm.NodeType) bool { return t == tm.DirectiveInterface }
	DirectivePrio        = func(t tm.NodeType) bool { return t == tm.DirectivePrio }
	DirectiveSet         = func(t tm.NodeType) bool { return t == tm.DirectiveSet }
	ExclusiveStartConds  = func(t tm.NodeType) bool { return t == tm.ExclusiveStartConds }
	File                 = func(t tm.NodeType) bool { return t == tm.File }
	Header               = func(t tm.NodeType) bool { return t == tm.Header }
	Identifier           = func(t tm.NodeType) bool { return t == tm.Identifier }
	Implements           = func(t tm.NodeType) bool { return t == tm.Implements }
	Import               = func(t tm.NodeType) bool { return t == tm.Import }
	InclusiveStartConds  = func(t tm.NodeType) bool { return t == tm.InclusiveStartConds }
	InlineParameter      = func(t tm.NodeType) bool { return t == tm.InlineParameter }
	Inputref             = func(t tm.NodeType) bool { return t == tm.Inputref }
	IntegerLiteral       = func(t tm.NodeType) bool { return t == tm.IntegerLiteral }
	InterfaceType        = func(t tm.NodeType) bool { return t == tm.InterfaceType }
	KeyValue             = func(t tm.NodeType) bool { return t == tm.KeyValue }
	Lexeme               = func(t tm.NodeType) bool { return t == tm.Lexeme }
	LexemeAttribute      = func(t tm.NodeType) bool { return t == tm.LexemeAttribute }
	LexemeAttrs          = func(t tm.NodeType) bool { return t == tm.LexemeAttrs }
	LexerSection         = func(t tm.NodeType) bool { return t == tm.LexerSection }
	LexerState           = func(t tm.NodeType) bool { return t == tm.LexerState }
	ListSeparator        = func(t tm.NodeType) bool { return t == tm.ListSeparator }
	LookaheadPredicate   = func(t tm.NodeType) bool { return t == tm.LookaheadPredicate }
	NamedPattern         = func(t tm.NodeType) bool { return t == tm.NamedPattern }
	Nonterm              = func(t tm.NodeType) bool { return t == tm.Nonterm }
	NontermParams        = func(t tm.NodeType) bool { return t == tm.NontermParams }
	ParamModifier        = func(t tm.NodeType) bool { return t == tm.ParamModifier }
	ParamRef             = func(t tm.NodeType) bool { return t == tm.ParamRef }
	ParamType            = func(t tm.NodeType) bool { return t == tm.ParamType }
	ParserSection        = func(t tm.NodeType) bool { return t == tm.ParserSection }
	Pattern              = func(t tm.NodeType) bool { return t == tm.Pattern }
	Predicate            = func(t tm.NodeType) bool { return t == tm.Predicate }
	PredicateAnd         = func(t tm.NodeType) bool { return t == tm.PredicateAnd }
	PredicateEq          = func(t tm.NodeType) bool { return t == tm.PredicateEq }
	PredicateNot         = func(t tm.NodeType) bool { return t == tm.PredicateNot }
	PredicateNotEq       = func(t tm.NodeType) bool { return t == tm.PredicateNotEq }
	PredicateOr          = func(t tm.NodeType) bool { return t == tm.PredicateOr }
	RawType              = func(t tm.NodeType) bool { return t == tm.RawType }
	References           = func(t tm.NodeType) bool { return t == tm.References }
	ReportAs             = func(t tm.NodeType) bool { return t == tm.ReportAs }
	ReportClause         = func(t tm.NodeType) bool { return t == tm.ReportClause }
	RhsAnnotated         = func(t tm.NodeType) bool { return t == tm.RhsAnnotated }
	RhsAsLiteral         = func(t tm.NodeType) bool { return t == tm.RhsAsLiteral }
	RhsAssignment        = func(t tm.NodeType) bool { return t == tm.RhsAssignment }
	RhsCast              = func(t tm.NodeType) bool { return t == tm.RhsCast }
	RhsIgnored           = func(t tm.NodeType) bool { return t == tm.RhsIgnored }
	RhsLookahead         = func(t tm.NodeType) bool { return t == tm.RhsLookahead }
	RhsNested            = func(t tm.NodeType) bool { return t == tm.RhsNested }
	RhsOptional          = func(t tm.NodeType) bool { return t == tm.RhsOptional }
	RhsPlusAssignment    = func(t tm.NodeType) bool { return t == tm.RhsPlusAssignment }
	RhsPlusList          = func(t tm.NodeType) bool { return t == tm.RhsPlusList }
	RhsQuantifier        = func(t tm.NodeType) bool { return t == tm.RhsQuantifier }
	RhsSet               = func(t tm.NodeType) bool { return t == tm.RhsSet }
	RhsStarList          = func(t tm.NodeType) bool { return t == tm.RhsStarList }
	RhsSuffix            = func(t tm.NodeType) bool { return t == tm.RhsSuffix }
	RhsSymbol            = func(t tm.NodeType) bool { return t == tm.RhsSymbol }
	Rule                 = func(t tm.NodeType) bool { return t == tm.Rule }
	SetAnd               = func(t tm.NodeType) bool { return t == tm.SetAnd }
	SetComplement        = func(t tm.NodeType) bool { return t == tm.SetComplement }
	SetCompound          = func(t tm.NodeType) bool { return t == tm.SetCompound }
	SetOr                = func(t tm.NodeType) bool { return t == tm.SetOr }
	SetSymbol            = func(t tm.NodeType) bool { return t == tm.SetSymbol }
	StartConditions      = func(t tm.NodeType) bool { return t == tm.StartConditions }
	StartConditionsScope = func(t tm.NodeType) bool { return t == tm.StartConditionsScope }
	StateMarker          = func(t tm.NodeType) bool { return t == tm.StateMarker }
	Stateref             = func(t tm.NodeType) bool { return t == tm.Stateref }
	StringLiteral        = func(t tm.NodeType) bool { return t == tm.StringLiteral }
	SubType              = func(t tm.NodeType) bool { return t == tm.SubType }
	Symref               = func(t tm.NodeType) bool { return t == tm.Symref }
	SymrefArgs           = func(t tm.NodeType) bool { return t == tm.SymrefArgs }
	SyntaxProblem        = func(t tm.NodeType) bool { return t == tm.SyntaxProblem }
	TemplateParam        = func(t tm.NodeType) bool { return t == tm.TemplateParam }
	VoidType             = func(t tm.NodeType) bool { return t == tm.VoidType }
	InvalidToken         = func(t tm.NodeType) bool { return t == tm.InvalidToken }
	MultilineComment     = func(t tm.NodeType) bool { return t == tm.MultilineComment }
	Comment              = func(t tm.NodeType) bool { return t == tm.Comment }
	Templates            = func(t tm.NodeType) bool { return t == tm.Templates }
	Annotation           = OneOf(tm.Annotation...)
	Argument             = OneOf(tm.Argument...)
	Expression           = OneOf(tm.Expression...)
	GrammarPart          = OneOf(tm.GrammarPart...)
	LexerPart            = OneOf(tm.LexerPart...)
	Literal              = OneOf(tm.Literal...)
	NontermParam         = OneOf(tm.NontermParam...)
	NontermType          = OneOf(tm.NontermType...)
	Option               = OneOf(tm.Option...)
	ParamValue           = OneOf(tm.ParamValue...)
	PredicateExpression  = OneOf(tm.PredicateExpression...)
	RhsPart              = OneOf(tm.RhsPart...)
	Rule0                = OneOf(tm.Rule0...)
	SetExpression        = OneOf(tm.SetExpression...)
)

func OneOf(types ...tm.NodeType) Selector {
	if len(types) == 0 {
		return func(tm.NodeType) bool { return false }
	}
	const bits = 32
	max := 1
	for _, t := range types {
		if int(t) > max {
			max = int(t)
		}
	}
	size := (max + bits) / bits
	bitarr := make([]uint32, size)
	for _, t := range types {
		bitarr[uint(t)/bits] |= 1 << (uint(t) % bits)
	}
	return func(t tm.NodeType) bool {
		i := uint(t) / bits
		return int(i) < len(bitarr) && bitarr[i]&(1<<(uint(t)%bits)) != 0
	}
}
