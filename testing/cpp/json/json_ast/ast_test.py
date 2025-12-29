"""Python unit tests for the JSON AST, mirroring the parser tests."""

import unittest
from nbjson import parse, Tree
from nbjson import ast
from nbjson import selector
import markup

class AstTest(unittest.TestCase):    
    def test_parse_simple_value(self):
        """Test parsing 1.0, check root is JSONText and JSONValue child exists."""
        tree = parse("1.0")
        root = tree.root
        
        self.assertIsNotNone(root)
        self.assertEqual(root.type, ast.JSONText)
        
        json_value = root.child(selector.JSONValue)
        self.assertIsNotNone(json_value)
        self.assertEqual(json_value.type, ast.JSONValue)

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
        for node_type, srcs in tests:
            for src in srcs:
                want_ranges, stripped = markup.parse(src)
                tree = parse(stripped)
                got = []
                def collect_node(n):
                    if n.type == node_type:
                        seen.add(n.type)
                        got.append((n.offset(), n.endoffset()))
                    return True  # Continue traversal
                if tree.root:
                    tree.root.traverse(collect_node)
                self.assertEqual(sorted(got), sorted(want_ranges),
                               f"Node type test failed for {node_type} on input:\n{src}")


if __name__ == "__main__":
    unittest.main()
