#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h> 
#include <string.h> 
#include "symbol.h"
#include "common.h"

extern int yyparse(void);
extern void yy_scan_string(const char *);

int main(void) {
    char input[1024];

    if (!symbol_table_create(10)) {
        return 1;
    }    

    while (true) {
        temp = 1;
        printf("Enter expression:\n");

        if (!fgets(input, sizeof(input), stdin))
            break;

        if (strcmp(input, "\n") == 0) {
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
