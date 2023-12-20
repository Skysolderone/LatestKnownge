#include <iostream>

int main()
{
    // 定义一个固定长度为5的数组
    const int arraySize = 5;
    int numbers[arraySize] = {1, 2, 3, 4, 5};

    // 显示原始数组
    std::cout << "Original Array: ";
    for (int i = 0; i < arraySize; ++i)
    {
        std::cout << numbers[i] << " ";
    }
    std::cout << std::endl;

    // 在索引为3的位置插入一个值（假设插入的值为99）
    int insertIndex1 = 3;
    int insertValue1 = 99;

    if (insertIndex1 >= 0 && insertIndex1 < arraySize)
    {
        numbers[insertIndex1] = insertValue1;
    }
    else
    {
        std::cout << "Invalid index for insertion." << std::endl;
    }

    // 显示第一次更新后的数组
    std::cout << "Array After Inserting Value at Index 3: ";
    for (int i = 0; i < arraySize; ++i)
    {
        std::cout << numbers[i] << " ";
    }
    std::cout << std::endl;

    // 在索引为3的位置再次插入一个值（假设插入的值为88）
    int insertIndex2 = 3;
    int insertValue2 = 88;

    if (insertIndex2 >= 0 && insertIndex2 < arraySize)
    {
        numbers[insertIndex2] = insertValue2;
    }
    else
    {
        std::cout << "Invalid index for insertion." << std::endl;
    }

    // 显示第二次更新后的数组
    std::cout << "Array After Inserting Another Value at Index 3: ";
    for (int i = 0; i < arraySize; ++i)
    {
        std::cout << numbers[i] << " ";
    }
    std::cout << std::endl;

    return 0;
}