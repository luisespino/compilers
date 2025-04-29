grammar Grammar;

prog
  : stmt* EOF
  ;

stmt
  : fun
  | call SEMI
  ;

fun
  : id LPAREN args? RPAREN LBRACE ret RBRACE
  ;

ret
  : 'return' expr SEMI
  ;

args
  : id (COMMA id)*
  ;

call
  : id LPAREN argsend? RPAREN
  ;
  
argsend
  : num (COMMA num)*
  ;

expr
  : id
  | num
  ;

id
  : IDENTIFIER
  ;
  
num
  : NUMBER
  ;

SEMI    : ';';
COMMA   : ',';
LPAREN  : '(';
RPAREN  : ')';
LBRACE  : '{';
RBRACE  : '}';
PLUS    : '+';
MINUS   : '-';
MUL     : '*';
DIV     : '/';

NUMBER
  : [0-9]+
  ;

IDENTIFIER
  : [a-zA-Z_][a-zA-Z0-9_]*
  ;
  
WS : [ \t\n\r\f]+ -> skip;


