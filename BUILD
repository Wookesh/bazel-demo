load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/wookesh/bazel-demo
# gazelle:go_generate_proto false
# gazelle:proto disable
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)
