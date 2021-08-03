# Figure 5.4
# Luis Espino 2021

tokens = ('DIGIT','TIMES')

t_TIMES   = r'\*'

def t_DIGIT(t):
    r'\d+'
    try:
        t.value = int(t.value)
    except ValueError:
        print("Integer value too large %d", t.value)
        t.value = 0
    return t

t_ignore = " \t"

def t_newline(t):
    r'\n+'
    t.lexer.lineno += t.value.count("\n")
    
def t_error(t):
    print("Illegal character '%s'" % t.value[0])
    t.lexer.skip(1)
    
import ply.lex as lex
lexer = lex.lex()

def p_line(t):
    'line : t'
    print(t[1])

def p_t_f_tp(t):
    't : f tp'
    t[0] = t[2]
    print ("t: f tp, tp_sin:"+"{}".format(t[0]))
    
def p_tp_times_f_tp(t):
    'tp : TIMES f tp'
    t[0] = t[3]*t[-1]
    print ("tp: TIMES f tp, tp_sin:{} = tp_her:{} x f_val_sin:{} ".format(t[0], t[-1], t[3]))

def p_tp_empty(t):
    'tp :'
    t[0] = t[-1]
    print ("tp: empty, tp_sin=tp_her:{}".format(t[0]))

def p_expression_digit(t):
    'f : DIGIT'
    t[0] = t[1]
    print ("f: DIGIT, f_val_sin:{}".format(t[0]))

def p_error(t):
    print("Syntax error at '%s'" % t.value)

import ply.yacc as yacc
parser = yacc.yacc()

while True:
    try:
        s = input('calc > ')
    except EOFError:
        break
    parser.parse(s)
