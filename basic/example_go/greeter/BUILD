load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "greeter",
    srcs = ["greeter.go"],
    importpath = "github.com/wookesh/bazel-demo/basic/example_go/greeter",
    visibility = ["//visibility:public"],
)

go_test(
    name = "greeter_test",
    srcs = ["greeter_test.go"],
    embed = [":greeter"],
    importpath = "basic/example_go/greeter",
)
