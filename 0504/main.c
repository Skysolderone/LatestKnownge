#include <stdio.h>
#include <string.h>
#include <ctype.h>
#include "stack.h"

int parse(char str[])
{
    int len = strlen(str);
    Stack stack;
    int res = init_stack(&stack);
    if (res != 0)
    {
        printf("init error");
    }

    // ��ջ
    for (int i = 0; i < len / 2; i++)
    {
        int res = push_stack(&stack, str[i]);
        if (res != 0)
        {
            printf("push error");
        }
    }
    // �Ƚ�
    for (int i = len / 2; i < len; i++)
    {

        int res = pop_stack(&stack);
        if (res == -1)
        {
            printf("pop error");
            return -1;
        }
        if (str[i] != res)
        {
            return -1;
        }
    }

    return 0;
}
void main()
{

    for (;;)
    {

        char str[100];
        printf("��������Ҫ�жϵ�����,����q�˳�\n");
        // input
        int i = scanf("%s", str);
        int len = strlen(str);
        //'q' exit
        if (str[0] == 'q' && len == 1)
        {
            printf("quit....\n");
            break;
        }
        // verify
        int index = 1;
        int range = 0;
        while (str[range] != '\0')
        {
            if (!isdigit(str[range]))
            {
                index = 0;
                break;
            }
            range++;
        }
        if (i != 1 || index == 0)
        {
            printf("�����ʽ����\n");
            // ���������
            fflush(stdin);
            continue;
        }
        // parse
        int result = parse(str);
        if (result == 0)
        {
            printf("%s �����ǻ�����\n", str);
            continue;
        }

        printf("%s �������ǻ�����\n", str);
    };
}
