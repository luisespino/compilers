/* Figure 5.4 */

%{
#include "fig54.tab.hpp"
#include <iostream>
using namespace std;
%}

%option noyywrap

%%

[0-9]+  	{ yylval.val = atoi(yytext); return DIGIT; }
"+"		{ return PLUS; }
"*"		{ return TIMES; }
"("      	{ return LPAREN; }
")"      	{ return RPAREN; }
"\n"      	{ return N; }
[ \t]+ 	{ }

%%
