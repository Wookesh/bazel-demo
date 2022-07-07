load("@io_bazel_rules_go//go:def.bzl", "go_binary")
load(":list_generator.bzl", "list_generator")

go_binary(
    name = "generator",
    srcs = ["main.go"],
)

genrule(
    name = "list",
    srcs = [],
    outs = ["list.txt"],
    cmd = "./$(location generator) > \"$@\"",
    tools = [":generator"],
)

list_generator(
    name = "simple-list",
    count = 15,
    visibility = ["//visibility:public"],
)

go_binary(
    name = "consumer",
    srcs = ["consumer.go"],
    data = [":list"],
)
