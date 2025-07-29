# Control Flow Translation to RISC-V
# Luis Espino 2024

# global
addr = 0;
label = 0;
data = '.data\n';
text = '\n.text\n.globl _start\n_start:\n'

# synthesized attributes 
class syn:
    def __init__(self, addr, code, true, false):
        self.addr  = addr   # int or name
        self.code  = code   # code for riscv
        self.true  = true   # true label
        self.false = false  # false label

# auto increment label
def newlabel():
    global label
    label = label + 1
    return label

# Reserved words
reserved = {
    'if'        : 'IF',
    'endif'     : 'ENDIF',
    'else'      : 'ELSE',
    'while'     : 'WHILE',
    'endwhile'  : 'ENDWHILE',
    'true'      : 'TRUE',
    'false'     : 'FALSE',
    'print'     : 'PRINT',
    'int'       : 'INT'
    }

# token names
tokens = [
    'ID', 'NUMBER', 'EQUALS',
    'LPAREN','RPAREN','EQUAL','NEQUAL','OR','AND','NOT',
    'LESS','LESSEQ','GREAT','GREATEQ'
    ] + list(reserved.values())

# token regex
t_EQUAL   = r'=='
t_NEQUAL  = r'!='
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


# id token regex
def t_ID(t):
    r'[a-zA-Z_][a-zA-Z0-9_]*'
    t.type = reserved.get(t.value,'ID')
    return t

# int token regex
def t_NUMBER(t):
    r'\d+'
    try:
        t.value = int(t.value)
    except ValueError:
        print("Integer value too large %d", t.value)
        t.value = 0
    return t

# Ignored characters
t_ignore = " \t"

# newline token regex
def t_newline(t):
    r'\n+'
    t.lexer.lineno += t.value.count("\n")

# error
def t_error(t):
    print("Illegal character '%s'" % t.value[0])
    t.lexer.skip(1)
    
# Build the lexer
import ply.lex as lex
lexer = lex.lex()

# operation precedence
precedence = (
    ('left','OR'),
    ('left','AND'),
    ('right','NOT'),    
    )

# symbol table as dictionary
symbol = { }

# productions

def p_p_s(p):
    'p : s'
    p[0] = p[1].code
    
def p_statement_declaration(p):
    's : INT ID EQUALS expression'
    global data
    symbol[p[2]] = p[4].addr
    data += p[2]+': .word '+str(p[4].addr)+'\n'
    p[0] = syn('declaration', p[4].code, '', '') 

def p_statement_assign(p):
    's : ID EQUALS expression'
    symbol[p[1]] = p[3].addr
    p[3].code += '\t'+'la t0, '+p[1]+'\n'
    p[3].code += '\t'+'li t1, '+str(p[3].addr)+'\n'
    p[3].code += '\t'+'sw t1, 0(t0)\n'
    p[0] = syn('assign', p[3].code, '', '') 

def p_s_if(p):
    's : IF LPAREN b RPAREN s ENDIF'
    p[3].code += p[3].true + p[5].code + p[3].false  
    p[0] = syn('if', p[3].code, '', '')

def p_s_if_else(p):
    's : IF LPAREN b RPAREN s ELSE s ENDIF'
    global label
    newlabel()
    next = 'L'+str(label)
    p[3].code += p[3].true + p[5].code
    p[3].code += '\t'+'j '+next+'\n'
    p[3].code += p[3].false + t[7].code
    next += ':\n'
    p[3].code += next
    p[0] = syn('ifelse', p[3].code, '', '')

def p_s_while(p):
    's : WHILE LPAREN b RPAREN s ENDWHILE'
    global label
    newlabel()
    begin = 'L'+str(label)
    p[3].code = begin+':\n' + p[3].code
    p[3].code += p[3].true + p[5].code
    p[3].code += '\t'+'j '+begin+'\n'
    p[3].code += p[3].false
    p[0] = syn('while', p[3].code, '', '')

