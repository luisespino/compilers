%{
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "symbol.h"
#include "common.h"

void yyerror(const char *s);
int yylex(void);

%}

%union {
    int num;
    float fnum;
    char* str;
    Attr attr;
}

%token <str> ID
%token <num> NUM
%token <fnum> FNUM
%token INT FLOAT

%type <attr> p d t b c 

%%

p:
    d {
        printf("ARM64: \n");
        char* program = ".global _start\n\n.data\n";
        program = strconcat(program, $1.data);
        program = strconcat(program, "\n.text\n_start:\n\tmov x8, #93\n\tsvc #0\n");
        printf(program);
    }
    ;

d:
    t ID d {
        if ($1.type == T_INT) {
            $$.data = strconcat($2, ": .word 0");
            if ($1.width > 4) {
                int size = $1.width / 4;
                for (int i = 1; i < size; i++)
                    $$.data = strconcat($$.data, ", 0");
            }
        }
        else if ($1.type == T_FLOAT) {
            $$.data = strconcat($2, ": .double 0.0");
            if ($1.width > 8) {
                int size = $1.width / 8;
                for (int i = 1; i < size; i++)
                    $$.data = strconcat($$.data, ", 0.0");
            }
        }
        $$.data = strconcat($$.data, "\n");
        $$.data = strconcat($$.data, $3.data);
        symbol_table_insert($2,0);
    }
    | /* empty */ { $$.data = ""; }
    ;

t:
    b c {
        $$.type = $1.type;
        $$.width = ($2.width != 1) ? $2.width * $1.width : $1.width;
    }
    ;

b:
    INT {
        $$.type = T_INT;
        $$.width = 4;
    }
    | FLOAT {
        $$.type = T_FLOAT;
        $$.width = 8;
    }
    ;

c:
    '[' NUM ']' c {
        $$.width = $2 * $4.width;
    }
    | /* empty */ { $$.width = 1; }
    ;



%%

void yyerror(const char *s) {
    fprintf(stderr, "Error: %s\n", s);
}
