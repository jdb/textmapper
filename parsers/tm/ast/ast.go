// generated by Textmapper; DO NOT EDIT

package ast

import (
	"github.com/inspirer/textmapper/parsers/tm/selector"
)

// Interfaces.

type TmNode interface {
	TmNode() *Node
}

type Token struct {
	*Node
}

type NilNode struct{}

var nilInstance = &NilNode{}

// All types implement TmNode.
func (n ArgumentFalse) TmNode() *Node        { return n.Node }
func (n ArgumentTrue) TmNode() *Node         { return n.Node }
func (n ArgumentVal) TmNode() *Node          { return n.Node }
func (n Array) TmNode() *Node                { return n.Node }
func (n Assoc) TmNode() *Node                { return n.Node }
func (n BooleanLiteral) TmNode() *Node       { return n.Node }
func (n Command) TmNode() *Node              { return n.Node }
func (n DirectiveAssert) TmNode() *Node      { return n.Node }
func (n DirectiveBrackets) TmNode() *Node    { return n.Node }
func (n DirectiveExpect) TmNode() *Node      { return n.Node }
func (n DirectiveExpectRR) TmNode() *Node    { return n.Node }
func (n DirectiveInject) TmNode() *Node      { return n.Node }
func (n DirectiveInput) TmNode() *Node       { return n.Node }
func (n DirectiveInterface) TmNode() *Node   { return n.Node }
func (n DirectivePrio) TmNode() *Node        { return n.Node }
func (n DirectiveSet) TmNode() *Node         { return n.Node }
func (n Empty) TmNode() *Node                { return n.Node }
func (n ExclusiveStartConds) TmNode() *Node  { return n.Node }
func (n Extend) TmNode() *Node               { return n.Node }
func (n File) TmNode() *Node                 { return n.Node }
func (n Header) TmNode() *Node               { return n.Node }
func (n Identifier) TmNode() *Node           { return n.Node }
func (n Import) TmNode() *Node               { return n.Node }
func (n InclusiveStartConds) TmNode() *Node  { return n.Node }
func (n Inline) TmNode() *Node               { return n.Node }
func (n InlineParameter) TmNode() *Node      { return n.Node }
func (n Inputref) TmNode() *Node             { return n.Node }
func (n IntegerLiteral) TmNode() *Node       { return n.Node }
func (n Lexeme) TmNode() *Node               { return n.Node }
func (n LexemeAttribute) TmNode() *Node      { return n.Node }
func (n LexemeAttrs) TmNode() *Node          { return n.Node }
func (n LexerSection) TmNode() *Node         { return n.Node }
func (n LexerState) TmNode() *Node           { return n.Node }
func (n ListSeparator) TmNode() *Node        { return n.Node }
func (n LookaheadPredicate) TmNode() *Node   { return n.Node }
func (n Name) TmNode() *Node                 { return n.Node }
func (n NamedPattern) TmNode() *Node         { return n.Node }
func (n NoEoi) TmNode() *Node                { return n.Node }
func (n NonEmpty) TmNode() *Node             { return n.Node }
func (n Nonterm) TmNode() *Node              { return n.Node }
func (n NontermAlias) TmNode() *Node         { return n.Node }
func (n NontermParams) TmNode() *Node        { return n.Node }
func (n Not) TmNode() *Node                  { return n.Node }
func (n Option) TmNode() *Node               { return n.Node }
func (n ParamModifier) TmNode() *Node        { return n.Node }
func (n ParamRef) TmNode() *Node             { return n.Node }
func (n ParamType) TmNode() *Node            { return n.Node }
func (n ParserSection) TmNode() *Node        { return n.Node }
func (n Pattern) TmNode() *Node              { return n.Node }
func (n Predicate) TmNode() *Node            { return n.Node }
func (n PredicateAnd) TmNode() *Node         { return n.Node }
func (n PredicateEq) TmNode() *Node          { return n.Node }
func (n PredicateNot) TmNode() *Node         { return n.Node }
func (n PredicateNotEq) TmNode() *Node       { return n.Node }
func (n PredicateOr) TmNode() *Node          { return n.Node }
func (n RawType) TmNode() *Node              { return n.Node }
func (n ReportAs) TmNode() *Node             { return n.Node }
func (n ReportClause) TmNode() *Node         { return n.Node }
func (n RhsAlias) TmNode() *Node             { return n.Node }
func (n RhsAssignment) TmNode() *Node        { return n.Node }
func (n RhsCast) TmNode() *Node              { return n.Node }
func (n RhsIgnored) TmNode() *Node           { return n.Node }
func (n RhsLookahead) TmNode() *Node         { return n.Node }
func (n RhsNested) TmNode() *Node            { return n.Node }
func (n RhsOptional) TmNode() *Node          { return n.Node }
func (n RhsPlusAssignment) TmNode() *Node    { return n.Node }
func (n RhsPlusList) TmNode() *Node          { return n.Node }
func (n RhsPlusQuantifier) TmNode() *Node    { return n.Node }
func (n RhsSet) TmNode() *Node               { return n.Node }
func (n RhsStarList) TmNode() *Node          { return n.Node }
func (n RhsStarQuantifier) TmNode() *Node    { return n.Node }
func (n RhsSuffix) TmNode() *Node            { return n.Node }
func (n RhsSymbol) TmNode() *Node            { return n.Node }
func (n Rule) TmNode() *Node                 { return n.Node }
func (n SetAnd) TmNode() *Node               { return n.Node }
func (n SetComplement) TmNode() *Node        { return n.Node }
func (n SetCompound) TmNode() *Node          { return n.Node }
func (n SetOr) TmNode() *Node                { return n.Node }
func (n SetSymbol) TmNode() *Node            { return n.Node }
func (n StartConditions) TmNode() *Node      { return n.Node }
func (n StartConditionsScope) TmNode() *Node { return n.Node }
func (n StateMarker) TmNode() *Node          { return n.Node }
func (n Stateref) TmNode() *Node             { return n.Node }
func (n StringLiteral) TmNode() *Node        { return n.Node }
func (n Symref) TmNode() *Node               { return n.Node }
func (n SymrefArgs) TmNode() *Node           { return n.Node }
func (n SyntaxProblem) TmNode() *Node        { return n.Node }
func (n TemplateParam) TmNode() *Node        { return n.Node }
func (n InvalidToken) TmNode() *Node         { return n.Node }
func (n MultilineComment) TmNode() *Node     { return n.Node }
func (n Comment) TmNode() *Node              { return n.Node }
func (n Templates) TmNode() *Node            { return n.Node }
func (NilNode) TmNode() *Node                { return nil }

