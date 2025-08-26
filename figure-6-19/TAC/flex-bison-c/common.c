#include "common.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int temp = 1;

char* newTemp(void) {
    char* name = malloc(12); 
    if (!name) return NULL;

    snprintf(name, 12, "t%d", temp++);
    return name;  
}

char* strconcat(const char* a, const char* b) {
    size_t len = strlen(a) + strlen(b) + 1; 
    char* result = malloc(len);
    if (!result) return NULL;

    strcpy(result, a);
    strcat(result, b);

    return result;
}