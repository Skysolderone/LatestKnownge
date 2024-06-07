#include <iostream>

struct ListNode
{
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(nullptr) {}
};

// 递归反转链表
ListNode *reverseListRecursive(ListNode *head)
{
    //
    // if (head == nullptr || head->next == nullptr)
    // {
    //     return head;
    // }
    // ListNode *newHead = reverseListRecursive(head->next);
    // head->next = head;
    // head->next = nullptr;
    // return newHead;
    ListNode *prev = nullptr;
    ListNode *curr = head;

    while (curr != nullptr)
    {
        ListNode *nextTemp = curr->next; // 暂存下一个节点
        curr->next = prev;               // 反转指针方向
        prev = curr;                     // 前一个节点后移
        curr = nextTemp;                 // 当前节点后移
    }

    return prev; // 当curr为空时，prev就是新的头节点
}

void printList(ListNode *head)
{
    while (head != nullptr)
    {
        std::cout << head->val << " ";
        head = head->next;
    }
    std::cout << std::endl;
}

int main()
{
    ListNode *head = new ListNode(1);
    head->next = new ListNode(2);
    head->next->next = new ListNode(3);
    head->next->next->next = new ListNode(4);
    std::cout << "init link:";
    printList(head);
    ListNode *newhead = reverseListRecursive(head);
    std::cout << "reverse link";
    printList(newhead);
    return 0;
}