type Argument interface {
	TmNode
	argumentNode()
}

// argumentNode() ensures that only the following types can be
// assigned to Argument.
func (ArgumentFalse) argumentNode() {}
func (ArgumentTrue) argumentNode()  {}
func (ArgumentVal) argumentNode()   {}
func (NilNode) argumentNode()       {}

type Expression interface {
	TmNode
	expressionNode()
}

// expressionNode() ensures that only the following types can be
// assigned to Expression.
func (Array) expressionNode()          {}
func (BooleanLiteral) expressionNode() {}
func (IntegerLiteral) expressionNode() {}
func (StringLiteral) expressionNode()  {}
func (Symref) expressionNode()         {}
func (SyntaxProblem) expressionNode()  {}
func (NilNode) expressionNode()        {}

type GrammarPart interface {
	TmNode
	grammarPartNode()
}

// grammarPartNode() ensures that only the following types can be
// assigned to GrammarPart.
func (DirectiveAssert) grammarPartNode()    {}
func (DirectiveExpect) grammarPartNode()    {}
func (DirectiveExpectRR) grammarPartNode()  {}
func (DirectiveInject) grammarPartNode()    {}
func (DirectiveInput) grammarPartNode()     {}
func (DirectiveInterface) grammarPartNode() {}
func (DirectivePrio) grammarPartNode()      {}
func (DirectiveSet) grammarPartNode()       {}
func (Nonterm) grammarPartNode()            {}
func (SyntaxProblem) grammarPartNode()      {}
func (TemplateParam) grammarPartNode()      {}
func (NilNode) grammarPartNode()            {}

type LexerPart interface {
	TmNode
	lexerPartNode()
}

