"Client module"
load("@npm//:defs.bzl", "npm_link_all_packages")
load("@aspect_bazel_lib//lib:copy_to_bin.bzl", "copy_to_bin")
load("@npm//:vite/package_json.bzl", vite_bin = "bin")
# load("@npm//:vitest/package_json.bzl", vitest_bin = "bin")

npm_link_all_packages(name = "node_modules")

CONFIGS = [
    "vite.config.ts",
    "vitest.config.ts",
    "tsconfig.json",
    "tsconfig.node.json",
]

BUILD_DEPS = [":node_modules/" + d for d in [
    # "@bazel-example/vue-library",
    # "@vitejs/plugin-vue",
    # "@vitejs/plugin-vue-jsx",
    # "@vue/tsconfig",
    "vite",
    # "vue",
    # "vue-router",
]]

copy_to_bin(
    name = "vite_srcs_export",
    srcs = glob([
        "src/**/*.ts",
        "src/**/*.tsx",
        "public/**/*",
    ]),
)

vite_bin.vite(
    name = "client",
    srcs = BUILD_DEPS + CONFIGS + [":vite_srcs_export"],
    args = ["build --config vite.config.ts"],
    chdir = package_name(),
    out_dirs = ["dist"],
    visibility = ["//visibility:public"],
)

# vitest_bin.vitest_test(
#     name = "vitest",
#     args = [
#         "run",
#         "--config=vite.config.ts",
#     ],
#     chdir = package_name(),
#     data = BUILD_DEPS + CONFIGS + [":vite_srcs_export"] + [":node_modules/vitest"] + glob([
#         "src/**/*.test.ts",
#     ]),
# )

# load("@npm//:defs.bzl", "npm_link_all_packages")
# load("@aspect_rules_js//js:defs.bzl", "js_binary")

# npm_link_all_packages(name = "node_modules")

# [platform(
#   name = "linux_{}".format(arch),
#   constraint_values = [
#       "@platforms//os:linux",
#       "@platforms//cpu:{}".format(arch if arch != "amd64" else "x86_64"),
#   ],
# ) for arch in ["amd64", "arm64"]]

# js_binary(
#     name = "binary",
#     entry_point = "main.js",
# )

# load("//:ext.bzl", "oci_binimage_index")

# oci_binimage_index(
#     name = "image",
#     binary = "binary",
#     platforms = [
#         "amd64",
#         "arm64"
#     ],
# )
