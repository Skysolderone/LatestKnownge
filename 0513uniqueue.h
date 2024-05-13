template <typename T>
class UniquePtr
{
    // 构造函数
    explicit UniquePtr(T *ptr = nullptr) : ptr_(ptr) {}
    // 禁止拷贝构造和拷贝赋值
    UniquePtr(const UniquePtr &) = delete;
    UniquePtr operator=(const UniquePtr &) = delete;
    // 允许移动构造和移动赋值
    UniquePtr(UniquePtr &&other) noexcept;
    UniquePtr &operator=(UniquePtr &&other) noexcept;
    // 重置指针
    void reset(T *ptr = nullptr);
    // get point
    T *get() const { return ptr_; }
    // release point
    void release();
    // 解引用操作
    T &operator*() const;
    T *operator->() const;
    // check point
    explicit operator bool() const { return ptr_ != nullptr; }

    ~UniquePtr();

private:
    T *ptr_;
};

template <typename T>
UniquePtr<T>::UniquePtr(UniquePtr &&other) noexcept : ptr_(other.ptr_)
{
    other.ptr_ = nullptr; // 将源对象的指针设为nullptr，以确保资源的独占性
}

template <typename T>
UniquePtr<T> &UniquePtr<T>::operator=(UniquePtr &&other) noexcept
{
    if (this != &other)
    {                         // 防止自赋值
        delete ptr_;          // 删除当前对象所指向的资源
        ptr_ = other.ptr_;    // 接管源对象的资源
        other.ptr_ = nullptr; // 将源对象的指针设为nullptr
    }
    return *this;
}

template <typename T>
void UniquePtr<T>::reset(T *ptr)
{
    delete ptr_; // 删除当前对象所指向的资源
    ptr_ = ptr;  // 接管新资源
}

template <typename T>
void UniquePtr<T>::release()
{
    T *temp = ptr_; // 保存原始指针
    ptr_ = nullptr; // 将内部指针设为nullptr
    return temp;    // 返回原始指针
}

template <typename T>
T &UniquePtr<T>::operator*() const
{
    return *ptr_; // 返回指向对象的引用
}

template <typename T>
T *UniquePtr<T>::operator->() const
{
    return ptr_; // 返回原始指针
}

template <typename T>
UniquePtr<T>::~UniquePtr()
{
    delete ptr_; // 删除所指向的对象
}