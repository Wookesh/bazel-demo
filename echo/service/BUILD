load("@io_bazel_rules_go//go:def.bzl", "go_binary")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")

# TODO: change to go_image
go_binary(
    name = "echo",
    srcs = ["main.go"],
    deps = ["//echo/service/server"],
)

go_image(
    name = "echo_image",
    srcs = ["main.go"],
    deps = ["//echo/service/server"],
)
