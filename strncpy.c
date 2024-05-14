#include <stddef.h>
char *my_strncpy(char *dest, const char *src, size_t n)
{
    char *dest_start = dest;
    while (n-- && *src != '\0')
    {
        *dest++ = *src++;
    }
    while (n--)
    {
        *dest++ = '\0';
    }
    return dest_start;
}
char *safe_strncpy(char *dest, const char *src, size_t n)
{
    char *dest_start = dest;
    if (dest < src || (dest >= src + n))
    {
        // 正向复制（无重叠或dest在src之后）
        while (n-- && *src != '\0')
        {
            *dest++ = *src++;
        }
        // 填充剩余的n个字符为'\0'
        while (n--)
        {
            *dest++ = '\0';
        }
    }
    else
    {
        // 反向复制（dest在src之前且有重叠）
        src += n - 1;  // 移动到src的末尾
        dest += n - 1; // 移动到dest的末尾
        while (n--)
        {
            if (*src != '\0')
            {
                *dest-- = *src;
            }
            src--;
            if (dest < src)
            { // 当dest指针小于src时，说明已无重叠部分，可以提前结束
                break;
            }
        }
        // 从后往前填充'\0'直到达到原始dest位置
        dest++; // 因为上面循环结束时dest多减了一次
        while (dest < dest_start + n)
        {
            *dest++ = '\0';
        }
    }
    return dest_start;
}