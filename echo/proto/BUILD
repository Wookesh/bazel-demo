load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_proto_grpc//go:defs.bzl", "go_grpc_library", "go_proto_library")
load("@rules_proto_grpc//js:defs.bzl", "js_grpc_web_library", "js_proto_library")
load("@rules_proto_grpc//python:defs.bzl", "python_grpc_library", "python_proto_library")

proto_library(
    name = "message_proto",
    srcs = ["message.proto"],
)

go_proto_library(
    name = "message_go_proto",
    importpath = "github.com/wookesh/bazel-demo/echo/proto",
    protos = [":message_proto"],
    visibility = ["//visibility:public"],
)

python_proto_library(
    name = "message_python_proto",
    protos = [":message_proto"],
    visibility = ["//visibility:public"],
)

js_proto_library(
    name = "message_js_proto",
    protos = [":message_proto"],
    visibility = ["//visibility:public"],
)

proto_library(
    name = "echo_grpc",
    srcs = ["echo.proto"],
    deps = [":message_proto"],
)

go_grpc_library(
    name = "echo_go_grpc",
    importpath = "github.com/wookesh/bazel-demo/echo/proto",
    protos = [
        ":echo_grpc",
        ":message_proto",
    ],
    visibility = ["//visibility:public"],
)

python_grpc_library(
    name = "echo_python_grpc",
    protos = [
        ":echo_grpc",
        ":message_proto",
    ],
    visibility = ["//visibility:public"],
)

js_grpc_web_library(
    name = "echo_js_grpc",
    protos = [
        ":echo_grpc",
        ":message_proto",
    ],
    visibility = ["//visibility:public"],
)
