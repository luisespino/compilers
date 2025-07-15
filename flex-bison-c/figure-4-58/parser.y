%{
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int yylex(void);
void yyerror(const char *s);
%}

%token DIGIT

%%
line    : expr '\n'         { printf("Result: %d\n", $1); }
        ;

expr    : expr '+' term     { $$ = $1 + $3; }
        | term
        ;

term    : term '*' factor   { $$ = $1 * $3; }
        | factor
        ;

factor  : '(' expr ')'      { $$ = $2; }
        | DIGIT
        ;

%%
void yyerror(const char *s) {
    fprintf(stderr, "Error: %s\n", s);
}
