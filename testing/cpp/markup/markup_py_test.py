"""Python unit tests for markup functions, replicating the C++ tests."""

import unittest

from markup import create, parse


class MarkupTest(unittest.TestCase):
    def test_create(self):
        """Test Create function."""
        self.assertEqual(create("foo", []), "foo")
        self.assertEqual(create("foo", [(1, 2)]), "f«o»o")
        self.assertEqual(create("foo", [(1, 2), (1, 3)]), "f««o»o»")
        self.assertEqual(create("foo", [(0, 3)]), "«foo»")

    def test_parse(self):
        """Test Parse function."""
        self.assertEqual(parse("f«o»o"), ([(1, 2)], "foo"))
        self.assertEqual(parse("f««o»»o"), ([(1, 2), (1, 2)], "foo"))
        self.assertEqual(parse("f««o»o»"), ([(1, 2), (1, 3)], "foo"))
        self.assertEqual(parse("f«中»o"), ([(1, 2)], "f中o"))
        self.assertEqual(parse("f«🎉»o"), ([(1, 2)], "f🎉o"))


if __name__ == "__main__":
    unittest.main()

