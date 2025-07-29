grammar Grammar;

// tokens
ID  : [a-zA-Z_][a-zA-Z0-9_]*;
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
  | ID                  # Id 
  | NUM                 # Num
  ;
