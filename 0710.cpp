#include <iostream>
#include <climits>
#include <cstdint>
using namespace std;

int main()
{
    cout << "int min type " << INT_MIN << endl;
    cout << "int 类型最大值: " << INT_MAX << endl;
    cout << "int 类型大小: " << sizeof(int) << " 字节" << endl;

    std::cout << "long 类型最小值: " << LONG_MIN << std::endl;
    std::cout << "long 类型最大值: " << LONG_MAX << std::endl;
    std::cout << "long 类型大小: " << sizeof(long) << " 字节" << std::endl;

    std::int32_t i32 = -2147483648;
    std::int64_t i64 = -9223372036854775808LL;
    std::cout << "int32_t 类型值: " << i32 << std::endl;
    std::cout << "int64_t 类型值: " << i64 << std::endl;
    std::cout << "int32_t 类型大小: " << sizeof(std::int32_t) << " 字节" << std::endl;
    std::cout << "int64_t 类型大小: " << sizeof(std::int64_t) << " 字节" << std::endl;

    return 0;
}