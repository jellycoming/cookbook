#include <sstream>
#include <ctime>
#include <iomanip>

#include "time_util.h"

std::string current_time() {
    std::time_t t = std::time(nullptr);
    std::tm* tm_ptr = std::localtime(&t);

    std::ostringstream oss;
    oss << std::put_time(tm_ptr, "%Y-%m-%d %H:%M:%S");
    return oss.str();
}
