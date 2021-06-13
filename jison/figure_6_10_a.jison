/* description: generate the DOT string to create the AST of a expression. */

/* lexical grammar */
%lex

%{
if (!('id' in yy)) {
  yy.id = 0;
}
%}

%%

\s+                   /* skip whitespace */
[0-9]+("."[0-9]+)?\b  return 'NUMBER'
[a-zA-Z][a-zA-Z0-9]*  return 'ID'
"*"                   return '*'
"/"                   return '/'
"-"                   return '-'
"+"                   return '+'
"="                   return '='
";"                   return ';'
"("                   return '('
")"                   return ')'
<<EOF>>               return 'EOF'
.                     return 'INVALID'

/lex

/* operator associations and precedence */

%left '+' '-'
%left '*' '/'
%left UMINUS

%start expressions

%% /* language grammar */

expressions
    : ID '=' e ';' EOF
        {$$ = $3[1]+$1[0]+' = '+$3[0]+'\n';
         prompt("3 Address Code:",$$);}
    ;

e
    : e '+' e
        {$$ = ['t'+ ++yy.id];
         $$.push($1[1]+$3[1]+$$[0]+' = '+$1[0]+' '+$2+' '+$3[0]+'\n');}
    | e '-' e
        {$$ = ['t'+ ++yy.id];
         $$.push($1[1]+$3[1]+$$[0]+' = '+$1[0]+' '+$2+' '+$3[0]+'\n');}
    | e '*' e
        {$$ = ['t'+ ++yy.id];
         $$.push($1[1]+$3[1]+$$[0]+' = '+$1[0]+' '+$2+' '+$3[0]+'\n');}
    | e '/' e
        {$$ = ['t'+ ++yy.id];
         $$.push($1[1]+$3[1]+$$[0]+' = '+$1[0]+' '+$2+' '+$3[0]+'\n');}
    | '-' e %prec UMINUS
        {$$ = ['t'+ ++yy.id];
         $$.push($2[1]+$$[0]+' = minus '+$2[0]+'\n');}
    | '(' e ')'
        {$$ = $2;}
    | ID
        {$$ = [yytext];
        $$.push('');}
    | NUMBER
        {$$ = [yytext];
        $$.push('');}
    ;
