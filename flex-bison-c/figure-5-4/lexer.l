%{
#include "parser.tab.h"
#include <stdlib.h>
%}

%%

[0-9]+      { yylval = atoi(yytext); return DIGIT; }
[+*()]      { return yytext[0]; }
[n]         { return N; }
[ \t\n]     ; // ignore
.           { return yytext[0]; }

%%

