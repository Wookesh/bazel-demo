load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library")

cc_library(
    name = "greeter",
    srcs = ["greeter.cpp"],
    hdrs = ["greeter.h"],
)

cc_binary(
    name = "hello",
    srcs = ["main.cpp"],
    deps = [":greeter"],
)
