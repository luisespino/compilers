#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#include <string.h>

char *x;

int F();
int T();
int E();
void S();

void S() {
    int result = E();
    printf("Result: %d\n", result);
}

int E() {
    int e = T();
    while (*x == '+') {
        x++; 
        int t = T();
        e = e + t;
    }
    return e;
}

int T() {
    int t = F();
    while (*x == '*') {
        x++;
        int f = F();
        t = t * f;
    }
    return t;
}

int F() {
    if (*x == '(') {
        x++;  
        int val = E();  
        if (*x == ')') {
            x++;  
        } else {
            printf("Error: Paren missing.\n");
            exit(1);
        }
        return val;
    } else if (isdigit(*x)) {
        int num = *x - '0';
        x++;
        return num;
    }

    printf("Error: character invalid '%c'\n", *x);
    exit(1);
}

int main() {
    char input[100];

    while (1) {
        printf("Enter expression (x to exit): ");
        fgets(input, sizeof(input), stdin);

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