// lexerPartNode() ensures that only the following types can be
// assigned to LexerPart.
func (DirectiveBrackets) lexerPartNode()    {}
func (ExclusiveStartConds) lexerPartNode()  {}
func (InclusiveStartConds) lexerPartNode()  {}
func (Lexeme) lexerPartNode()               {}
func (NamedPattern) lexerPartNode()         {}
func (StartConditionsScope) lexerPartNode() {}
func (SyntaxProblem) lexerPartNode()        {}
func (NilNode) lexerPartNode()              {}

type Literal interface {
	TmNode
	literalNode()
}

// literalNode() ensures that only the following types can be
// assigned to Literal.
func (BooleanLiteral) literalNode() {}
func (IntegerLiteral) literalNode() {}
func (StringLiteral) literalNode()  {}
func (NilNode) literalNode()        {}

type NontermParam interface {
	TmNode
	nontermParamNode()
}

// nontermParamNode() ensures that only the following types can be
// assigned to NontermParam.
func (InlineParameter) nontermParamNode() {}
func (ParamRef) nontermParamNode()        {}
func (NilNode) nontermParamNode()         {}

type ParamValue interface {
	TmNode
	paramValueNode()
}

// paramValueNode() ensures that only the following types can be
// assigned to ParamValue.
func (BooleanLiteral) paramValueNode() {}
func (IntegerLiteral) paramValueNode() {}
func (ParamRef) paramValueNode()       {}
func (StringLiteral) paramValueNode()  {}
func (NilNode) paramValueNode()        {}

type PredicateExpression interface {
	TmNode
	predicateExpressionNode()
}

// predicateExpressionNode() ensures that only the following types can be
// assigned to PredicateExpression.
func (ParamRef) predicateExpressionNode()       {}
func (PredicateAnd) predicateExpressionNode()   {}
func (PredicateEq) predicateExpressionNode()    {}
func (PredicateNot) predicateExpressionNode()   {}
func (PredicateNotEq) predicateExpressionNode() {}
func (PredicateOr) predicateExpressionNode()    {}
func (NilNode) predicateExpressionNode()        {}

type RhsPart interface {
	TmNode
	rhsPartNode()
}

// rhsPartNode() ensures that only the following types can be
// assigned to RhsPart.
func (Command) rhsPartNode()           {}
func (RhsAlias) rhsPartNode()          {}
func (RhsAssignment) rhsPartNode()     {}
func (RhsCast) rhsPartNode()           {}
func (RhsIgnored) rhsPartNode()        {}
func (RhsLookahead) rhsPartNode()      {}
func (RhsNested) rhsPartNode()         {}
func (RhsOptional) rhsPartNode()       {}
func (RhsPlusAssignment) rhsPartNode() {}
func (RhsPlusList) rhsPartNode()       {}
func (RhsPlusQuantifier) rhsPartNode() {}
func (RhsSet) rhsPartNode()            {}
func (RhsStarList) rhsPartNode()       {}
func (RhsStarQuantifier) rhsPartNode() {}
func (RhsSymbol) rhsPartNode()         {}
func (StateMarker) rhsPartNode()       {}
func (SyntaxProblem) rhsPartNode()     {}
func (NilNode) rhsPartNode()           {}

type Rule0 interface {
	TmNode
	rule0Node()
}

// rule0Node() ensures that only the following types can be
// assigned to Rule0.
func (Rule) rule0Node()          {}
func (SyntaxProblem) rule0Node() {}
func (NilNode) rule0Node()       {}

type SetExpression interface {
	TmNode
	setExpressionNode()
}

// setExpressionNode() ensures that only the following types can be
// assigned to SetExpression.
func (SetAnd) setExpressionNode()        {}
func (SetComplement) setExpressionNode() {}
func (SetCompound) setExpressionNode()   {}
func (SetOr) setExpressionNode()         {}
func (SetSymbol) setExpressionNode()     {}
func (NilNode) setExpressionNode()       {}

// Types.

type ArgumentFalse struct {
	*Node
}

func (n ArgumentFalse) Name() ParamRef {
	child := n.Child(selector.ParamRef)
	return ParamRef{child}
}

type ArgumentTrue struct {
	*Node
}

func (n ArgumentTrue) Name() ParamRef {
	child := n.Child(selector.ParamRef)
	return ParamRef{child}
}

type ArgumentVal struct {
	*Node
}

func (n ArgumentVal) Name() ParamRef {
	child := n.Child(selector.ParamRef)
	return ParamRef{child}
}

