#include <iostream>
#include <future>
#include <thread>

using namespace std;

int main()
{
    packaged_task<int()> task([]
                              { return 7; });
    future<int> f1 = task.get_future();
    thread t(std::move(task));
    future<int> f2 = async(launch::async, []
                           { return 8; });
    promise<int> p;
    future<int> f3 = task.get_future();
    thread([&p]
           { p.set_value_at_thread_exit(20); })
        .detach();
    cout << "waiting ....." << endl;
    f1.wait();
    f2.wait();
    f3.wait();
    cout << "result are:" << f1.get() << "" << f2.get() << "" << f3.get() << endl;
    t.join();
}