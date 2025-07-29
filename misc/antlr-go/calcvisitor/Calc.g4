grammar Calc;

// tokens
INT : [0-9]+ ;
WS  : [ \t]+ -> skip ;

// rules
expr:	left=expr op=('*'|'/') right=expr  # Op
    |	left=expr op=('+'|'-') right=expr  # Op
    |	digit=INT                          # Digit
    |	'(' expr ')'                       # Paren
    ;
