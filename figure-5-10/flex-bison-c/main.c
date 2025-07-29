#include "node.h"
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h> 
#include <string.h> 

extern int yyparse(void);
extern void yy_scan_string(const char *);

int main(void) {
    char input[1024];

    while (true) {
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

        if (strcmp(input, "dot\n") == 0) {
            dot();
            continue;
        }


        yy_scan_string(input);
        yyparse();

        printf("Postorder traversal: ");
        postorder(root);
        printf("\n");        
        
    }

    return 0;
}
