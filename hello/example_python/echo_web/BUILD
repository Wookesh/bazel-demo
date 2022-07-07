load("@rules_python//python:defs.bzl", "py_library")
load("@django_env//:requirements.bzl", "requirement")

py_library(
    name = "project",
    srcs = glob(["*.py"]),  # all python files in directory
    visibility = ["//hello/example_python:__pkg__"],
    deps = [
        requirement("django"),
    ],
)
