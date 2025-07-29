%{
#include "node.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int yylex(void);
void yyerror(const char *s);

%}

%union {
    int num;
    char* id;
    Node* node;
}

%token <num> NUM
%token <id> ID

%type <node> s e t f

%%

s   : e                { root = $1; }
    ;

e   : e '+' t          { $$ = new_node("+", $1, $3); }
e   : e '-' t          { $$ = new_node("-", $1, $3); }
    | t                { $$ = $1; }
    ;

t   : t '*' f          { $$ = new_node("*", $1, $3); }
t   : t '/' f          { $$ = new_node("/", $1, $3); }
    | f                { $$ = $1; }
    ;

f   : '(' e ')'        { $$ = $2; }
    | NUM              {
                          char buffer[32];
                          snprintf(buffer, sizeof(buffer), "%d", $1);
                          $$ = new_node(buffer, NULL, NULL);
                       }
    | ID               { $$ = new_node($1, NULL, NULL); }
    ;

%%

extern char *yytext;
extern int yychar;

void yyerror(const char *s) {
    if (yychar != YYEMPTY)
        fprintf(stderr, "Error: %s at '%s'\n", s, yytext);
    else
        fprintf(stderr, "Error: %s\n", s);
}
