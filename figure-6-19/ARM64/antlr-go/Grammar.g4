grammar Grammar;

// tokens
ID  : [a-zA-Z_][a-zA-Z0-9_]*;
NUM : [0-9]+;
WS  : [ \r\n\t]+ -> skip;

// rules

p : s+ EOF              # Program
  ;

s : ID '=' e            # Assign
  ;

e : e op=('*'|'/') e    # Muldiv
  | e op=('+'|'-') e    # Sumres
  | '(' e ')'           # PassE
  | ID                  # Id 
  | NUM                 # Num
  ;