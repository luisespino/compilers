
grammar Grammar;

WS : [ \t\n\r]+ -> skip;
NAME : [a-zA-Z]+;
DIGIT : [0-9]+;

prog
    : stmt* EOF                     # Program
    ;

stmt
    : decl                          # StmtDeclaration
    | assg                          # StmtAssigment
    ;

decl
    : type='int' var '=' expr ';'   # DeclExpr
    | type='int' var '=' DIGIT ';'  # DeclVal
    ;

assg
    : var '=' var index ';'         # AssgArray
    ;

var
    : NAME                          # VarName
    ;

expr
    : '[' ']'                       # ExprEmpty
    | '[' list ']'                  # ExprValues
    ;
    
val
    : DIGIT                         # ValDigit
    | expr                          # ValExpr
    ;

index
    : '[' val ']' index             # IndexMany
    | '[' val ']'                   # IndexOne
    ;

list
    : val ',' list                  # Values
    | val                           # Value
    ;
