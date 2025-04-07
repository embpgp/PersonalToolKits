#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define IP_BYTE_COUNT 4
#define BYTE_MAX_VALUE 256

// Trie 树节点结构
typedef struct TrieNode {
    struct TrieNode *children[BYTE_MAX_VALUE];
    int is_end_of_rule; // 标记是否为规则的结束
} TrieNode;

// 创建新的 Trie 节点
TrieNode* createNode() {
    TrieNode* node = (TrieNode*)malloc(sizeof(TrieNode));
    if (node) {
        for (int i = 0; i < BYTE_MAX_VALUE; i++) {
            node->children[i] = NULL;
        }
        node->is_end_of_rule = 0;
    }
    return node;
}

// 插入 IP 地址到 Trie 树
void insert(TrieNode* root, const unsigned char ip[IP_BYTE_COUNT]) {
    TrieNode* current = root;
    for (int i = 0; i < IP_BYTE_COUNT; i++) {
        int index = ip[i];
        if (!current->children[index]) {
            current->children[index] = createNode();
        }
        current = current->children[index];
    }
    current->is_end_of_rule = 1;
}

// 匹配 IP 地址
int match(TrieNode* root, const unsigned char ip[IP_BYTE_COUNT]) {
    TrieNode* current = root;
    for (int i = 0; i < IP_BYTE_COUNT; i++) {
        int index = ip[i];
        if (!current->children[index]) {
            return 0; // 未匹配到
        }
        current = current->children[index];
    }
    return current->is_end_of_rule; // 返回是否为规则的结束
}

// 释放 Trie 树内存
void freeTrie(TrieNode* node) {
    if (!node) return;
    for (int i = 0; i < BYTE_MAX_VALUE; i++) {
        if (node->children[i]) {
            freeTrie(node->children[i]);
        }
    }
    free(node);
}

int main() {
    TrieNode* root = createNode();

    // 模拟 100 条随机 IP 规则
    unsigned char rules[100][IP_BYTE_COUNT];
    for (int i = 0; i < 100; i++) {
        for (int j = 0; j < IP_BYTE_COUNT; j++) {
            rules[i][j] = rand() % BYTE_MAX_VALUE;
        }
        if (i == 88) {
            rules[i][0] = 1;
            rules[i][1] = 1;
            rules[i][2] = 1;
            rules[i][3] = 1;
        }
        printf("%d: %d.%d.%d.%d\n",i, rules[i][0],rules[i][1],rules[i][2],rules[i][3]);
        insert(root, rules[i]);
    }

    // 测试匹配一个随机 IP
    unsigned char test_ip[IP_BYTE_COUNT];
    for (int i = 0; i < IP_BYTE_COUNT; i++) {
        test_ip[i] = 1;//rand() % BYTE_MAX_VALUE;
    }
    printf("test_ip: %d.%d.%d.%d\n", test_ip[0],test_ip[1],test_ip[2],test_ip[3]);
    if (match(root, test_ip)) {
        printf("IP 匹配成功！\n");
    } else {
        printf("IP 未匹配到规则。\n");
    }

    // 释放 Trie 树内存
    freeTrie(root);

    return 0;
}