func (n ArgumentVal) Val() (ParamValue, bool) {
	child := n.Child(selector.ParamRef).Next(selector.ParamValue)
	return ToTmNode(child).(ParamValue), child.IsValid()
}

type Array struct {
	*Node
}

func (n Array) Expression() []Expression {
	nodes := n.Children(selector.Expression)
	var ret = make([]Expression, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(Expression))
	}
	return ret
}

type Assoc struct {
	*Node
}

type BooleanLiteral struct {
	*Node
}

type Command struct {
	*Node
}

type DirectiveAssert struct {
	*Node
}

func (n DirectiveAssert) Empty() (Empty, bool) {
	child := n.Child(selector.Empty)
	return Empty{child}, child.IsValid()
}

func (n DirectiveAssert) NonEmpty() (NonEmpty, bool) {
	child := n.Child(selector.NonEmpty)
	return NonEmpty{child}, child.IsValid()
}

func (n DirectiveAssert) RhsSet() RhsSet {
	child := n.Child(selector.RhsSet)
	return RhsSet{child}
}

type DirectiveBrackets struct {
	*Node
}

func (n DirectiveBrackets) Opening() Symref {
	child := n.Child(selector.Symref)
	return Symref{child}
}

func (n DirectiveBrackets) Closing() Symref {
	child := n.Child(selector.Symref).Next(selector.Symref)
	return Symref{child}
}

type DirectiveExpect struct {
	*Node
}

func (n DirectiveExpect) IntegerLiteral() IntegerLiteral {
	child := n.Child(selector.IntegerLiteral)
	return IntegerLiteral{child}
}

type DirectiveExpectRR struct {
	*Node
}

func (n DirectiveExpectRR) IntegerLiteral() IntegerLiteral {
	child := n.Child(selector.IntegerLiteral)
	return IntegerLiteral{child}
}

type DirectiveInject struct {
	*Node
}

func (n DirectiveInject) Symref() Symref {
	child := n.Child(selector.Symref)
	return Symref{child}
}

func (n DirectiveInject) ReportClause() ReportClause {
	child := n.Child(selector.ReportClause)
	return ReportClause{child}
}

type DirectiveInput struct {
	*Node
}

func (n DirectiveInput) InputRefs() []Inputref {
	nodes := n.Children(selector.Inputref)
	var ret = make([]Inputref, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, Inputref{node})
	}
	return ret
}

type DirectiveInterface struct {
	*Node
}

func (n DirectiveInterface) Ids() []Identifier {
	nodes := n.Children(selector.Identifier)
	var ret = make([]Identifier, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, Identifier{node})
	}
	return ret
}

type DirectivePrio struct {
	*Node
}

func (n DirectivePrio) Assoc() Assoc {
	child := n.Child(selector.Assoc)
	return Assoc{child}
}

func (n DirectivePrio) Symbols() []Symref {
	nodes := n.Children(selector.Symref)
	var ret = make([]Symref, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, Symref{node})
	}
	return ret
}

type DirectiveSet struct {
	*Node
}

func (n DirectiveSet) Name() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

func (n DirectiveSet) RhsSet() RhsSet {
	child := n.Child(selector.RhsSet)
	return RhsSet{child}
}

type Empty struct {
	*Node
}

type ExclusiveStartConds struct {
	*Node
}

func (n ExclusiveStartConds) States() []LexerState {
	nodes := n.Children(selector.LexerState)
	var ret = make([]LexerState, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, LexerState{node})
	}
	return ret
}

type Extend struct {
	*Node
}

type File struct {
	*Node
}

func (n File) Header() Header {
	child := n.Child(selector.Header)
	return Header{child}
}

func (n File) Imports() []Import {
	nodes := n.Children(selector.Import)
	var ret = make([]Import, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, Import{node})
	}
	return ret
}

func (n File) Options() []Option {
	nodes := n.Children(selector.Option)
	var ret = make([]Option, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, Option{node})
	}
	return ret
}

func (n File) SyntaxProblem() (SyntaxProblem, bool) {
	child := n.Child(selector.SyntaxProblem)
	return SyntaxProblem{child}, child.IsValid()
}

func (n File) Lexer() (LexerSection, bool) {
	child := n.Child(selector.LexerSection)
	return LexerSection{child}, child.IsValid()
}

