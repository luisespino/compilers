%{
#include "attr.h"
#include "parser.tab.h"
%}

%%

[ \t\r\n]+         ; // Ignorar espacios y saltos de línea

[a-zA-Z_][a-zA-Z0-9_]*  { yylval.str = strdup(yytext); return ID; }
[0-9]+                  { yylval.num = atoi(yytext); return NUM; }

"="     return '=';
"*"     return '*';
"/"     return '/';
"+"     return '+';
"-"     return '-';
"("     return '(';
")"     return ')';

.       return yytext[0];

%%

int yywrap() { return 1; }
