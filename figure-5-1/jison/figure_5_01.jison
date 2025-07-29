// MIT License
// Copyright (c) 2021 Luis Espino

/* lexical grammar */
%lex
%%

\s+                   /* skip whitespace */
[0-9]+("."[0-9]+)?\b  return 'DIGIT'
"*"                   return '*'
"+"                   return '+'
"("                   return '('
")"                   return ')'
<<EOF>>               return 'EOF'
.                     return 'INVALID'

/lex

%start expressions

%% /* language grammar */

expressions
    : e EOF
        { typeof console !== 'undefined' ? console.log($1) : print($1);
          return $1; }
    ;

e
    : e '+' t
        {$$ = $1+$3;}
    | t
        {$$ = $1;}
    ;

t
    : t '*' f
        {$$ = $1*$3;}
    | f
        {$$ = $1;}
    ;

f
    : '(' e ')'
        {$$ = $2;}
    | DIGIT
        {$$ = Number(yytext);}
    ;
