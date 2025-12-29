package(default_visibility = ["//visibility:public"])

cc_library(
    name = "selector",
    hdrs = ["selector.h"],
    deps = [
        "//cpp/json:parser",
    ],
)

cc_library(
    name = "parse",
    srcs = ["builder.cc"],
    hdrs = [
        "tree.h",
        "builder.h",
    ],
    deps = [
        "//cpp/json:lexer",
        "//cpp/json:parser",
        ":selector",
        "@abseil-cpp//absl/strings",
    ],
)

cc_test(
    name = "parse_test",
    size = "small",
    srcs = ["ast_test.cc"],
    deps = [
        ":parse",
        "//cpp/markup",
        "@googletest//:gtest_main",
    ],
)

cc_binary(
    name = "display_tree",
    srcs = ["display_tree.cc"],
    deps = [
        ":parse",
        ":selector",
    ],
)

py_test(
    name = "parse_test_py",
    srcs = ["ast_test.py"],
    main = "ast_test.py",
    deps = [
        "//cpp/json:nbjson_wheel_lib",
        "//cpp/markup:markup_py",
        "@rules_python//python/runfiles",
    ],
    python_version = "PY3",
    size = "small",
    imports = ["../.."],  # Adjust to maintain access to markup module
)
