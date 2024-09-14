// generated by Textmapper; DO NOT EDIT

package tm

import (
	"github.com/inspirer/textmapper/parsers/tm/token"
)

const tmNumClasses = 36

var tmRuneClass = []uint8{
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 1, 1, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 21, 22, 23, 24, 25, 26, 27, 28, 28,
	28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28,
	28, 28, 28, 28, 28, 29, 30, 31, 1, 28, 1, 28, 28, 28, 28, 28, 28, 28, 28, 28,
	28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 32, 33,
	34, 35,
}

const tmRuneClassLen = 127
const tmFirstRule = -4

var tmStateMap = []int{
	0, 0, 50, 60,
}

var tmToken = []token.Type{
	1, 0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
	80, 81, 82,
}

var tmLexerAction = []int8{
	-4, -4, 49, 49, 49, 47, 44, 43, 42, 40, 38, 35, 32, 31, 30, 28, 27, 25, 24,
	20, 19, 17, 16, 15, 13, 12, 11, 10, 8, 7, -4, 6, 5, 3, 2, 1, -37, -37, -37,
	-37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37,
	-37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37,
	-37, -37, -37, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29,
	-29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29,
	-29, -29, -29, -29, -29, -29, -29, -29, -29, -14, -14, -14, -14, -14, -14,
	-14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14,
	-14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, 4, -14, -14, -15,
	-15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15,
	-15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15,
	-15, -15, -15, -15, -15, -84, -84, -84, -84, -84, -84, -84, -84, -84, -84,
	-84, -84, -84, -84, -84, -84, -84, -84, -84, -84, -84, -84, -84, -84, -84,
	-84, -84, -84, -84, -84, -84, -84, -84, -84, -84, -84, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23,
	-23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23,
	-23, -23, -23, -23, -23, -23, -23, -23, -44, -44, -44, -44, -44, -44, -44,
	-44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -1, -44, -44, 8, -44, -44,
	-44, -44, -44, -44, -44, 8, -44, -44, -44, -44, -44, -44, -44, -4, -4, -4,
	-4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, 9, -4, -4, 8, -4, -4,
	-4, -4, -4, -4, -4, 8, -4, -4, -4, -4, -4, -4, -4, -41, -41, -41, -41, -41,
	-41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41,
	-41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41,
	-41, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35,
	-35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35,
	-35, -35, -35, -35, -35, -35, -35, -31, -31, -31, -31, -31, -31, -31, -31,
	-31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31,
	-31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -16, -16,
	-16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16,
	-16, -16, -16, -16, -16, -16, -16, 14, -16, -16, -16, -16, -16, -16, -16,
	-16, -16, -16, -16, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17,
	-17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17,
	-17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -30, -30, -30, -30, -30,
	-30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30,
	-30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30,
	-30, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19,
	-19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19,
	-19, -19, -19, -19, -19, -19, -19, -22, -22, -22, -22, -22, -22, -22, -22,
	-22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, 18, -22,
	-22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, 19, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42,
	-42, -2, -42, -42, -42, -42, 43, -42, -42, -42, -42, -42, -42, -42, -42, -42,
	-42, -42, -42, -42, -42, -42, -42, -4, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 22, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, -4, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 22, 21, 21, 21, 21, 23, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -20, -20, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20,
	-20, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4,
	-4, -4, 19, -4, -4, -4, -4, 26, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21,
	-21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21,
	-21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -33, -33, -33, -33,
	-33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33,
	-33, -33, -33, -33, -33, 29, -33, -33, -33, -33, -33, -33, -33, -33, -33,
	-33, -33, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34,
	-34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34,
	-34, -34, -34, -34, -34, -34, -34, -34, -32, -32, -32, -32, -32, -32, -32,
	-32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32,
	-32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -28,
	-28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28,
	-28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28,
	-28, -28, -28, -28, -28, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25,
	-25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25,
	-25, -3, -25, -25, -25, -25, -25, -25, -25, -25, -25, -4, -4, -4, -4, -4, -4,
	-4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, 34,
	-4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -26, -26, -26, -26, -26, -26,
	-26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26,
	-26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26,
	-4, 35, 35, -4, 35, 35, 35, 35, 35, 35, 35, 37, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 36, 35, 35, 35, 35, 35, -4, 35,
	35, -4, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, -45, -45, -45,
	-45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45,
	-45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45,
	-45, -45, -45, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, 39, -38,
	-38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38,
	-38, -38, -38, -38, -38, -38, -38, -38, -38, -39, -39, -39, -39, -39, -39,
	-39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39,
	-39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, 41, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40,
	-40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40,
	-40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -10, 43, 43, -10, -10, 43,
	43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43,
	43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, -4, 44, 44, -4, 44, 44, 46, 44,
	44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44,
	44, 44, 44, 45, 44, 44, 44, 44, 44, -4, 44, 44, -4, 44, 44, 44, 44, 44, 44,
	44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44,
	44, 44, 44, 44, 44, 44, 44, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36,
	-36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, 48, -36,
	-36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -18, -18, -18, -18, -18,
	-18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18,
	-18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18,
	-18, -9, -9, 49, 49, 49, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -4,
	-4, 49, 49, 49, 47, 44, 43, 42, 40, 38, 35, 32, 31, 30, 28, 27, 25, 24, 51,
	19, 17, 16, 15, 13, 12, 11, 10, 8, 7, -4, 6, 5, 3, 2, 1, -4, 55, 55, -4, -4,
	55, 55, 55, 55, 55, 55, 55, 55, 55, 21, 55, 55, 55, 55, 43, 55, 55, 55, 55,
	55, 55, 55, 55, 55, 53, 52, 55, 55, 55, 55, 55, -4, 55, 55, -4, 55, 55, 55,
	55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55,
	55, 55, 55, 55, 55, 55, 55, 55, 55, 55, -4, 53, 53, -4, -4, 53, 53, 53, 53,
	53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53,
	53, 53, 54, 55, 53, 53, 53, 53, -4, 53, 53, -4, 53, 53, 53, 53, 53, 53, 53,
	53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53,
	53, 53, 53, 53, 53, 53, -4, 55, 55, -4, -4, 55, 55, 55, 55, 55, 55, 55, 55,
	55, 55, 55, 55, 55, 55, 59, 55, 55, 55, 55, 55, 55, 55, 55, 55, 57, 56, 55,
	55, 55, 55, 55, -4, 55, 55, -4, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55,
	55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55,
	55, 55, -4, 57, 57, -4, -4, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57,
	57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 58, 55, 57, 57, 57, 57,
	-4, 57, 57, -4, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57,
	57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, -85, -85,
	-85, -85, -85, -85, -85, -85, -85, -85, -85, -85, -85, -85, -85, -85, -85,
	-85, -85, -85, -85, -85, -85, -85, -85, -85, -85, -85, -85, -85, -85, -85,
	-85, -85, -85, -85, -4, -4, 49, 49, 49, 47, 44, 43, 42, 40, 38, 35, 32, 31,
	30, 28, 27, 25, 24, 20, 19, 17, 16, 15, 13, 12, 11, 10, 8, 7, -4, 6, 61, 3,
	2, 1, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43,
	-43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43,
	-43, -43, -43, -43, -43, -43, -43,
}

var tmBacktracking = []int{
	40, 9, // in ID
	38, 21, // in '/'
	21, 33, // in '('
}
