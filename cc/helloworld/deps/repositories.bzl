load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Define all external repositories
def repositories():
    # =============================
    # Bazel Skylib
    # =============================
    http_archive(
        name = "bazel_skylib",
        sha256 = "74d544d96f4a5bb630d465ca8bbcfe231e3594e5aae57e1edbf17a6eb3ca2506",
        url = "https://github.com/bazelbuild/bazel-skylib/releases/download/1.3.0/bazel-skylib-1.3.0.tar.gz",
    )

    # =============================
    # gtest / gmock
    # =============================
    http_archive(
        name = "com_google_googletest",
        sha256 = "81964fe578e9bd7c94dfdb09c8e4d6e6759e19967e397dbea48d1c10e45d0df2",
        strip_prefix = "googletest-release-1.12.1",
        urls = ["https://github.com/google/googletest/archive/refs/tags/release-1.12.1.tar.gz"],
    )

    # =============================
    # gRPC v1.72.2 + Protobuf v30.0/v4.30.0/v6.30.0
    # =============================
    http_archive(
        name = "com_github_grpc_grpc",
        urls = [
            "https://github.com/grpc/grpc/archive/d6b42b2fd3e522a60318c7e2f2021219c6b5c470.tar.gz",
        ],
        strip_prefix = "grpc-d6b42b2fd3e522a60318c7e2f2021219c6b5c470",
    )