func (n File) Parser() (ParserSection, bool) {
	child := n.Child(selector.ParserSection)
	return ParserSection{child}, child.IsValid()
}

type Header struct {
	*Node
}

func (n Header) Name() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

func (n Header) Target() (Identifier, bool) {
	child := n.Child(selector.Identifier).Next(selector.Identifier)
	return Identifier{child}, child.IsValid()
}

type Identifier struct {
	*Node
}

type Import struct {
	*Node
}

func (n Import) Alias() (Identifier, bool) {
	child := n.Child(selector.Identifier)
	return Identifier{child}, child.IsValid()
}

func (n Import) Path() StringLiteral {
	child := n.Child(selector.StringLiteral)
	return StringLiteral{child}
}

type InclusiveStartConds struct {
	*Node
}

func (n InclusiveStartConds) States() []LexerState {
	nodes := n.Children(selector.LexerState)
	var ret = make([]LexerState, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, LexerState{node})
	}
	return ret
}

type Inline struct {
	*Node
}

type InlineParameter struct {
	*Node
}

func (n InlineParameter) ParamType() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

func (n InlineParameter) Name() Identifier {
	child := n.Child(selector.Identifier).Next(selector.Identifier)
	return Identifier{child}
}

func (n InlineParameter) ParamValue() (ParamValue, bool) {
	child := n.Child(selector.ParamValue)
	return ToTmNode(child).(ParamValue), child.IsValid()
}

type Inputref struct {
	*Node
}

func (n Inputref) Reference() Symref {
	child := n.Child(selector.Symref)
	return Symref{child}
}

func (n Inputref) NoEoi() (NoEoi, bool) {
	child := n.Child(selector.NoEoi)
	return NoEoi{child}, child.IsValid()
}

type IntegerLiteral struct {
	*Node
}

type Lexeme struct {
	*Node
}

func (n Lexeme) StartConditions() (StartConditions, bool) {
	child := n.Child(selector.StartConditions)
	return StartConditions{child}, child.IsValid()
}

func (n Lexeme) Name() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

func (n Lexeme) RawType() (RawType, bool) {
	child := n.Child(selector.RawType)
	return RawType{child}, child.IsValid()
}

func (n Lexeme) Pattern() (Pattern, bool) {
	child := n.Child(selector.Pattern)
	return Pattern{child}, child.IsValid()
}

func (n Lexeme) Priority() (IntegerLiteral, bool) {
	child := n.Child(selector.IntegerLiteral)
	return IntegerLiteral{child}, child.IsValid()
}

func (n Lexeme) Attrs() (LexemeAttrs, bool) {
	child := n.Child(selector.LexemeAttrs)
	return LexemeAttrs{child}, child.IsValid()
}

func (n Lexeme) Command() (Command, bool) {
	child := n.Child(selector.Command)
	return Command{child}, child.IsValid()
}

type LexemeAttribute struct {
	*Node
}

type LexemeAttrs struct {
	*Node
}

func (n LexemeAttrs) LexemeAttribute() LexemeAttribute {
	child := n.Child(selector.LexemeAttribute)
	return LexemeAttribute{child}
}

type LexerSection struct {
	*Node
}

func (n LexerSection) LexerPart() []LexerPart {
	nodes := n.Children(selector.LexerPart)
	var ret = make([]LexerPart, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(LexerPart))
	}
	return ret
}

type LexerState struct {
	*Node
}

func (n LexerState) Name() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

type ListSeparator struct {
	*Node
}

func (n ListSeparator) Separator() []Symref {
	nodes := n.Children(selector.Symref)
	var ret = make([]Symref, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, Symref{node})
	}
	return ret
}

type LookaheadPredicate struct {
	*Node
}

func (n LookaheadPredicate) Not() (Not, bool) {
	child := n.Child(selector.Not)
	return Not{child}, child.IsValid()
}

func (n LookaheadPredicate) Symref() Symref {
	child := n.Child(selector.Symref)
	return Symref{child}
}

type Name struct {
	*Node
}

type NamedPattern struct {
	*Node
}

func (n NamedPattern) Name() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

func (n NamedPattern) Pattern() Pattern {
	child := n.Child(selector.Pattern)
	return Pattern{child}
}

