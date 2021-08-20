#include <iostream>
// using namespace std;
void HelloWorld();
void Version();

int main() {
    HelloWorld();
    Version();
    return 0;
}

void HelloWorld() {
    std::cout << "Hello World" << std::endl;
}

void Version() {
    std::cout << __cplusplus << "\n";
}