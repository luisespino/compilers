#ifndef NODE_H
#define NODE_H

#include <stdio.h>

typedef struct Node {
    char *value;
    struct Node *left;
    struct Node *right;
} Node;

Node* new_node(const char *value, Node *left, Node *right);

void postorder(Node *node);

void dot(void);

extern Node* root;

#endif
