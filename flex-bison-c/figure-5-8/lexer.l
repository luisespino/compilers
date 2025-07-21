%{
#include "parser.tab.h"
#include <stdlib.h>
%}

%%
"int"           { return INT; }
"float"         { return FLOAT; }
","             { return COMMA; }
[a-zA-Z]+       { yylval.id = strdup(yytext); return ID; }
[ \t\n]         ;  // ignore
.               { return yytext[0]; }
%%

int yywrap()    { return 1; }

