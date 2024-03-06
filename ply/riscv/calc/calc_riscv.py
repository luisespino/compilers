# Calc translator to RISC-V assembler
# (Synthesizing temporary addresses)
# Luis Espino 2024

t = 0

def get_temp():
    global t
    t += 4    # offset 4 bytes
    return t

tokens = (
    'NUMBER',
    'PLUS','MINUS','TIMES','DIVIDE','EQUALS',
    'LPAREN','RPAREN',
    )

# Tokens

t_PLUS    = r'\+'
t_MINUS   = r'-'
t_TIMES   = r'\*'
t_DIVIDE  = r'/'
t_EQUALS  = r'='
t_LPAREN  = r'\('
t_RPAREN  = r'\)'

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
    ('left','PLUS','MINUS'),
    ('left','TIMES','DIVIDE'),
    ('right','UMINUS'),
    )

def p_statement_expr(t):
    'statement : expression'
    print()
    print('\t'+'li t3, '+str(t[1]))
    print('\t'+'lw a0, 0(t3)')
    print('\t'+'li a7, 1')
    print('\t'+'ecall\n')
    print('\t'+'li a0, 0')
    print('\t'+'li a7, 93')
    print('\t'+'ecall\n')

def p_expression_binop(t):
    '''expression : expression PLUS expression
                  | expression MINUS expression
                  | expression TIMES expression
                  | expression DIVIDE expression'''
    addr = get_temp()
    print()
    print('\t'+'li t3, '+str(t[1]))
    print('\t'+'lw t1, 0(t3)')
    print('\t'+'li t3, '+str(t[3]))
    print('\t'+'lw t2, 0(t3)')
    if t[2] == '+'  : print('\t'+'add t0, t1, t2')
    elif t[2] == '-': print('\t'+'sub t0, t1, t2')
    elif t[2] == '*': print('\t'+'mul t0, t1, t2')
    elif t[2] == '/': print('\t'+'div t0, t1, t2')
    print('\t'+'li t3, '+str(addr))
    print('\t'+'sw t0, 0(t3)')
    t[0] = addr   

def p_expression_uminus(t):
    'expression : MINUS expression %prec UMINUS'
    print()
    print('\t'+'li t3, '+str(t[2]))
    print('\t'+'lw t1, 0(t3)')
    print('\t'+'neg t0, t1')
    print('\t'+'sw t0, 0(t3)')
    t[0] = t[2]

def p_expression_group(t):
    'expression : LPAREN expression RPAREN'
    t[0] = t[2]

def p_expression_number(t):
    'expression : NUMBER'
    addr = get_temp()
    print()
    print('\t'+'li t0, '+str(t[1]))
    print('\t'+'li t3, '+str(addr))
    print('\t'+'sw t0, 0(t3)')
    t[0] = addr

def p_error(t):
    print("Syntax error at '%s'" % t.value)

import ply.yacc as yacc
parser = yacc.yacc()

while True:
    t = 0
    try:
        s = input('calc > ')   # Use raw_input on Python 2
        print('.globl _start\n_start:')
    except EOFError:
        break
    parser.parse(s)
