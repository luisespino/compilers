grammar Grammar;


WS : [ \t\n\r]+ -> skip;
NAME : [a-zA-Z]+;
DIGIT : [0-9]+;

prog
    : stmt* EOF
    ;

stmt
    : decl
    | assg
    ;

decl
    : 'int' var '=' array ';'
    | 'int' var '=' val ';'
    ;

assg
    : var '=' var index ';'
    ;

var
    : NAME
    ;

array
    : '[' ']'
    | '[' list ']'
    ;
    
val
    : DIGIT
    | array
    ;

index
    : '[' val ']' index
    | '[' val ']'
    ;

list
    : val ',' list
    | val
    ;