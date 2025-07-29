from graphviz import Graph
i = 0

def inc():
    global i
    i += 1
    return i

tokens = ('DIGIT','PLUS','TIMES','LPAREN','RPAREN')

t_PLUS    = r'\+'
t_TIMES   = r'\*'
t_LPAREN  = r'\('
t_RPAREN  = r'\)'

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
    'line : expr'

def p_expr_plus(t):
    'expr : expr PLUS term'    
    id = inc()
    t[0] = id
    dot.node(str(id),str(t[2]))
    dot.edge(str(id),str(t[1]))
    dot.edge(str(id),str(t[3]))
    
def p_expr_term(t):
    'expr : term'
    t[0] = t[1] 

def p_expr_times(t):
    'term : term TIMES factor'
    id = inc()
    t[0] = id
    dot.node(str(id),str(t[2]))
    dot.edge(str(id),str(t[1]))
    dot.edge(str(id),str(t[3]))

def p_expression_times(t):
    'term : factor'
    t[0] = t[1] 

def p_expression_group(t):
    'factor : LPAREN expr RPAREN'
    t[0] = t[2]

def p_expression_digit(t):
    'factor : DIGIT'
    id = inc()
    t[0] = id
    dot.node(str(id),str(t[1]))

def p_error(t):
    print("Syntax error at '%s'" % t.value)

import ply.yacc as yacc
parser = yacc.yacc()

while True:
    i = 0
    dot = Graph()
    dot.attr(splines='false')
    dot.node_attr.update(shape='circle', fontname='arial',
                         color='blue4', fontcolor='blue4')
    dot.edge_attr.update(color='blue4')        
    try:
        s = input('calc > ')
    except EOFError:
        break
    parser.parse(s)
    dot.view()
