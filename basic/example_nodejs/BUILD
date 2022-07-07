load("@build_bazel_rules_nodejs//:index.bzl", "js_library", "nodejs_binary")

nodejs_binary(
    name = "example_nodejs",
    data = [
        ":index",
    ],
    entry_point = ":index.js",
)

js_library(
    name = "index",
    srcs = ["index.js"],
)
