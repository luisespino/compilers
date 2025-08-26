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
        printf("ARM64: \n");
        char* program = strconcat(data, code);
        printf(program);
    }
    ;

s:
    s s {
        
    }
    ;

s:
    ID '=' e {
        
        $$.addr = 0;
        $$.code = NULL;
        int value;

        switch ($3.type) {
            case ADDR_INT:
                
                if (symbol_table_search($1, &value)) {
                    code = strconcat(code,"\tldr X0, =");
                    code = strconcat(code,$1);
                    code = strconcat(code,"\n");

                    code = strconcat(code,"\tmov X1, #");
                    char buffer[12];
                    value = *(int*)$3.addr;
                    sprintf(buffer, "%d", value);
                    code = strconcat(code,buffer);
                    code = strconcat(code,"\n");
                    
                    code = strconcat(code,"\tstr X1, [X0]\n\n");
                } else {
                    data = strconcat(data, $1);
                    data = strconcat(data, ": .word ");
                    char buffer[12];
                    value = *(int*)$3.addr;
                    sprintf(buffer, "%d", value);
                    data = strconcat(data,buffer);
                    data = strconcat(data,"\n");                    
                }        
                break;
            case ADDR_ID:
            case ADDR_TMP:
                if (!symbol_table_search($1, &value)) {
                    data = strconcat(data, $1);
                    data = strconcat(data, ": .word ");
                    data = strconcat(data, "0");
                    data = strconcat(data,"\n");                                        
                } 
                code = strconcat(code,"\tldr X0, =");
                code = strconcat(code,$1);
                code = strconcat(code,"\n");

                code = strconcat(code,"\tldr X1, =");
                code = strconcat(code,$3.addr);
                code = strconcat(code,"\n");

                code = strconcat(code,"\tldr X2, [X1]\n");

                code = strconcat(code,"\tstr X2, [X0]\n\n");                       
                break;
        }
        symbol_table_insert($1,0);
    }
    ;

e:
    e '+' e {
      
        $$.addr = newTemp();
        $$.type = ADDR_TMP;
        $$.code = NULL;

        char buffer[12];
        int value;
 
        data = strconcat(data, $$.addr);
        data = strconcat(data, ": .word ");
        data = strconcat(data, "0");
        data = strconcat(data,"\n");                                        

        code = strconcat(code,"\tldr X0, =");
        code = strconcat(code,$$.addr);
        code = strconcat(code,"\n");



        switch ($1.type) {
            case ADDR_INT:
                code = strconcat(code,"\tldr X1, [X0]\n");
                switch ($3.type) {
                    case ADDR_INT:

                        code = strconcat(code,"\tadd X1, X1, #");
                        
                        value = *(int*)$1.addr;
                        sprintf(buffer, "%d", value);
                        code = strconcat(code, buffer);
                        code = strconcat(code,"\n");

                        code = strconcat(code,"\tadd X1, X1, #");
                        value = *(int*)$3.addr;
                        sprintf(buffer, "%d", value);
                        code = strconcat(code, buffer);
                        code = strconcat(code,"\n");

                        code = strconcat(code,"\tstr X1, [X0]\n");
                        break;
                    case ADDR_ID:
                    case ADDR_TMP:
                        code = strconcat(code,"\tadd X1, X1, #");
                        value = *(int*)$1.addr;
                        sprintf(buffer, "%d", value);
                        code = strconcat(code, buffer);
                        code = strconcat(code,"\n");

                        code = strconcat(code,"\tadd X1, X1, =");
                        code = strconcat(code, $3.addr);
                        code = strconcat(code,"\n");

                        code = strconcat(code,"\tldr X3, [X2]\n");

                        code = strconcat(code,"\tadd X1, X1. X3\n");

                        code = strconcat(code,"\tstr X1, [X0]\n");
                        break;
                }
                break;
            case ADDR_ID:
            case ADDR_TMP:
                switch ($3.type) {
                    case ADDR_INT:
                        code = strconcat(code,"\tldr X1, =");
                        code = strconcat(code, $1.addr);
                        code = strconcat(code,"\n");

                        code = strconcat(code,"\tldr X2, [X1]\n");

                        code = strconcat(code,"\tadd X2, X2, #");
                        int value = *(int*)$3.addr;
                        sprintf(buffer, "%d", value);
                        code = strconcat(code, buffer);
                        code = strconcat(code,"\n");

                        code = strconcat(code,"\tstr X2, [X0]\n");
                        break;
                    case ADDR_ID:
                    case ADDR_TMP:
                        code = strconcat(code,"\tldr X1, =");
                        code = strconcat(code, $1.addr);
                        code = strconcat(code,"\n");

                        code = strconcat(code,"\tldr X2, [X1]\n");

                        code = strconcat(code,"\tldr X3, =");
                        code = strconcat(code, $3.addr);
                        code = strconcat(code,"\n");

                        code = strconcat(code,"\tldr X4, [X3]\n");

                        code = strconcat(code,"\tadd X2, X2, X4");

                        code = strconcat(code,"\tstr X2, [X0]\n");
                        break;
                }
                break;
        }
        code = strconcat(code, "\n");      
    }

  | e '*' e {
        $$.addr = newTemp();
        $$.type = ADDR_TMP;
        $$.code = NULL;

        /* missing translate code */
    }

  | '-' e  %prec UMINUS {
        $$.addr = newTemp();
        $$.type = ADDR_TMP;
        $$.code = NULL;

        /* missing translate code */
        
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
