#include <iostream>

int divideByThree(int n)
{
    // process negative
    bool isNegative = false;
    if (n < 0)
    {
        isNegative = true;
        n = -n;
    }
    // int result = 0;
    // while (n >= (result + 1) * 3)
    // {
    //     // eg
    //     // int t = result << 2;
    //     // result = (t + n) / 3;
    //     // bit
    //     result = result << 2;
    //     result = result + (n >= (result + 1) * 3);
    //     n = n - (result + 1) * 3 + (n >= (result + 1) * 3) * 3;
    // }
    // return isNegative ? -result : result;
    int result = 0;
    while (n >= (result + 1) * 3)
    {
        result = result << 2; // result乘以4
        result = result + (n >= (result + 1) * 3);
        n = n - (result + 1) * 3 + (n >= (result + 1) * 3) * 3;
    }

    return isNegative ? -result : result;
}

int main()
{
    std::cout << divideByThree(10) << std::endl;  // 输出: 3
    std::cout << divideByThree(-10) << std::endl; // 输出: -3
    return 0;
}