load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@k8s_deploy//:defaults.bzl", "k8s_deploy")

go_binary(
    name = "echo",
    embed = [":service_lib"],
    visibility = ["//visibility:public"],
)

go_image(
    name = "echo_image",
    srcs = ["main.go"],
    deps = [
        "//echo/proto:echo_go_grpc",
        "//echo/service/server",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

k8s_deploy(
    name = "echo_deploy",
    images = {
        "echo_image:latest": ":echo_image",
    },
    template = "deployment.yaml",
)

go_library(
    name = "service_lib",
    srcs = ["main.go"],
    importpath = "github.com/wookesh/bazel-demo/echo/service",
    visibility = ["//visibility:private"],
    deps = [
        "//echo/proto:echo_go_grpc",
        "//echo/service/server",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
