#include <iostream>
#include <future>
#include <thread>
#include <chrono>
using namespace std;
int do_something(int x)
{
    cout << "wait..." << endl;
    this_thread::sleep_for(chrono::seconds(2));
    return x * x;
}
int main()
{
    future<int> f = async(launch::async, do_something, 10);

    cout << "main do some thing" << endl;
    int v = f.get();
    cout << "get future value" << v << endl;
}
