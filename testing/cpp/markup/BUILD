package(default_visibility = ["//visibility:public"])

cc_library(
    name = "markup",
    srcs = ["markup.cc"],
    hdrs = ["markup.h"],
    copts = ["-std=c++17"],
    # Make headers available as cpp/markup/markup.h
    include_prefix = "cpp/markup",
    deps = [
        "@abseil-cpp//absl/log:check",
        "@abseil-cpp//absl/log:log",
        "@abseil-cpp//absl/strings",
    ],
)

cc_test(
    name = "markup_test",
    size = "small",
    srcs = ["markup_test.cc"],
    copts = ["-std=c++17"],
    deps = [
        ":markup",
        "@googletest//:gtest_main",
    ],
)

# Python markup module
load("@rules_python//python:defs.bzl", "py_library", "py_test")

py_library(
    name = "markup_py",
    srcs = ["markup.py"],
    imports = ["."],  # Make markup.py importable as 'markup'
    visibility = ["//visibility:public"],
)

# Python test for markup functions
py_test(
    name = "markup_py_test",
    srcs = ["markup_py_test.py"],
    deps = [
        ":markup_py",
    ],
    python_version = "PY3",
    size = "small",
    imports = ["."],  # Add current directory to Python path
)

# Test suite to run all markup tests (C++ and Python)
test_suite(
    name = "all_tests",
    tests = [
        ":markup_test",      # C++ tests
        ":markup_py_test",   # Python tests
    ],
    visibility = ["//visibility:public"],
)
