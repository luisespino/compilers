#ifndef SYMBOL_H
#define SYMBOL_H

#include <search.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    char* name;
    int value;
} Symbol;

int symbol_table_create(size_t size);
void symbol_table_destroy(void);
int symbol_table_insert(const char* name, int value);
int symbol_table_search(const char* name, int* value);

#endif // SYMBOL_H