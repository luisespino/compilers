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
    : expr ('+'|'-') term               # ExprSumRes
    | term                              # ExprTerm
    ;
    
term
    : term ('*'|'/') fact               # ExprMulDiv
    | fact                              # ExprFact
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
    ;

or
    : and '||' or                       # OrExpr
    | and                               # OrPass
    ;

and
    : not '&&' and                      # AndExpr
    | not                               # AndPass
    ;

not
    : '!' not                          	# NotExpr
    | rel                               # NotPass
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