type NoEoi struct {
	*Node
}

type NonEmpty struct {
	*Node
}

type Nonterm struct {
	*Node
}

func (n Nonterm) Extend() (Extend, bool) {
	child := n.Child(selector.Extend)
	return Extend{child}, child.IsValid()
}

func (n Nonterm) Inline() (Inline, bool) {
	child := n.Child(selector.Inline)
	return Inline{child}, child.IsValid()
}

func (n Nonterm) Name() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

func (n Nonterm) Params() (NontermParams, bool) {
	child := n.Child(selector.NontermParams)
	return NontermParams{child}, child.IsValid()
}

func (n Nonterm) Alias() (NontermAlias, bool) {
	child := n.Child(selector.NontermAlias)
	return NontermAlias{child}, child.IsValid()
}

func (n Nonterm) RawType() (RawType, bool) {
	child := n.Child(selector.RawType)
	return RawType{child}, child.IsValid()
}

func (n Nonterm) ReportClause() (ReportClause, bool) {
	child := n.Child(selector.ReportClause)
	return ReportClause{child}, child.IsValid()
}

func (n Nonterm) Rule0() []Rule0 {
	nodes := n.Children(selector.Rule0)
	var ret = make([]Rule0, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(Rule0))
	}
	return ret
}

type NontermAlias struct {
	*Node
}

func (n NontermAlias) Name() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

type NontermParams struct {
	*Node
}

func (n NontermParams) List() []NontermParam {
	nodes := n.Children(selector.NontermParam)
	var ret = make([]NontermParam, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(NontermParam))
	}
	return ret
}

type Not struct {
	*Node
}

type Option struct {
	*Node
}

func (n Option) Key() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

func (n Option) Value() Expression {
	child := n.Child(selector.Expression)
	return ToTmNode(child).(Expression)
}

type ParamModifier struct {
	*Node
}

type ParamRef struct {
	*Node
}

func (n ParamRef) Identifier() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

type ParamType struct {
	*Node
}

type ParserSection struct {
	*Node
}

func (n ParserSection) GrammarPart() []GrammarPart {
	nodes := n.Children(selector.GrammarPart)
	var ret = make([]GrammarPart, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(GrammarPart))
	}
	return ret
}

type Pattern struct {
	*Node
}

type Predicate struct {
	*Node
}

func (n Predicate) PredicateExpression() PredicateExpression {
	child := n.Child(selector.PredicateExpression)
	return ToTmNode(child).(PredicateExpression)
}

type PredicateAnd struct {
	*Node
}

func (n PredicateAnd) Left() PredicateExpression {
	child := n.Child(selector.PredicateExpression)
	return ToTmNode(child).(PredicateExpression)
}

func (n PredicateAnd) Right() PredicateExpression {
	child := n.Child(selector.PredicateExpression).Next(selector.PredicateExpression)
	return ToTmNode(child).(PredicateExpression)
}

type PredicateEq struct {
	*Node
}

func (n PredicateEq) ParamRef() ParamRef {
	child := n.Child(selector.ParamRef)
	return ParamRef{child}
}

func (n PredicateEq) Literal() Literal {
	child := n.Child(selector.Literal)
	return ToTmNode(child).(Literal)
}

type PredicateNot struct {
	*Node
}

func (n PredicateNot) ParamRef() ParamRef {
	child := n.Child(selector.ParamRef)
	return ParamRef{child}
}

type PredicateNotEq struct {
	*Node
}

func (n PredicateNotEq) ParamRef() ParamRef {
	child := n.Child(selector.ParamRef)
	return ParamRef{child}
}

func (n PredicateNotEq) Literal() Literal {
	child := n.Child(selector.Literal)
	return ToTmNode(child).(Literal)
}

type PredicateOr struct {
	*Node
}

func (n PredicateOr) Left() PredicateExpression {
	child := n.Child(selector.PredicateExpression)
	return ToTmNode(child).(PredicateExpression)
}

func (n PredicateOr) Right() PredicateExpression {
	child := n.Child(selector.PredicateExpression).Next(selector.PredicateExpression)
	return ToTmNode(child).(PredicateExpression)
}

type RawType struct {
	*Node
}

type ReportAs struct {
	*Node
}

func (n ReportAs) Identifier() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

type ReportClause struct {
	*Node
}

