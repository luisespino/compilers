#include "node.h"
#include <stdlib.h>
#include <string.h>

Node* root = NULL;

Node* new_node(const char *value, Node *left, Node *right) {
    Node *node = malloc(sizeof(Node));
    node->value = strdup(value);
    node->left = left;
    node->right = right;
    return node;
}

void postorder(Node *node) {
    if (node == NULL) return;

    postorder(node->left);
    postorder(node->right);
    printf("%s ", node->value);
}

static int create_dot(Node* node, FILE* out) {
    static int id = 0;
    if (!node) return -1;

    int my_id = id++;
    fprintf(out, "    %d [label=\"%s\"];\n", my_id, node->value);

    int left_id = create_dot(node->left, out);
    int right_id = create_dot(node->right, out);

    if (left_id != -1)
        fprintf(out, "    %d -> %d;\n", my_id, left_id);
    if (right_id != -1)
        fprintf(out, "    %d -> %d;\n", my_id, right_id);

    return my_id;
}

void dot() {
    FILE *out = fopen("tree.dot", "w");
    if (!out) {
        perror("Error creating tree.dot");
        return;
    }

    fprintf(out, "digraph AST {\n");
    fprintf(out, "    node [shape=ellipse];\n");

    create_dot(root, out);

    fprintf(out, "}\n");
    fclose(out);

    printf("File tree.dot created.\n");
}
