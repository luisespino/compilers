grammar Calc;

// Tokens
MUL: '*';
ADD: '+';
LB: '(';
RB: ')';
DIGIT: [0-9]+;
WS: [ \r\n\t]+ -> skip;

// Rules
l : e EOF;

e : e '+' t  # Sum
  | t        # PassT
  ;

t : t '*' f  # Mul
  | f        # PassF
  ;

f : '(' e ')' # PassE
  | DIGIT     # Digit
  ;

