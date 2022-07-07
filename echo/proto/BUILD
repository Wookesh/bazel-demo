load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_proto_grpc//go:defs.bzl", "go_grpc_library", "go_proto_library")
load("@rules_proto_grpc//python:defs.bzl", "python_grpc_library", "python_proto_library")

#load("@rules_proto_grpc//js:defs.bzl", "js_proto_compile")

proto_library(
    name = "message_proto",
    srcs = ["message.proto"],
)

go_proto_library(
    name = "message_go_proto",
    importpath = "echo/proto",
    protos = [":message_proto"],
    visibility = ["//visibility:public"],
)

python_grpc_library(
    name = "message_python_proto",
    protos = [":message_proto"],
    visibility = ["//visibility:public"],
)

#js_proto_compile(
#    name = "message_js_proto",
#    protos = [":message_proto"],
#)

proto_library(
    name = "echo_grpc",
    srcs = ["echo.proto"],
    deps = [":message_proto"],
)

go_grpc_library(
    name = "echo_go_grpc",
    importpath = "echo/proto",
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
