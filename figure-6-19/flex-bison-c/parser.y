%{
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "attr.h"

void yyerror(const char *s);
int yylex(void);


%}

%union {
    int num;
    char* str;
    Attr attr;
}

%token <str> ID
%token <num> NUM
%type <attr> p s e 

%left '+' '-'
%left '*' '/'

%%

p:
    s { printf("program: \n"); }
    ;

s:
    s s  { printf("setences: \n"); }
  ;

s:
    ID '=' e   { printf("Asignación: \n"); }
  ;

e:
    e '+' e     { printf("suma: \n"); }
  | e '-' e     { printf("resta: \n"); }
  | e '*' e     { printf("multi: \n"); }
  | e '/' e     { printf("divi: \n"); }
  | '(' e ')'   {   $$.addr = $2.addr;
                    $$.code = $2.code; }
  | ID          {   $$.addr = $1.addr;
                    $$.code = ""; }
  | NUM         {   $$.addr = $1.addr;
                    $$.code = ""; }
  ;

%%

void yyerror(const char *s) {
    fprintf(stderr, "Error: %s\n", s);
}
