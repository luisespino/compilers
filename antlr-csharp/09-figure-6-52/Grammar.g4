grammar Grammar;

prog
  : stmt* EOF
  ;

stmt
  : fun
  | ret SEMI
  | call SEMI
  ;

fun
  : id LPAREN args? RPAREN LBRACE stmt* RBRACE
  ;

ret
  : 'return' expr
  ;

args
  : expr (COMMA expr)*
  ;

call
  : id LPAREN args? RPAREN
  ;

expr
  : term ((PLUS | MINUS) term)*
  ;

term
  : fact ((MUL | DIV) fact)*
  ;

fact
  : call
  | id
  | NUMBER
  | LPAREN expr RPAREN
  ;

id
  : IDENTIFIER
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


