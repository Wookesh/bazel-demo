load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github_com_wookesh_bazel_demo
# gazelle:go_generate_proto false
# gazelle:proto disable
gazelle(
    name = "gazelle",
    prefix = "github_com_wookesh_bazel_demo",
)

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)
