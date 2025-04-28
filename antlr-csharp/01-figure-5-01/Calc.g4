grammar Calc;

// tokens
NUM : [0-9]+;
WS  : [ \r\n\t]+ -> skip;

// rules
e : e op=('+'|'-') t    # Sumres
  | t                   # PassT
  ;

t : t op=('*'|'/') f    # Muldiv
  | f                   # PassF
  ;

f : '(' e ')'           # PassE
  | NUM                 # Num
  ;
