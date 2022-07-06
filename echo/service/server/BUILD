load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "server",
    srcs = ["echo_server.go"],
    importpath = "echo/service/server",
    visibility = ["//echo:__subpackages__"],
    deps = ["//echo/proto:echo_go_grpc"],
)
