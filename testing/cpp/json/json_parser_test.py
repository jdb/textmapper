"""Python unit tests for the JSON parser, mirroring the C++ tests."""

import unittest
from nbjson import Lexer, Parser, NodeFlags, Location, Token
from nbjson import ast
import markup


class ParserRangeTest(unittest.TestCase):
    def test_instantiate(self):
        """Test that parser can be instantiated."""
        def listener(node, flags, loc):
            pass
        def always_recover(status):
            return True
        parser = Parser(byteLocation=False, listener=listener, error_handler=always_recover, x=0, y=False)
        self.assertIsNotNone(parser)

    def test_parse_with_lexer(self):
        """Test parsing with a lexer."""
        output = []
        def listener(node, flags, loc):
            output.append((node, loc.begin, loc.end))
        def always_recover(status):
            return True
        parser = Parser(byteLocation=False, listener=listener, error_handler=always_recover, x=8, y=False)
        lexer = Lexer("1.0")
        status = parser.parse(lexer)
        self.assertTrue(status.ok, f"Parse failed: {status.message}")
        self.assertEqual(output, [
            (ast.JSONValue, 0, 3),
            (ast.JSONText, 0, 3),
        ])

    def test_node_types(self):
        """Test parsing different node types."""
        tests = [
            (ast.JSONEmptyObject, [
                R"«{}»",
                R"«{ /* comment */ }»",
                R'{"aa": «{}» }',
            ]),
            (ast.JSONObject, [
                R'«{ "a" : "b" }»',
                R'«{ "a" : ["b"] }»',
                R'«{ "a" : {} }»',
                R'«{ "a" : «{"q":B}» }»',
            ]),
            (ast.JSONArray, [
                R'{ "a" : «["b"]» }',
                R' «[]» ',
            ]),
            (ast.JSONText, [
                R'«{ "a" : ["b", A] }»',
                R' «"aa"» ',
                R' «A» ',
            ]),
            (ast.JSONMember, [
                R'[{ «"a" : ["b"]», «"q":[]» }]',
            ]),
            (ast.JSONValue, [
                R'«{ "a" : «[«"b"»]» }»',
                R' «"aa"» ',
            ]),
            (ast.NonExistingType, []),
            (ast.MultiLineComment, [
                R'{ "a"«/* abc */» : [] }',
            ]),
            (ast.JSONString, [
                R'{ «"a"» : [«"b"»] }',
            ]),
            (ast.SyntaxProblem, [
                R'{ «"a" "b"» }',
            ]),
        ]
        seen = set()
        for node_type, inputs in tests:
            for input_str in inputs:
                want_ranges, text = markup.parse(input_str)
                lexer = Lexer(text)
                got = []
                def listener(node, flags, loc):
                    if node == node_type:
                        seen.add(node)
                        got.append((loc.begin, loc.end))
                def always_recover(status):
                    return True
                parser = Parser(byteLocation=False, listener=listener, error_handler=always_recover, x=9, y=True)
                status = parser.parse(lexer)
                self.assertTrue(status.ok, f"Parse failed for {node_type} on input: {input_str}: {status.message}")
                self.assertEqual(sorted(got), sorted(want_ranges),
                               f"Node type test failed for {node_type} on input:\n{input_str}")


if __name__ == "__main__":
    unittest.main()

