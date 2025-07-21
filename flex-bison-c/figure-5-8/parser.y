%{
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int yylex();
void yyerror(const char *s);

extern void addType(const char *id, const char *type);
%}

%union {
    char *type;   // Tipo (int/float)
    char *id;     // Nombre del identificador
}

/* Tokens */
%token <id> ID
%token INT FLOAT COMMA

/* Tipos de los no terminales */
%type <type> T L D


%%
D   : T L           { 
                        printf("Declaración completada. Tipo heredado: %s\n", $1); 
                        free($1); 
                    }
    ;

T   : INT           {   $$ = strdup("integer"); }
    | FLOAT         {   $$ = strdup("float"); }
    ;

L   : L COMMA ID    {   
                        addType($3, $1); 
                        printf("addType('%s', '%s')\n", $3, $1);
                        $$ = $1;
                    }
    | ID            { 
                        addType($1, $<type>0);
                        printf("addType('%s', '%s')\n", $1, $<type>0);
                        $$ = $<type>0;
                    }
    ;
%%

extern char *yytext;
extern int yychar; 

void yyerror(const char *s) {
    fprintf(stderr, "Error: %s", s);
    if (yychar != YYEMPTY) {
            fprintf(stderr, " on '%s'", yytext);
    }
    fprintf(stderr, "\n");  
}