func (n ReportClause) Action() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

func (n ReportClause) Flags() []Identifier {
	nodes := n.Child(selector.Identifier).NextAll(selector.Identifier)
	var ret = make([]Identifier, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, Identifier{node})
	}
	return ret
}

func (n ReportClause) ReportAs() (ReportAs, bool) {
	child := n.Child(selector.ReportAs)
	return ReportAs{child}, child.IsValid()
}

type RhsAlias struct {
	*Node
}

func (n RhsAlias) Inner() RhsPart {
	child := n.Child(selector.RhsPart)
	return ToTmNode(child).(RhsPart)
}

func (n RhsAlias) Name() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

type RhsAssignment struct {
	*Node
}

func (n RhsAssignment) Id() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

func (n RhsAssignment) Inner() RhsPart {
	child := n.Child(selector.RhsPart)
	return ToTmNode(child).(RhsPart)
}

type RhsCast struct {
	*Node
}

func (n RhsCast) Inner() RhsPart {
	child := n.Child(selector.RhsPart)
	return ToTmNode(child).(RhsPart)
}

func (n RhsCast) Target() Symref {
	child := n.Child(selector.Symref)
	return Symref{child}
}

type RhsIgnored struct {
	*Node
}

func (n RhsIgnored) Rule0() []Rule0 {
	nodes := n.Children(selector.Rule0)
	var ret = make([]Rule0, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(Rule0))
	}
	return ret
}

type RhsLookahead struct {
	*Node
}

func (n RhsLookahead) Predicates() []LookaheadPredicate {
	nodes := n.Children(selector.LookaheadPredicate)
	var ret = make([]LookaheadPredicate, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, LookaheadPredicate{node})
	}
	return ret
}

type RhsNested struct {
	*Node
}

func (n RhsNested) Rule0() []Rule0 {
	nodes := n.Children(selector.Rule0)
	var ret = make([]Rule0, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(Rule0))
	}
	return ret
}

type RhsOptional struct {
	*Node
}

func (n RhsOptional) Inner() RhsPart {
	child := n.Child(selector.RhsPart)
	return ToTmNode(child).(RhsPart)
}

type RhsPlusAssignment struct {
	*Node
}

func (n RhsPlusAssignment) Id() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

func (n RhsPlusAssignment) Inner() RhsPart {
	child := n.Child(selector.RhsPart)
	return ToTmNode(child).(RhsPart)
}

type RhsPlusList struct {
	*Node
}

func (n RhsPlusList) RuleParts() []RhsPart {
	nodes := n.Children(selector.RhsPart)
	var ret = make([]RhsPart, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(RhsPart))
	}
	return ret
}

func (n RhsPlusList) ListSeparator() ListSeparator {
	child := n.Child(selector.ListSeparator)
	return ListSeparator{child}
}

type RhsPlusQuantifier struct {
	*Node
}

func (n RhsPlusQuantifier) Inner() RhsPart {
	child := n.Child(selector.RhsPart)
	return ToTmNode(child).(RhsPart)
}

type RhsSet struct {
	*Node
}

func (n RhsSet) Expr() SetExpression {
	child := n.Child(selector.SetExpression)
	return ToTmNode(child).(SetExpression)
}

type RhsStarList struct {
	*Node
}

func (n RhsStarList) RuleParts() []RhsPart {
	nodes := n.Children(selector.RhsPart)
	var ret = make([]RhsPart, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(RhsPart))
	}
	return ret
}

func (n RhsStarList) ListSeparator() ListSeparator {
	child := n.Child(selector.ListSeparator)
	return ListSeparator{child}
}

type RhsStarQuantifier struct {
	*Node
}

func (n RhsStarQuantifier) Inner() RhsPart {
	child := n.Child(selector.RhsPart)
	return ToTmNode(child).(RhsPart)
}

type RhsSuffix struct {
	*Node
}

func (n RhsSuffix) Name() Name {
	child := n.Child(selector.Name)
	return Name{child}
}

func (n RhsSuffix) Symref() Symref {
	child := n.Child(selector.Symref)
	return Symref{child}
}

type RhsSymbol struct {
	*Node
}

func (n RhsSymbol) Reference() Symref {
	child := n.Child(selector.Symref)
	return Symref{child}
}

