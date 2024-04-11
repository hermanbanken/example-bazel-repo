# Example of Bazel go monorepos
# https://github.com/jankremlacek/go-bazel/blob/main/proto/servicea/BUILD.bazel
# https://github.com/aspect-build/bazel-examples/blob/main/oci_go_image/BUILD.bazel

load("@aspect_bazel_lib//lib:transitions.bzl", "platform_transition_filegroup")
load("@container_structure_test//:defs.bzl", "container_structure_test")
load("@rules_pkg//pkg:tar.bzl", "pkg_tar")
load("@rules_oci//oci:defs.bzl", "oci_image", "oci_tarball")

pkg_tar(
    name = "tar",
    srcs = ["//project-backend/cmd/server"],
)

oci_image(
    name = "image",
    base = "@distroless_base",
    entrypoint = ["/server"],
    tars = [":tar"],
    visibility = ["//visibility:public"],
)

# This is the target that should be released to the target platform
platform_transition_filegroup(
    name = "transitioned_image",
    srcs = [":image"],
    target_platform = select({
        "@platforms//cpu:arm64": "@rules_go//go/toolchain:linux_arm64",
        "@platforms//cpu:x86_64": "@rules_go//go/toolchain:linux_amd64",
    }),
)

# $ bazel run //server-client:tarball
# $ docker run --rm eu.gcr.io/hermanbanken/example-bazel-monorepo:latest
oci_tarball(
    name = "tarball",
    image = ":image",
    repo_tags = ["eu.gcr.io/hermanbanken/example-bazel-monorepo:latest"],
    visibility = ["//test:__pkg__"],
)

container_structure_test(
    name = "test",
    configs = ["test.yaml"],
    # Use the image built for the exec platform rather than the target platform
    image = ":image",
)

# multi-platform image
# https://github.com/macourteau/aspect-rules_oci/commit/ae2fa02779f1220d8cb81219095e7659485db296
