"""Python unit tests for the JSON lexer, mirroring the C++ tests."""

import unittest
from nbjson import Lexer, Token
import markup


class LexerTest(unittest.TestCase):
    def test_token(self):
        """Unit tests for each token type."""
        tests = [
            ("id", Token.ID, [R' 0«foo» «barB1»']),
            ("string", Token.JSONSTRING, [R'«"foo"» «"b\nar"» «"α"»']),
            ("number", Token.JSONNUMBER, ["«1» «534»", "«1e9» «1.2» «1e-2»"]),
            ("true", Token.TRUE, ["«true»", "/* true */ «true» "]),
            ("false", Token.FALSE, ["  «false» "]),
            ("null", Token.KW_NULL, ["  «null» "]),
            ("lbrace", Token.LBRACE, ["«{»"]),
            ("rbrace", Token.RBRACE, ["«}»"]),
            ("lbrack", Token.LBRACK, ["«[»"]),
            ("rbrack", Token.RBRACK, ["«]»"]),
            ("colon", Token.COLON, ["«:»"]),
            ("comma", Token.COMMA, ["«,»"]),
            ("comment", Token.MULTILINECOMMENT, ["  «/*  asda *** */» bar"]),
            ("char_a", Token.CHAR_A, ["  «A» «α» «A»"]),
        ]
        for name, expected_token, test_in in tests:
            with self.subTest():
                for test_case in test_in:
                    with self.subTest():
                        want_ranges, source = markup.parse(test_case)                        
                        lexer = Lexer(source)
                        got = []
                        while (tok := lexer.next()) != Token.EOI:
                            if tok == expected_token:
                                got.append((lexer.begin, lexer.end))
                        self.assertEqual(got, want_ranges,
                                       f"{name}: want {test_case}, got {markup.create(source, got)}")

    def test_multi_byte_utf8(self):
        """Test that offsets by characters correctly handle multi-byte UTF-8."""
        tests = [
            (R'"α"', 3),
            (R'"中"', 3),
            (R'"🎉"', 3),
            (R'"aα中🎉"', 6),
        ]
        for test_in, want in tests:
            with self.subTest():
                lexer = Lexer(test_in)
                token = lexer.next()
                self.assertEqual(token, Token.JSONSTRING, f"want type Token.JSONSTRING, got {token}")
                got = lexer.end - lexer.begin
                self.assertEqual(want, got, f"lexer.end - lexer.begin = {got}, want {want}")
         

if __name__ == "__main__":
    unittest.main()

