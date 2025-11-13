# -fPIC：生成位置无关代码（Position Independent Code），是生成 .so 必需的；
# -shared：告诉编译器输出动态链接库；
# -o libtime.so：输出文件名遵循惯例 lib<name>.so。
g++ -fPIC -shared -o libtime.so time.cc
