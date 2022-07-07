load("@rules_python//python:defs.bzl", "py_binary")
load("@demo_env//:requirements.bzl", "requirement")

py_binary(
    name = "demo",
    srcs = ["demo.py"],
    deps = [requirement("lorem")],
)
