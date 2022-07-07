load("@rules_python//python:defs.bzl", "py_binary")
load("@django_env//:requirements.bzl", "requirement")

py_binary(
    name = "manage",
    srcs = [":manage.py"],
    deps = [
        "//basic/example_python/echo_web:project",
        requirement("django"),
    ],
)
