load("@rules_python//python:defs.bzl", "py_library")
load("@django_env//:requirements.bzl", "requirement")

package(default_visibility = ["//hello/example_python:__subpackages__"])

py_library(
    name = "project",
    srcs = glob(["*.py"]),  # all python files in directory
    deps = [
        requirement("django"),
    ],
)
