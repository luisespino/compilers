grammar Grammar;

prog
    : stmt* EOF
    ;

stmt
    : decl
    | assg
    | if
    | while
    ;

decl
    : type='int' var '=' expr ';'
    ;

var
    : NAME
    ;

expr
    : expr ('+'|'-') term
    | term
    ;
    
term
    : term ('+'|'-') fact
    | fact
    ;
    
fact
    : '(' expr ')'
    | var
    | num
    ;

num
    : DIGIT
    ;

assg
    : var '=' expr ';'
    ;

if
    : 'if' cond stmt* 'endif'
    ;
    
while
    : 'while' cond stmt* 'endwhile'
    ;

    
cond
    : '(' bool ')'
    ;

bool
    : or
    | and
    | not
    | rel
    | T
    | F
    ;

or
    : expr '||' expr
    ;

and
    : expr '&&' expr
    ;

not
    : '!' expr
    ;

rel
    : expr oprel expr
    ;
    
oprel 
    : EQ 
    | NEQ 
    | LE 
    | LT 
    | GE 
    | GT
    ;
    

NAME : [a-zA-Z_][a-zA-Z_0-9]* ;

DIGIT : [0-9]+;

WS : [ \t\n\r\f]+ -> skip;

EQ  : '==';
NEQ : '!=';
LE  : '<=';
LT  : '<';
GE  : '>=';
GT  : '>';

T : 'true';
F : 'false';

