#include <iostream>
#include <cstdint>
using namespace std;

int main()
{
    uintptr_t address = 0x482040;

    typedef int (*FunctionType)(uint32_t *);

    // 将地址转换为函数指针
    FunctionType function = reinterpret_cast<FunctionType>(address);

    // 现在你可以像调用普通函数一样调用它
    uint32_t arg = 10;
    int result = function(&arg);
    cout << result << endl;
    return 0;
}