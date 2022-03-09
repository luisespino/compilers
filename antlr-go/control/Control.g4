grammar Control;

// tokens

INT     : [0-9]+ ;
ID      : [a-zA-Z_][a-zA-Z0-9_]* ;
STRING  : '"' (~["\r\n] | '""')* '"' ;
WS      : [ \n\t]+ -> skip ;

// rules

prog: block EOF;

block: (stmt)*
     ;

stmt: assignstmt
    | printlnstmt
    | printstmt
    | ifstmt
    | whilestmt
    ;

assignstmt
    : ID '=' expr
    ;

printlnstmt
    : 'println' '(' expr ')'
    ;

printstmt
    : 'print' '(' expr ')'
    ;

ifstmt  
    : 'if' '(' expr ')' '{' block '}'
    ;

whilestmt
    : 'while' '(' expr ')' '{' block '}'
    ;

expr
    : '!' right=expr                        # NotExpr
    | left=expr op=('*'|'/') right=expr     # OpExpr
    | left=expr op=('+'|'-') right=expr     # OpExpr
    | left=expr op=('>='|'>') right=expr    # OpExpr
    | left=expr op=('<='|'<') right=expr    # OpExpr
    | left=expr op=('=='|'!=') right=expr   # OpExpr
    | left=expr '&&' right=expr             # OpExpr
    | left=expr '||' right=expr             # OpExpr
    | '(' expr ')'                          # ParExpr
    | INT                                   # IntExpr
    | ID                                    # IdExpr
    | STRING                                # StrExpr    
    | ('true' | 'false')                    # BoolExpr
    ;
