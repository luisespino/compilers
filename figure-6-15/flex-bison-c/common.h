#ifndef COMMON_H
#define COMMON_H

extern int temp;
extern char* data;
extern char* code;

typedef enum { T_INT, T_FLOAT } Type;

typedef struct {    
    Type type;
    int width;
    char* data;
} Attr;

char* newTemp(void);

char* strconcat(const char* a, const char* b);

#endif