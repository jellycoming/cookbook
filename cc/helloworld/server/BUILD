# 服务端可执行文件
cc_binary(
    name = "greeter_grpc_server",
    srcs = ["greeter_grpc_server.cc"],
    deps = [
        "@com_github_grpc_grpc//:grpc++",
        "@com_github_grpc_grpc//:grpc++_reflection",
        "//protos:helloworld_cc_grpc",
        "//third_party/lib_time:lib_time",
    ],
)

# 客户端可执行文件
cc_binary(
    name = "greeter_grpc_client",
    srcs = ["greeter_grpc_client.cc"],
    deps = [
        "@com_github_grpc_grpc//:grpc++",
         "//protos:helloworld_cc_grpc",
    ],
)
