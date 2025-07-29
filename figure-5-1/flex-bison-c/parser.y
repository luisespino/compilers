%{
#include <stdio.h>

int yylex(void);
void yyerror(const char *s);
%}

%token DIGIT
%token N

%%
l       : e N           { printf("%d\n",$1); }
        ;

e       : e '+' t       { $$ = $1 + $3; }
        | t             { $$ = $1; }
        ;

t       : t '*' f       { $$ = $1 * $3; }
        | f             { $$ = $1; }
        ;

f       : '(' e ')'     { $$ = $2; }
        | DIGIT         { $$ = $1; }
        ;

%%
extern char *yytext;
extern int yychar; 

void yyerror(const char *s) {
    if (yychar != YYEMPTY) {
        fprintf(stderr, "Error: %s on '%s'\n", s, yytext);
    }
}
