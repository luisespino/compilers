grammar Grammar;

// Tokens
DIGIT : [0-9]+;
WS    : [ \r\n\t]+ -> skip;

// Rules
l : e EOF;

e : t ep        # ET
  ;

ep : '+' t ep   # Sum
   |            # EpsSum
   ;

t : f tp        # TF
  ;       

tp : '*' f tp   # Mul
   |            # EpsMul
   ;

f : '(' e ')'   # Brace
  | DIGIT       # Digit
  ;
  
