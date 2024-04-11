# Bazel for n00bs
Created by a Bazel n00b just learning Bazel for the first time.

Preparations:
```bash
brew install bazel
go install github.com/bazelbuild/buildtools/buildozer@latest
```

Generate something useful:
```bash
bazel run //project-backend:tarball
```

Multi-platform Docker OCI images: https://github.com/bazel-contrib/rules_oci/pull/531. New way:
```starlark
# file=BUILD
oci_image_index(
    name = "app",
    image = ":app_linux",
    platforms = [
        "@io_bazel_rules_go//go/toolchain:linux_amd64",
        "@io_bazel_rules_go//go/toolchain:linux_arm64",
    ]
)
```