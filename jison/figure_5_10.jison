/* description: generate the DOT string to create the AST of a expression. */

/* lexical grammar */
%lex

%{
if (!('id' in yy)) {
  yy.id = 0;
  yy.dot = 'graph{'+'\n';
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
"("                   return '('
")"                   return ')'
<<EOF>>               return 'EOF'
.                     return 'INVALID'

/lex

/* operator associations and precedence */

%left '+' '-'
%left '*' '/'

%start expressions

%% /* language grammar */

expressions
    : e EOF
        {prompt("DOT:",yy.dot+'}');}
    ;

e
    : e '+' t
        {$$ = ++yy.id;
        yy.dot += yy.id+' [label="'+$2+'"];'+'\n';
        yy.dot += yy.id+'--'+$1+';'+'\n';
        yy.dot += yy.id+'--'+$3+';'+'\n';}
    | e '-' t
        {$$ = ++yy.id;
        yy.dot += yy.id+' [label="'+$2+'"];'+'\n';
        yy.dot += yy.id+'--'+$1+';'+'\n';
        yy.dot += yy.id+'--'+$3+';'+'\n';}
    | t
        {$$ = $1;}
    ;
t
    : t '*' f
        {$$ = ++yy.id;
        yy.dot += yy.id+' [label="'+$2+'"];'+'\n';
        yy.dot += yy.id+'--'+$1+';'+'\n';
        yy.dot += yy.id+'--'+$3+';'+'\n';}
    | t '/' f
        {$$ = ++yy.id;
        yy.dot += yy.id+' [label="'+$2+'"];'+'\n';
        yy.dot += yy.id+'--'+$1+';'+'\n';
        yy.dot += yy.id+'--'+$3+';'+'\n';}
    | f
        {$$ = $1;}
    ;
f
    : '(' e ')'
        {$$ = $2;}
    | ID
        {$$ = ++yy.id;
        yy.dot += yy.id+' [label="'+yytext+'"];'+'\n';}
    | NUMBER
        {$$ = ++yy.id;
        yy.dot += yy.id+' [label="'+yytext+'"];'+'\n';}
    ;
