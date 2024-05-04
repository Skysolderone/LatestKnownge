#ifndef STACK_H
#define STACK_H

#define MAX_SIEZ 100
typedef struct stack
{
    int top;
    int data[MAX_SIEZ];
} Stack;
// init
int init_stack(Stack *stack)
{
    (*stack).top = -1;
    return 0;
}
// push
int push_stack(Stack *stack, int i)
{

    if ((*stack).top == MAX_SIEZ - 1)
    {
        return -1;
    }
    (*stack).top++;

    (*stack).data[(*stack).top] = i;
    return 0;
}
// pop
int pop_stack(Stack *stack)
{

    if ((*stack).top == -1)
    {
        return -1;
    }
    int res = (*stack).data[(*stack).top];
    (*stack).top--;
    return res;
}
#endif