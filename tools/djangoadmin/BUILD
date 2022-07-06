load("@rules_python//python:defs.bzl", "py_binary", "py_library")
load("@django_env//:requirements.bzl", "requirement")
#load("@boto3//:requirements.bzl", "requirement")

py_library(
    name = "__init__",
    srcs = ["__init__.py"],
)

py_binary(
    name = "django-admin",
    srcs = ["django-admin.py"],
    visibility = ["//visibility:public"],
    deps = [
        requirement("django"),
    ],
)

py_binary(
    name = "list",
    srcs = ["list.py"],
    python_version = "PY3",
    deps = [
        requirement("django"),
    ],
)
