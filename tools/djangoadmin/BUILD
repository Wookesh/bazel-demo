load("@rules_python//python:defs.bzl", "py_binary")
load("@django_env//:requirements.bzl", "requirement")

py_binary(
    name = "django-admin",
    srcs = ["django-admin.py"],
    visibility = ["//visibility:public"],
    deps = [
        requirement("django"),
    ],
)
