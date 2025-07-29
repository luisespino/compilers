class tac:
    def __init__(self, addr, code, type, width, true, false):
        self.addr  = addr
        self.code  = code
        self.type  = type
        self.width = width
        self.true  = true
        self.false = false

reserved = {
    'if'    : 'IF',
    'else'  : 'ELSE',
    'while' : 'WHILE',
    'true'  : 'TRUE',
    'false' : 'FALSE'
    }

tokens = [
    'NAME','FLOAT','INT',
    'PLUS','MINUS','TIMES','DIVIDE','EQUALS',
    'LPAREN','RPAREN','EQUAL','NEQUAL','OR','AND','NOT',
    'LESS','LESSEQ','GREAT','GREATEQ'
    ] + list(reserved.values())

# Tokens

t_EQUAL   = r'=='
t_NEQUAL  = r'!='
t_PLUS    = r'\+'
t_MINUS   = r'-'
t_TIMES   = r'\*'
t_DIVIDE  = r'/'
t_EQUALS  = r'='
t_LPAREN  = r'\('
t_RPAREN  = r'\)'
t_OR      = r'\|\|'
t_AND     = r'&&'
t_NOT     = r'!'
t_LESS    = r'<'
t_LESSEQ  = r'<='
t_GREAT   = r'>'
t_GREATEQ = r'>='


def t_NAME(t):
    r'[a-zA-Z_][a-zA-Z0-9_]*'
    t.type = reserved.get(t.value,'NAME')
    return t

def t_FLOAT(t):
    r'\d+\.\d+'
    try:
        t.value = float(t.value)
    except ValueError:
        print("Float value too large %d", t.value)
        t.value = 0
    return t

def t_INT(t):
    r'\d+'
    try:
        t.value = int(t.value)
    except ValueError:
        print("Integer value too large %d", t.value)
        t.value = 0
    return t

# Ignored characters
t_ignore = " \t"

def t_newline(t):
    r'\n+'
    t.lexer.lineno += t.value.count("\n")
    
def t_error(t):
    print("Illegal character '%s'" % t.value[0])
    t.lexer.skip(1)
    
# Build the lexer
import ply.lex as lex
lexer = lex.lex()

# Parsing rules

precedence = (
    ('left','OR'),
    ('left','AND'),
    ('right','NOT'),    
    ('left','PLUS','MINUS'),
    ('left','TIMES','DIVIDE'),
    ('right','UMINUS'),
    )

# dictionary of names
names = { }
types = { }

def p_p_s(t):
    'p : s'
    print(t[1].code)

def p_statement_assign(t):
    's : NAME EQUALS expression'
    names[t[1]] = t[3].addr
    types[t[1]] = t[3].type
    c = t[3].code + t[1]+' = '+str(t[3].addr)+'\n'
    t[0] = tac ('assign',c,t[3].type,1,'','') 

#production in C and C++ but not in Java
""" def p_statement_expr(t):
    's : expression'
    t[0] = t[1]"""

def p_s_if(t):
    's : IF LPAREN b RPAREN s'
    c = t[3].code
    c += t[3].true
    c += t[5].code
    c += t[3].false
    t[0] = tac('if',c,'',0,'','')

def p_s_if_else(t):
    's : IF LPAREN b RPAREN s ELSE s'
    newlabel()
    next = 'L'+str(label)
    c = t[3].code
    c += t[3].true
    c += t[5].code
    c += 'goto '+next+'\n'
    c += t[3].false
    c += t[7].code
    next += ':\n'
    c += next
    t[0] = tac('if',c,'',0,'','')

def p_s_while(t):
    's : WHILE LPAREN b RPAREN s'
    newlabel()
    begin = 'L'+str(label)
    c = begin+':\n'
    c += t[3].code
    c += t[3].true
    c += t[5].code
    c += 'goto '+begin+'\n'
    c += t[3].false
    t[0] = tac('while',c,'',0,'','')

def p_s_ss(t):
    's : s s'
    t[0] = tac('ss',t[1].code+t[2].code,'',0,'','')
    

def p_expression_bool(t):
    'expression : b'
    t[0] = tac (str(t[1].addr),t[1].code,'BOOLEAN',1,t[1].true,t[1].false)    

