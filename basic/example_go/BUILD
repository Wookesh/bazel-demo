load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "hello",
    embed = [":example_go_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "example_go_lib",
    srcs = ["main.go"],
    importpath = "github.com/wookesh/bazel-demo/basic/example_go",
    visibility = ["//visibility:private"],
    deps = [
        "//basic/example_go/greeter",
        "@com_github_sirupsen_logrus//:logrus",
    ],
)
