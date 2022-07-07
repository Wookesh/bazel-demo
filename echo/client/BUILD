load("@rules_python//python:defs.bzl", "py_binary")
load("@django_env//:requirements.bzl", "requirement")

py_binary(
    name = "echo_client",
    srcs = [":echo_client.py"],
    deps = [
        #        requirement("lorem"),
        "//echo/proto:echo_python_grpc",
        "//echo/proto:message_python_proto",
    ],
)
