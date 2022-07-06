load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")

go_library(
    name = "echo_server",
    srcs = ["echo_server.go"],
    importpath = "echo/service/server",
    deps = ["//echo/proto:echo_go_grpc"],
)

# TODO: change to go_image
go_binary(
    name = "echo",
    srcs = ["main.go"],
    deps = [":echo_server"],
)
