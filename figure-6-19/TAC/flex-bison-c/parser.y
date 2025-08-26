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
    char* str;
    Attr attr;
}

%token <str> ID
%token <num> NUM
%type <attr> p s e 

%left '+' '-'
%left '*' '/'
%right UMINUS

%%

p:
    s {
        printf("Three Address Code: \n");
        printf($1.code);
    }
    ;

s:
    s s {
        $$.code = strconcat($1.code, $2.code);
    }
    ;

s:
    ID '=' e {
        symbol_table_insert($1, 0);
        $$.addr = 0;
        $$.type = ADDR_ID;
        $$.code = strconcat($3.code, $1);
        $$.code = strconcat($$.code, " = ");
        switch ($3.type) {
            case ADDR_INT:
                char buffer[12]; 
                int value = *(int*)$3.addr;
                sprintf(buffer, "%d", value);
                $$.code = strconcat($$.code, buffer);
                break;
            case ADDR_ID:
            case ADDR_TMP:
                $$.code = strconcat($$.code, $3.addr);
                break;
        }
        $$.code = strconcat($$.code, "\n");
    }
    ;

e:
    e '+' e {
        
        $$.addr = newTemp();
        $$.type = ADDR_TMP;
        $$.code = strconcat($1.code, $3.code);
        $$.code = strconcat($$.code, $$.addr);
        $$.code = strconcat($$.code, " = ");

        switch ($1.type) {
            case ADDR_INT:
                char buffer[12];
                int value = *(int*)$1.addr;
                sprintf(buffer, "%d", value);
                $$.code = strconcat($$.code, buffer);
                $$.code = strconcat($$.code, " + ");
                switch ($3.type) {
                    case ADDR_INT:
                        int value = *(int*)$3.addr;
                        sprintf(buffer, "%d", value);
                        $$.code = strconcat($$.code, buffer);
                        break;
                    case ADDR_ID:
                    case ADDR_TMP:
                        $$.code = strconcat($$.code, $3.addr);
                        break;
                }
                break;
            case ADDR_ID:
            case ADDR_TMP:
                $$.code = strconcat($$.code, $1.addr);
                $$.code = strconcat($$.code, " + ");
                switch ($3.type) {
                    case ADDR_INT:
                        int value = *(int*)$3.addr;
                        sprintf(buffer, "%d", value);
                        $$.code = strconcat($$.code, buffer);
                        break;
                    case ADDR_ID:
                    case ADDR_TMP:
                        $$.code = strconcat($$.code, $3.addr);
                        break;
                }
                break;
        }
        $$.code = strconcat($$.code, "\n");      
    }

  | e '*' e {
        $$.addr = newTemp();
        $$.type = ADDR_TMP;
        $$.code = strconcat($1.code, $3.code);
        $$.code = strconcat($$.code, $$.addr);
        $$.code = strconcat($$.code, " = ");
        switch ($1.type) {
            case ADDR_INT:
                char buffer[12];
                int value = *(int*)$1.addr;
                sprintf(buffer, "%d", value);
                $$.code = strconcat($$.code, buffer);
                $$.code = strconcat($$.code, " * ");
                switch ($3.type) {
                    case ADDR_INT:
                        int value = *(int*)$3.addr;
                        sprintf(buffer, "%d", value);
                        $$.code = strconcat($$.code, buffer);
                        break;
                    case ADDR_ID:
                    case ADDR_TMP:
                        $$.code = strconcat($$.code, $3.addr);
                        break;
                }
                break;
            case ADDR_ID:
            case ADDR_TMP:
                $$.code = strconcat($$.code, $1.addr);
                $$.code = strconcat($$.code, " * ");
                switch ($3.type) {
                    case ADDR_INT:
                        int value = *(int*)$3.addr;
                        sprintf(buffer, "%d", value);
                        $$.code = strconcat($$.code, buffer);
                        break;
                    case ADDR_ID:
                    case ADDR_TMP:
                        $$.code = strconcat($$.code, $3.addr);
                        break;
                }
                break;
        }
        $$.code = strconcat($$.code, "\n");      
    }

  | '-' e  %prec UMINUS {
        $$.addr = newTemp();
        $$.type = ADDR_TMP;
        $$.code = strconcat($2.code, $$.addr);

        switch ($2.type) {
            case ADDR_INT:
                $$.code = strconcat($$.code, " = -");
                char buffer[12];
                int value = *(int*)$2.addr;
                sprintf(buffer, "%d", value);
                $$.code = strconcat($$.code, buffer);
                break;
            case ADDR_ID:
            case ADDR_TMP:
                $$.code = strconcat($$.code, " = -");
                $$.code = strconcat($$.code, $2.addr);
                break;
        }
        $$.code = strconcat($$.code, "\n");      
    }    

  | '(' e ')' {
        $$.addr = $2.addr;
        $$.type = $2.type;
        $$.code = $2.code;
    }

  | ID {
        int value;
        if (symbol_table_search($1, &value)) {
            $$.addr = $1;
            $$.type = ADDR_ID;
        } else {
            char msg[128];
            snprintf(msg, sizeof(msg), "Undeclared variable: %s", $1);
            yyerror(msg);
            YYABORT;
        }
        $$.code = "";
    }

  | NUM {
        int* val = malloc(sizeof(int));
        *val = $1;
        $$.addr = val;
        $$.type = ADDR_INT;
        $$.code = "";
    }
    ;

%%

void yyerror(const char *s) {
    fprintf(stderr, "Error: %s\n", s);
}
