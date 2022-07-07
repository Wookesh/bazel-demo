load("@rules_python//python:defs.bzl", "py_binary")
load("@demo_env//:requirements.bzl", "requirement")

py_binary(
    name = "demo",
    srcs = ["demo.py"],
    python_version = "PY3",
    deps = [requirement("lorem")],
)
