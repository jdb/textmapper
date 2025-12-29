package(default_visibility = ["//visibility:public"])

cc_library(
    name = "lexer",
    srcs = ["json_lexer.cc"],
    hdrs = [
        "json_lexer.h",
        "json_token.h",
    ],
    deps = [
        "@abseil-cpp//absl/log",
        "@abseil-cpp//absl/strings",
    ],
)

cc_test(
    name = "lexer_test",
    size = "small",
    srcs = ["json_lexer_test.cc"],
    deps = [
        ":lexer",
        "//cpp/markup",
        "@googletest//:gtest_main",
    ],
)

cc_library(
    name = "parser",
    srcs = ["json_parser.cc"],
    hdrs = ["json_parser.h"],
    deps = [
        ":lexer",
        "@abseil-cpp//absl/functional:function_ref",
        "@abseil-cpp//absl/log",
        "@abseil-cpp//absl/status",
        "@abseil-cpp//absl/strings",
    ],
)

cc_library(
    name = "ast",
    deps = [
        ":lexer",
        ":parser",
        "//cpp/json/json_ast:parse",
        "//cpp/json/json_ast:selector",
        "@abseil-cpp//absl/strings",
    ],
)

cc_test(
    name = "parser_test",
    size = "small",
    srcs = ["json_parser_test.cc"],
    deps = [
        ":lexer",
        ":parser",
        "//cpp/markup",
        "@googletest//:gtest_main",
    ],
)

cc_binary(
    name = "benchmark",
    srcs = ["benchmark.cc"],
    deps = [
        ":lexer",
        "@google_benchmark//:benchmark",
    ],
)

# Python bindings using nanobind
load("@nanobind_bazel//:build_defs.bzl", "nanobind_extension")

nanobind_extension(
    name = "nblexer",
    srcs = ["json_py_binding.cpp"],
    deps = [
        ":ast",
        ":lexer",
        ":parser",
    ],
    copts = ["-std=c++17"],
    visibility = ["//visibility:public"],
)


pkg_init = """
from .nblexer import Lexer, Parser, Token, State, NodeFlags, Location, Tree, Node, parse
from . import nblexer
ast = nblexer.ast
selector = nblexer.selector

def traverse(node, callback):
    \"\"\"Traverse the AST tree, calling callback on each node. Returns False to skip subtree.\"\"\"
    if not node or callback(node) is False:
        return
    for child in node.children(selector.Any):
        traverse(child, callback)

Node.traverse = traverse

__all__ = [
    "Lexer",
    "Parser",
    "Token",
    "State",
    "NodeFlags",
    "Tree",
    "Node",
    "parse",
    "ast",
    "selector",
]
"""

# Prepare files for nbjson package
genrule(
    name = "nbjson_init_py",
    srcs = [],
    outs = ["__init__.py"],
    cmd = "cat > $@ <<EOF\n%s\nEOF" % pkg_init,
)

load("@rules_python//python:packaging.bzl", "py_wheel")
load("@rules_python//python:defs.bzl", "py_test", "py_library")

# Create a py_library that represents the installed wheel structure
# This mimics what gets installed when the wheel is installed via pip
# The wheel installs to platlib/nbjson/, so we create nbjson/ directory structure
genrule(
    name = "nbjson_package",
    srcs = [
        ":nbjson_init_py",
        ":nblexer.abi3.so",
    ],
    outs = [
        "nbjson/__init__.py",
        "nbjson/nblexer.abi3.so",
    ],
    cmd = """
        mkdir -p $$(dirname $(location nbjson/__init__.py))
        cp $(location :nbjson_init_py) $(location nbjson/__init__.py)
        cp $(location :nblexer.abi3.so) $(location nbjson/nblexer.abi3.so)
    """,
)

py_library(
    name = "nbjson_wheel_lib",
    srcs = ["nbjson/__init__.py"],
    data = ["nbjson/nblexer.abi3.so"],
    imports = ["."],
    # This creates a structure like: nbjson/__init__.py and nbjson/nblexer.abi3.so
    # which matches the wheel's platlib/nbjson/ structure when installed
)

py_wheel(
    name = "nblexer_wheel",
    distribution = "json-lexer",
    version = "1.0.0",
    deps = [],  # No deps to avoid duplication
    python_tag = "cp311",
    abi = "abi3",
    platform = "manylinux2014_x86_64",
    data_files = {
        "__init__.py": "platlib/nbjson/__init__.py",
        "nblexer.abi3.so": "platlib/nbjson/nblexer.abi3.so",
    },
)

# Python test for the lexer
# This test uses the wheel structure (nbjson package) to match the installed wheel
py_test(
    name = "json_lexer_test",
    srcs = ["json_lexer_test.py"],
    deps = [
        ":nbjson_wheel_lib",  # Use the wheel structure instead of direct nblexer
        "//cpp/markup:markup_py",  # Python markup module
        "@rules_python//python/runfiles",  # For runfiles access
    ],
    python_version = "PY3",
    size = "small",
    imports = [".."],  # Add cpp/ to path so markup module is accessible
)

# Python test for the parser
py_test(
    name = "json_parser_test",
    srcs = ["json_parser_test.py"],
    deps = [
        ":nbjson_wheel_lib",  # Use the wheel structure instead of direct nblexer
        "//cpp/markup:markup_py",  # Python markup module
        "@rules_python//python/runfiles",  # For runfiles access
    ],
    python_version = "PY3",
    size = "small",
    imports = [".."],  # Add cpp/ to path so markup module is accessible
)

