%{
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include "symbol.h"
#include "common.h"

void yyerror(const char *s);
int yylex(void);

extern char *yytext; 

%}

%union {
    int num;
    char* str;
    Attr attr;
}

%token <str> ID
%token <num> NUM
%token INT
//%type <attr> prog stmts stmt expr decl assign index array_init elements


%%

prog:
    stmts 
    ;

stmts:
    stmts stmt
    | stmt
    ;

stmt:
    decl 
    | assign
    ;

decl:
    INT ID '=' expr ';'
    | INT ID '=' array_init ';'
    ;

assign:
    ID index '=' expr ';'
    | ID '=' expr ';'
    ;

expr:
    NUM
    | ID index
    | ID
    ;

index:
    index '[' expr ']'
    | '[' expr ']'
    ;

array_init:
    '[' ']'
    | '[' elements ']'
    | '[' array_list ']'    
    ;

elements:
    elements ',' expr
    | expr
    ;

array_list:
    array_list ',' array_init
    | array_init
    ;


%%

void yyerror(const char *s) {
    fprintf(stderr, "Error: %s, near of '%s'\n", s, yytext);
}