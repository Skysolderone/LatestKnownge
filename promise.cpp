#include <iostream>
#include <future>
// #include <promise>

using namespace std;
int do_something(int x)
{
    cout << "wait..." << endl;
    this_thread::sleep_for(chrono::seconds(2));
    return x * x;
}
void worker(promise<int> p)
{
    int x = do_something(5);
    p.set_value(x);
}
int main()
{
    promise<int> p;
    future<int> f = p.get_future();
    thread t(worker, std::move(p));
    int value = f.get();
    cout << "get" << value << endl;
    t.join();
}