def p_bool_or(t):
    'b : b OR b'
    c = t[1].code 
    c += t[1].false
    c += t[3].code
    lt = t[1].true+t[3].true
    t[0] = tac ('log',c,'BOOLEAN',1,lt,t[3].false) 

def p_bool_and(t):
    'b : b AND b'
    c = t[1].code 
    c += t[1].true
    c += t[3].code
    lf = t[1].false+t[3].false
    t[0] = tac ('log',c,'BOOLEAN',1,t[3].true,lf) 

def p_bool_not(t):
    'b : NOT b'
    t[0] = tac ('log',t[2].code,'BOOLEAN',1,t[2].false,t[2].true) 

def p_bool_equal(t):
    '''b : expression EQUAL expression
         | expression NEQUAL expression
         | expression LESS expression
         | expression LESSEQ expression
         | expression GREAT expression         
         | expression GREATEQ expression'''
    newlabel()
    lt = 'L'+str(label)
    newlabel()
    lf = 'L'+str(label)
    c = t[1].code + t[3].code
    c += 'if '+str(t[1].addr)+' '+t[2]+' '+str(t[3].addr)+' goto '+lt+'\n'
    c += 'goto '+lf+'\n'
    lt += ':\n'
    lf += ':\n'
    t[0] = tac ('rel',c,'BOOLEAN',1,lt,lf) 


def p_bool_t(t):
    'b : TRUE'
    newlabel()
    lt = 'L'+str(label)
    c = 'goto '+lt+'\n'
    lt += ':\n'
    t[0] = tac('true',c,'BOOLEAN',1,lt,'')

def p_bool_f(t):
    'b : FALSE'
    newlabel()
    lf = 'L'+str(label)
    c = 'goto '+lf+'\n'
    lf += ':\n'    
    t[0] = tac('false',c,'BOOLEAN',1,'',lf)

def p_expression_binop(t):
    '''expression : expression PLUS expression
                  | expression MINUS expression
                  | expression TIMES expression
                  | expression DIVIDE expression'''
    #type max
    type = ""
    width = 0
    if t[1].width>=t[3].width:
        type = t[1].type
        width = t[1].width
    else:
        type = t[3].type
        width = t[3].width
    #widen
    widen = ""
    if t[1].type=="INT" and type=="FLOAT":
        newtemp()
        a = 't'+str(addr)
        widen = a+" = (float) "+str(t[1].addr) + "\n"
        t[1].addr = a
    elif t[3].type=="INT" and type=="FLOAT":
        newtemp()
        a = 't'+str(addr)
        widen = a+" = (float) "+str(t[3].addr) + "\n"
        t[3].addr = a
    #addr
    newtemp()
    a = 't'+str(addr)
    #code
    code =  a+' = '+str(t[1].addr)+' '+t[2]+' '+str(t[3].addr) + "\n"
    code = t[1].code + t[3].code + widen + code
    #instance
    t[0] = tac(a, code, type, width,'','')

def p_expression_uminus(t):
    'expression : MINUS expression %prec UMINUS'
    #t[0] = -t[2]
    newtemp()
    t[0] = 't'+str(addr)
    print (t[0]+' = minus '+str(t[2]))    

def p_expression_group(t):
    'expression : LPAREN expression RPAREN'
    t[0] = t[2]

def p_expression_float(t):
    'expression : FLOAT'
    t[0] = tac(t[1],"","FLOAT",8,'','')

def p_expression_int(t):
    'expression : INT'
    t[0] = tac(t[1],"","INT",4,'','')

def p_expression_name(t):
    'expression : NAME'
    try:
        t[0] = names[t[1]]
        t[0] = tac(t[1],"",types[t[1]],4,'','')        
    except LookupError:
        print("Undefined name '%s'" % t[1])
        t[0] = 0

def p_error(t):
    print("Syntax error at '%s'" % t.value)

import ply.yacc as yacc
parser = yacc.yacc()

def newtemp():
    global addr
    addr = addr + 1
    return addr
        

addr = 0;

def newlabel():
    global label
    label = label + 1
    return label
        
label = 0;

while True:
    try:
        addr = 0;
        label = 0;
        s = input('calc > ')   # Use raw_input on Python 2
    except EOFError:
        break
    parser.parse(s)