type Rule struct {
	*Node
}

func (n Rule) Predicate() (Predicate, bool) {
	child := n.Child(selector.Predicate)
	return Predicate{child}, child.IsValid()
}

func (n Rule) RhsPart() []RhsPart {
	nodes := n.Children(selector.RhsPart)
	var ret = make([]RhsPart, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(RhsPart))
	}
	return ret
}

func (n Rule) RhsSuffix() (RhsSuffix, bool) {
	child := n.Child(selector.RhsSuffix)
	return RhsSuffix{child}, child.IsValid()
}

func (n Rule) ReportClause() (ReportClause, bool) {
	child := n.Child(selector.ReportClause)
	return ReportClause{child}, child.IsValid()
}

type SetAnd struct {
	*Node
}

func (n SetAnd) Left() SetExpression {
	child := n.Child(selector.SetExpression)
	return ToTmNode(child).(SetExpression)
}

func (n SetAnd) Right() SetExpression {
	child := n.Child(selector.SetExpression).Next(selector.SetExpression)
	return ToTmNode(child).(SetExpression)
}

type SetComplement struct {
	*Node
}

func (n SetComplement) Inner() SetExpression {
	child := n.Child(selector.SetExpression)
	return ToTmNode(child).(SetExpression)
}

type SetCompound struct {
	*Node
}

func (n SetCompound) Inner() SetExpression {
	child := n.Child(selector.SetExpression)
	return ToTmNode(child).(SetExpression)
}

type SetOr struct {
	*Node
}

func (n SetOr) Left() SetExpression {
	child := n.Child(selector.SetExpression)
	return ToTmNode(child).(SetExpression)
}

func (n SetOr) Right() SetExpression {
	child := n.Child(selector.SetExpression).Next(selector.SetExpression)
	return ToTmNode(child).(SetExpression)
}

type SetSymbol struct {
	*Node
}

func (n SetSymbol) Operator() (Identifier, bool) {
	child := n.Child(selector.Identifier)
	return Identifier{child}, child.IsValid()
}

func (n SetSymbol) Symbol() Symref {
	child := n.Child(selector.Symref)
	return Symref{child}
}

type StartConditions struct {
	*Node
}

func (n StartConditions) Stateref() []Stateref {
	nodes := n.Children(selector.Stateref)
	var ret = make([]Stateref, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, Stateref{node})
	}
	return ret
}

type StartConditionsScope struct {
	*Node
}

func (n StartConditionsScope) StartConditions() StartConditions {
	child := n.Child(selector.StartConditions)
	return StartConditions{child}
}

func (n StartConditionsScope) LexerPart() []LexerPart {
	nodes := n.Children(selector.LexerPart)
	var ret = make([]LexerPart, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(LexerPart))
	}
	return ret
}

type StateMarker struct {
	*Node
}

func (n StateMarker) Name() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

type Stateref struct {
	*Node
}

func (n Stateref) Name() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

type StringLiteral struct {
	*Node
}

type Symref struct {
	*Node
}

func (n Symref) Name() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

func (n Symref) Args() (SymrefArgs, bool) {
	child := n.Child(selector.SymrefArgs)
	return SymrefArgs{child}, child.IsValid()
}

type SymrefArgs struct {
	*Node
}

func (n SymrefArgs) ArgList() []Argument {
	nodes := n.Children(selector.Argument)
	var ret = make([]Argument, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(Argument))
	}
	return ret
}

type SyntaxProblem struct {
	*Node
}

type TemplateParam struct {
	*Node
}

func (n TemplateParam) Modifier() (ParamModifier, bool) {
	child := n.Child(selector.ParamModifier)
	return ParamModifier{child}, child.IsValid()
}

func (n TemplateParam) ParamType() ParamType {
	child := n.Child(selector.ParamType)
	return ParamType{child}
}

func (n TemplateParam) Name() Identifier {
	child := n.Child(selector.Identifier)
	return Identifier{child}
}

func (n TemplateParam) ParamValue() (ParamValue, bool) {
	child := n.Child(selector.ParamValue)
	return ToTmNode(child).(ParamValue), child.IsValid()
}

type InvalidToken struct {
	*Node
}

type MultilineComment struct {
	*Node
}

type Comment struct {
	*Node
}

type Templates struct {
	*Node
}
