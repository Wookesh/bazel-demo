load("@io_bazel_rules_go//go:def.bzl", "go_test")

test_suite(
    name = "non-flaky",
    tags = ["-flaky"],
    tests = [
        "flaky",
        "long",
    ],
)

go_test(
    name = "flaky",
    size = "small",
    srcs = ["flaky_test.go"],
    flaky = True,
)

go_test(
    name = "long",
    srcs = ["long_test.go"],
)
