package(default_visibility = ["//visibility:public"])

cc_import(
    name = "mytime_lib",
    shared_library = "lib/libtime_util.so",
    hdrs = glob([
        "include/**/*.h",
    ]),
    visibility = ["//visibility:private"],
)

# cc_import不支持includes参数, 用cc_library包装一下
cc_library(
    name = "mytime",
    includes = ["include"],
    deps = [
        ":mytime_lib",
    ],
)

#################################################
# 这种方式也可生效
cc_library(
    name = "mytime_l",
    srcs = ["lib/libtime_util.so"],
    hdrs = glob([
        "include/**/*.h",
    ]),
    includes = ["include"],
    linkstatic = False,
)
