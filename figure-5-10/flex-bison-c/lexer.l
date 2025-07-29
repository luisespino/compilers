%{
#include "node.h"
#include "parser.tab.h"
#include <stdlib.h>
%}

%%

[0-9]+                  { yylval.num = atoi(yytext); return NUM; }
[a-zA-Z][a-zA-Z0-9_]*   { yylval.id = strdup(yytext); return ID; }
[+*()]                  { return yytext[0]; }
[ \t\n]                 ; // ignore
.                       { return yytext[0]; }

%%

