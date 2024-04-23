import ply.lex as lex
import ply.yacc as yacc

data = '.data\n';
text = '\n.text\n.globl _start\n_start:\n'
textfunc = '\t'+'li a7, 10\n\tecall\n'
symbol = { }

reserved = {
    'int': 'INT',
    'print': 'PRINT',
    'return': 'RETURN',
}

tokens = [
    'ID',
    'NUMBER',
    'SUM',
    'EQUALS',
    'LPAREN',
    'RPAREN',
    'COMMA',
    'LBRACE',
    'RBRACE'
] + list(reserved.values())

t_SUM = r'\+'
t_EQUALS = r'='
t_LPAREN = r'\('
t_RPAREN = r'\)'
t_COMMA = r','
t_LBRACE = r'{'
t_RBRACE = r'}'

def t_NUMBER(t):
    r'\d+'
    t.value = int(t.value)
    return t

def t_ID(t):
    r'[a-zA-Z_][a-zA-Z0-9_]*'
    t.type = reserved.get(t.value, 'ID')  # Check for reserved words
    return t

t_ignore = ' \t\n'

def t_error(t):
    print("Illegal character '%s'" % t.value[0])
    t.lexer.skip(1)

# Construcción del lexer
lexer = lex.lex()

# Reglas de gramática
def p_sentence(p):
    'p : s'
    p[0] = p[1]

def p_declaration(p):
    's : INT ID EQUALS expr'
    global data
    data += p[2]+': .word '+str(p[4])+'\n'
    symbol[p[2]] = p[4]  
    p[0] = ('declaration', p[2], p[4])

def p_print(p):
    '''s : PRINT LPAREN expr RPAREN
         | PRINT LPAREN funcall RPAREN'''
    global text
    if p[3] in symbol:
        text += '\t'+'lw a0, '+str(p[3])+'\n'
    elif p[3] == 'a0':
        pass
    else:
        text += '\t'+'li a0, '+str(p[3])+'\n'
    text += '\t'+'li a7, 1\n'
    text += '\t'+'ecall\n'
    text += '\t'+'li a0, 10\n'
    text += '\t'+'li a7, 11\n'
    text += '\t'+'ecall\n\n'
    p[0] = ('print', p[3])

def p_functioncall(p):
    'funcall : ID LPAREN fparam RPAREN'
    global text
    i = 0
    while i < len(p[3]):
        text += '\t'+'li a'+str(i)+', '+str(p[3][i])+'\n'
        i = i + 1
    text += '\t'+'jal ra, '+p[1]+'\n'
    p[0] = 'a0'

def p_function(p):
    's : function'
    p[0] = p[1]

def p_function_declare(p):
    'function : INT ID LPAREN fparam RPAREN LBRACE fblock RBRACE'
    global textfunc
    textfunc += p[2]+':\n'
    textfunc += p[7]
    p[0] = ('function', p[2], p[4], p[7])

def p_fparam(p):
    '''fparam : ID
              | ID COMMA fparams
              | empty '''
    if len(p) == 2:
        p[0] = [symbol[p[1]]] if p[1] is not None else []
    elif len(p) == 4:
        p[0] = [symbol[p[1]]] + p[3]

def p_fparams(p):
    '''fparams : ID
               | ID COMMA fparams'''
    if len(p) == 2:
        p[0] = [symbol[p[1]]]
    elif len(p) == 4:
        p[0] = [symbol[p[1]]] + p[3]

def p_fblock(p):
    '''fblock : s RETURN expr 
              | RETURN expr '''
    if len(p) == 4:
        p[0] = p[3]
        p[0] += '\t'+'mv a0, t0\n'
        p[0] += '\t'+'ret\n'        
    elif len(p) == 3:
        p[0] = p[2]
        p[0] += '\t'+'mv a0, t0\n'
        p[0] += '\t'+'ret\n'

def p_sentences(p):
    's : s s'
    p[0] = p[1] + p[2]

def p_expr(p):
    'expr : expr SUM expr'
    # missing parameter substitution
    global text
    p[0] = '\t'+'add t0, a0, a1\n'

def p_expr_number(p):
    'expr : NUMBER'
    p[0] = p[1]

def p_expr_id(p):
    'expr : ID'
    p[0] = p[1]

def p_empty(p):
    'empty :'
    pass

def p_error(p):
    print("Error de sintaxis en '%s'" % p.value)


parser = yacc.yacc()


input_code = '''
int a = 1
int b = 2
print(sum(a, b))

int sum(a, b){
    return a+b
}
'''

parser.parse(input_code)
output = data + text + textfunc
print(output)
