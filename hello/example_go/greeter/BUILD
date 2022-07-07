load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "greeter",
    srcs = ["greeter.go"],
    importpath = "hello/example_go/greeter",
)

go_test(
    name = "greeter_test",
    srcs = ["greeter_test.go"],
    embed = [":greeter"],
    importpath = "hello/example_go/greeter",
)
