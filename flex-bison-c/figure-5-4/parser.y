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

e       : t ep          { $$ = $1 + $2; }
        ;

ep      : '+' t ep      { $$ = $2 + $3; }
        |               { $$ = 0; }
        ;

t       : f tp          { $$ = $1 * $2; } 
        ;

tp      : '*' f tp      { $$ = $2 * $3; }     
        |               { $$ = 1; }
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
