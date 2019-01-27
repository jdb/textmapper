// generated by Textmapper; DO NOT EDIT

package selector

import (
	"github.com/inspirer/textmapper/tm-go/parsers/test"
)

type Selector func(nt test.NodeType) bool

var (
	Any               = func(t test.NodeType) bool { return true }
	Block             = func(t test.NodeType) bool { return t == test.Block }
	Decl1             = func(t test.NodeType) bool { return t == test.Decl1 }
	Decl2             = func(t test.NodeType) bool { return t == test.Decl2 }
	Int               = func(t test.NodeType) bool { return t == test.Int }
	Negation          = func(t test.NodeType) bool { return t == test.Negation }
	Test              = func(t test.NodeType) bool { return t == test.Test }
	MultiLineComment  = func(t test.NodeType) bool { return t == test.MultiLineComment }
	SingleLineComment = func(t test.NodeType) bool { return t == test.SingleLineComment }
	InvalidToken      = func(t test.NodeType) bool { return t == test.InvalidToken }
	Identifier        = func(t test.NodeType) bool { return t == test.Identifier }
	Declaration       = OneOf(test.Declaration...)
	TokenSet          = OneOf(test.TokenSet...)
)

func OneOf(types ...test.NodeType) Selector {
	if len(types) == 0 {
		return func(test.NodeType) bool { return false }
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
	return func(t test.NodeType) bool {
		i := uint(t) / bits
		return int(i) < len(bitarr) && bitarr[i]&(1<<(uint(t)%bits)) != 0
	}
}