def p_statement_print(p):
    's : PRINT LPAREN expression RPAREN'  
    if p[3].addr in symbol:
        p[3].code += '\t'+'lw a0, '+str(p[3].addr)+'\n'
    else:
        p[3].code += '\t'+'li a0, '+str(p[3].addr)+'\n'
    p[3].code += '\t'+'li a7, 1\n'
    p[3].code += '\t'+'ecall\n'
    p[3].code += '\t'+'li a0, 10\n'
    p[3].code += '\t'+'li a7, 11\n'
    p[3].code += '\t'+'ecall\n\n'
    p[0] = syn('print', p[3].code, '', '')    

def p_s_ss(p):
    's : s s'
    p[0] = syn('moresentences', p[1].code+p[2].code, '', '')    

def p_expression_bool(p):
    'expression : b'
    p[0] = syn(str(p[1].addr), p[1].code, p[1].true, p[1].false)    

def p_bool_or(p):
    'b : b OR b'
    p[1].code += p[1].false + p[3].code
    true = p[1].true + p[3].true
    p[0] = syn('or', p[1].code, true, p[3].false) 

def p_bool_and(p):
    'b : b AND b'
    p[1].code += p[1].true + p[3].code
    false = p[1].false + p[3].false
    p[0] = syn('and', p[1].code, p[3].true, false) 

def p_bool_not(p):
    'b : NOT b'
    p[0] = syn('not', p[2].code, p[2].false, p[2].true) 

def p_bool_equal(p):
    '''b : expression EQUAL expression
         | expression NEQUAL expression
         | expression LESS expression
         | expression LESSEQ expression
         | expression GREAT expression         
         | expression GREATEQ expression'''
    global label
    newlabel()
    true = 'L'+str(label)
    newlabel()
    false = 'L'+str(label)
    p[1].code += p[3].code
    if p[1].addr in symbol:
        p[1].code += '\t'+'lw t0, '+str(p[1].addr)+'\n'
    else:
        p[1].code += '\t'+'li t0, '+str(p[1].addr)+'\n'
    if p[3].addr in symbol:
        p[1].code += '\t'+'lw t1, '+str(p[3].addr)+'\n'
    else:
        p[1].code += '\t'+'li t1, '+str(p[3].addr)+'\n'
    if (p[2] == '=='):
        p[1].code += '\t'+'beq t0, t1, '+true+'\n'
    elif (p[2] == '!='):
        p[1].code += '\t'+'bne t0, t1, '+true+'\n'
    elif (p[2] == '<'):
        p[1].code += '\t'+'blt t0, t1, '+true+'\n'
    elif (p[2] == '<='):
        p[1].code += '\t'+'sub t2, t1, t0\n'
        p[1].code += '\t'+'blez t2, '+true+'\n'
    elif (p[2] == '>'):
        p[1].code += '\t'+'bgt t0, t1, '+true+'\n'
    elif (p[2] == '>='):
        p[1].code += '\t'+'bge t0, t1, '+true+'\n'
    p[1].code += '\t'+'j '+false+'\n'
    true += ':\n'
    false += ':\n'
    p[0] = syn('rel', p[1].code, true, false) 

def p_bool_t(p):
    'b : TRUE'
    global label
    newlabel()
    true = 'L'+str(label)
    code = '\t'+'j '+true+'\n'
    true += ':\n'
    p[0] = syn('true', code, true, '')

def p_bool_f(p):
    'b : FALSE'
    global label
    newlabel()
    false = 'L'+str(label)
    code = '\t'+'j '+false+'\n'
    false += ':\n'    
    p[0] = syn('false', code, '', false)

def p_expression_group(p):
    'expression : LPAREN expression RPAREN'
    p[0] = p[2]

def p_expression_int(p):
    'expression : NUMBER'
    p[0] = syn(p[1], '', '', '')

def p_expression_name(p):
    'expression : ID'
    try:
        p[0] = syn(p[1], '', '', '')        
    except LookupError:
        print("Undefined id '%s'" % p[1])
        p[0] = None

def p_error(p):
    print("Syntax error at '%s'" % p.value)

# build the parser
import ply.yacc as yacc
parser = yacc.yacc()        

# input example
code_input = '''
int x = 10
int y = 10
if (x<100 || x>200 && x!=y)
    x = 0
endif
print (x)
'''

# input analysis
res = parser.parse(code_input)
code_output = data+text+res+'\t'+'li a7, 10\n\tecall\n'  
print(code_output)
