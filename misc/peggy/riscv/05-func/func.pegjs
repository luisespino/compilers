{
  function evalExpr(expr) {
    return eval(expr); // Evalúa expresiones aritméticas
  }
}

Program
  = Statement*

Statement
  = FunctionDef
  / IfStatement
  / ReturnStatement
  / Call _ ";"
  / __

FunctionDef
  = "function" _ Name:Identifier _ "(" _ Args:Args _ ")" _ "{" _ Body:Statement* _ "}"

ReturnStatement
  = "return" _ RetVal:Expr _ ";"

IfStatement
  = "if" _ "(" _ Condition:Condition _ ")" _ "{" _ Body:Statement* _ "}" _ ("else" _ "{" _ ElseBody:Statement* _ "}")?

Args
  = Arg ("," _ Arg)*
  / ""

Arg
  = Expr 
  //= Identifier

Condition
  = Left:Expr _ Operator:("==" / "!=" / "<=" / "<" / ">=" / ">") _ Right:Expr

Expr
  = Call
  / Sum
  
Call
  = FunctionName:Identifier _ "(" _ ArgsList:Args? _ ")"


Sum
  = Product _ (("+" / "-") _ Product)*

Product
  = Value _ (("*" / "/") _ Value)*

Value
  = Call 
  / Number
  / Identifier
  / "(" _ Expr _ ")"

Identifier
  = [a-zA-Z_][a-zA-Z0-9_]*

Number
  = [0-9]+

_ = [ \t\n]* // Espacios en blanco

__ = [ \t\n]+ // Espacios en blanco