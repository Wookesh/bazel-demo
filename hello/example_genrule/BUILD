load("@io_bazel_rules_go//go:def.bzl", "go_binary")

genrule(
    name = "list",
    srcs = [],
    outs = ["list.txt"],
    cmd = "./$(location generator) > \"$@\"",
    tools = [":generator"],
)

go_binary(
    name = "generator",
    srcs = ["main.go"],
)

go_binary(
    name = "consumer",
    srcs = ["consumer.go"],
    data = [":list"],
)
