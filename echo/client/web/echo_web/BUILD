load("@rules_python//python:defs.bzl", "py_library")
load("@django_env//:requirements.bzl", "requirement")

py_library(
    name = "project",
    srcs = [
        "__init__.py",
        "asgi.py",
        "settings.py",
        "urls.py",
        "wsgi.py",
    ],
    visibility = ["//echo/client/web:__pkg__"],
    deps = [
        requirement("django"),
    ],
)
