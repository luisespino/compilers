%{
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

int yylex();
void yyerror(const char *s);

int size = 5;

int yylval;
%}

%token DIGIT

%left '+'
%left '*'

%%
input:
    /* vacío */
  | input expr '\n' { printf("%d\n", $2); }
  | input error '\n' { yyerrok; }   
  ;

expr:
      expr '+' expr { $$ = ($1 + $3 + 1) % size; }
    | expr '*' expr { $$ = ($1 * $3 + $1 + $3) % size; }
    | '(' expr ')'  { $$ = $2; }
    | DIGIT         { $$ = $1; }
    ;
%%

int yylex() {
    int c;
    while ((c = getchar()) == ' ' || c == '\t');  // saltar espacios

    if (c == EOF) return 0;

    if (isdigit(c)) {
        int digit = c - '0';
        if (digit < 0 || digit >= size) {
            fprintf(stderr, "Error: out of bounds [0, %d): %d\n", size, digit);
            while ((c = getchar()) != '\n' && c != EOF);
            return '\n';
        }
        yylval = digit;
        return DIGIT;
    }

    if (c == '+' || c == '*' || c == '(' || c == ')') return c;

    if (c == '\n') return '\n';

    return c;
}

void yyerror(const char *s) {
    fprintf(stderr, "Error: %s\n", s);
}

int main(int argc, char *argv[]) {
    int temp_size;
    char *endptr;

    if (argc > 1) {
        temp_size = strtol(argv[1], &endptr, 10);

        if (*endptr != '\0' || temp_size < 2 || temp_size > 10) {
            fprintf(stderr, "Error: the argument must be an integer between 2 and 10.\n");
            return 1;
        }

        size = temp_size;  // Solo aquí lo asignas si es válido
    }

    printf("Size: %dx%d\n", size, size);
    printf("Enter a expression, Ctrl+D to exit:\n");

    yyparse();
    return 0;
}
