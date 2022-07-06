load("@rules_python//python:defs.bzl", "py_binary")
load("@django//:requirements.bzl", "requirement")

py_binary(
    name = "echo_client",
    srcs = [":echo_client.py"],
    deps = [
        requirement("lorem"),
    ],
)
