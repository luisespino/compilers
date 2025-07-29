%{
#include "parser.tab.h"
#include <stdlib.h>
%}

%%
[0-9]+      { yylval = atoi(yytext); return DIGIT; }
[\n+*()]    { return yytext[0]; }
[ \t]       ; // ignore
.           { return yytext[0]; }
%%
