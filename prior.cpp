#include <iostream>
#include <queue>

using namespace std;
int main()
{
    // 创建一个空的优先级队列
    priority_queue<int> pq;

    // 向优先级队列中添加元素
    pq.push(3);
    pq.push(5);
    pq.push(1);
    pq.push(4);

    // 输出优先级队列的大小和队首元素
    cout << "Size of priority queue: " << pq.size() << endl;
    cout << "Top element: " << pq.top() << endl; // 输出5，因为5是优先级最高的元素

    // 删除队首元素并输出剩余元素的大小和队首元素
    pq.pop();
    cout << "Size after pop: " << pq.size() << endl;
    cout << "Top element after pop: " << pq.top() << endl; // 输出4，现在是优先级最高的元素

    return 0;
}