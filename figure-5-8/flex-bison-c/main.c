#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h> 
#include <string.h> 

extern int yyparse(void);
extern void yy_scan_string(const char *);

typedef struct Symbol {
    char *id;
    char *type;
    struct Symbol *next;
} Symbol;

Symbol *symbol_table = NULL;

void addType(const char *id, const char *type) {
    Symbol *s = malloc(sizeof(Symbol));
    s->id = strdup(id);
    s->type = strdup(type);
    s->next = symbol_table;
    symbol_table = s;
}

void printSymbolTable() {
    Symbol *current = symbol_table;
    while (current != NULL) {
        printf("ID: %s, Tipo: %s\n", current->id, current->type);
        current = current->next;
    }
}

int main(void) {
    char input[1024];

    while (true) {
        printf("Enter expression:\n");

        if (!fgets(input, sizeof(input), stdin))
            break;

        if (strcmp(input, "\n") == 0) {
            continue;
        }

        if (strcmp(input, "symbol\n") == 0) {
            printSymbolTable();
            continue;
        }        

        if (strcmp(input, "exit\n") == 0) {
            printf("Exiting...\n");
            break;
        }

        yy_scan_string(input);
        yyparse();
        
    }

    return 0;
}
