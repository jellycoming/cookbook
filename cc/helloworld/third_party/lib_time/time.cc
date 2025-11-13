#include <iostream>
#include <ctime>
#include <iomanip>

// extern "C"，防止 C++ name mangling 导致找不到符号
extern "C" void current_time() {
    std::time_t t = std::time(nullptr);
    std::tm* tm_ptr = std::localtime(&t);
    std::cout << std::put_time(tm_ptr, "%Y-%m-%d %H:%M:%S") << std::endl;
}
