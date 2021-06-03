/* lexical grammar */
%lex
%%

\s+                   /* skip whitespace */
[0-9]+("."[0-9]+)?\b  return 'DIGIT'
"*"                   return '*'
<<EOF>>               return 'EOF'
.                     return 'INVALID'

/lex

%start expressions

%% /* language grammar */


expressions
    : t EOF
        { typeof console !== 'undefined' ? console.log($1) : print($1);
          return $1; }
    ;

t
    : f tp
        {$$ = $2;}
    ;

tp
    : '*' f tp
        {$$ = $3*$2;}
    |     
        {$$ = $-1;} //only 3*5
        
    ;

f
    : DIGIT
        {$$ = Number(yytext);}
    ;
