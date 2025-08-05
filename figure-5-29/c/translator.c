#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#include <string.h>

char *x;
int temp_count = 0;
char code[1000];
int code_pos = 0;

char *new_temp() {
    static char temp[10];
    sprintf(temp, "t%d", ++temp_count);
    return strdup(temp);
}


void append_code(const char *line) {
    code_pos += sprintf(code + code_pos, "%s\n", line);
}


char *F();
char *T();
char *E();
void S();

void S() {
    temp_count = 0;
    code[0] = '\0';
    code_pos = 0;

    char *result = E();
    char line[100];
    sprintf(line, "// Final result on %s", result);
    append_code(line);

    printf("Code:\n%s", code);
    free(result);
}

char *E() {
    char *left = T();
    while (*x == '+') {
        x++;
        char *right = T();
        char *temp = new_temp();
        char line[100];
        sprintf(line, "%s = %s + %s", temp, left, right);
        append_code(line);
        free(left);
        free(right);
        left = temp;
    }
    return left;
}

char *T() {
    char *left = F();
    while (*x == '*') {
        x++;
        char *right = F();
        char *temp = new_temp();
        char line[100];
        sprintf(line, "%s = %s * %s", temp, left, right);
        append_code(line);
        free(left);
        free(right);
        left = temp;
    }
    return left;
}

char *F() {
    if (*x == '(') {
        x++;
        char *val = E();
        if (*x == ')') {
            x++;
        } else {
            printf("Error: missing paren.\n");
            exit(1);
        }
        return val;
    } else if (isdigit(*x)) {
        char *val = malloc(2);
        val[0] = *x;
        val[1] = '\0';
        x++;
        return val;
    } else {
        printf("Error: invalid character '%c'\n", *x);
        exit(1);
    }
}

int main() {
    char input[100];

    while (1) {
        printf("Enter expression (x to exit): ");
        fgets(input, sizeof(input), stdin);

        // Quitar '\n' si lo hay
        if (input[strlen(input) - 1] == '\n') {
            input[strlen(input) - 1] = '\0';
        }

        if (strcmp(input, "x") == 0 || strcmp(input, "X") == 0) {
            printf("Exiting...\n");
            break;
        }

        x = input;
        S();
    }

    return 0;
}
