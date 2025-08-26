#ifndef COMMON_H
#define COMMON_H

extern int temp;

typedef enum { ADDR_INT, ADDR_ID, ADDR_TMP } AddrType;

typedef struct {
    void* addr;
    AddrType type;
    char* code;
} Attr;

char* newTemp(void);

char* strconcat(const char* a, const char* b);

#endif