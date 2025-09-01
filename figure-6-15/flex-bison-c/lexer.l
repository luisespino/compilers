%{
#include "symbol.h"
#include "common.h"
#include "parser.tab.h"
%}

%%

int                     { return INT; }
float                   { return FLOAT; }

[a-zA-Z_][a-zA-Z0-9_]*  { yylval.str = strdup(yytext); return ID; }
[0-9]+                  { yylval.num = atoi(yytext); return NUM; }
[0-9]+\.[0-9]+          { yylval.fnum = atof(yytext); return FNUM; }

"["                     return '[';
"]"                     return ']';

[ \t\r\n]+              ; // Ignorar espacios y saltos de línea

.                       return yytext[0];

%%

int yywrap()            { return 1; }
