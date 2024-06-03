#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define HASH_TABLE_SIZE 100
#define MAX_KEY_LENGTH 50

typedef struct HashTableNode
{
    char key[MAX_KEY_LENGTH];
    int value;
    struct HashTableNoe *next;
} HashTableNode;
typedef struct HashTable
{
    HashTableNode *table[HASH_TABLE_SIZE];
} HashTable;

unsigned int hash_func(const char *key)
{
    unsigned int hash = 0;
    const char *ptr = key;
    while (*ptr)
    {
        hash += *ptr++;
    }
    return hash % HASH_TABLE_SIZE;
}
// insert
void insert(HashTable *hahstable, const char *key, int value)
{
    unsigned int hash = hash_func(key);
    HashTableNode *newNode = (HashTableNode *)malloc(sizeof(HashTableNode));
    strcpy(newNode->key, key);
    newNode->value = value;
    newNode->next = hahstable->table[hash];
    hahstable->table[hash] = newNode;
}
// search
int search(HashTable *hahstalbe, const char *key, int *value)
{
    unsigned int hash = hash_func(key);
    HashTableNode *currentHash = hahstalbe->table[hash];
    while (currentHash)
    {
        if (strcmp(currentHash->key, key) == 0)
        {
            *value = currentHash->value;
            return 1;
        }
        currentHash = currentHash->next;
    }
    return 0;
}

// 使用链地址法解决hash conflict
int main()
{
    HashTable hashtable;
    memset(&hashtable, 0, sizeof(HashTable));
    insert(&hashtable, "apple", 10);
    insert(&hashtable, "banana", 20);
    insert(&hashtable, "cherry", 30);

    int value;
    if (search(&hashtable, "banana", &value))
    {
        printf("The value of 'banana' is %d\n", value); // 输出：The value of 'banana' is 20
    }
    else
    {
        printf("'banana' not found\n");
    }

    return 0;
}