#include <iostream>
#include <vector>
using namespace std;
// 阶乘函数的递归实现
long long factorial(int n) {
    // 基本情况
    if (n == 0 || n == 1) {
        return 1;
    }
    // 递归情况
    else {
        return n * factorial(n - 1); // 递归调用
    }
}
// 斐波那契数列的递归实现
int fibonacci(int n) {
    // 基本情况
    if (n <= 1) {
        return n;
    }
    // 递归情况
    else {
        return fibonacci(n - 1) + fibonacci(n - 2); // 递归调用
    }
}
// 二分搜索的递归实现
bool binarySearchRecursive(const std::vector<int>& arr, int target, int left, int right) {
    if (left > right) {
        return false; // 基本情况：未找到目标值
    }
    int mid = left + (right - left) / 2; // 防止溢出
    if (arr[mid] == target) {
        return true; // 基本情况：找到目标值
    } else if (arr[mid] < target) {
        return binarySearchRecursive(arr, target, mid + 1, right); // 递归情况：在右半部分搜索
    } else {
        return binarySearchRecursive(arr, target, left, mid - 1); // 递归情况：在左半部分搜索
    }
}
// int main() {
//     std::vector<int> arr = {1, 3, 5, 7, 9, 11, 13, 15, 17, 19};
//     int target;
//     std::cout << "请输入要搜索的目标值：";
//     std::cin >> target;
//     if (binarySearchRecursive(arr, target, 0, arr.size() - 1)) {
//         std::cout << "找到目标值 " << target << std::endl;
//     } else {
//         std::cout << "未找到目标值 " << target << std::endl;
//     }
//     return 0;
// }
int main(){
    int number;
    cout<<"please enter num" << endl;
    cin>>number;
    cout << number << "的阶乘是：" << factorial(number) << endl;    
    return 0;
}