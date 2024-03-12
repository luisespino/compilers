# Declaración, inicialización y acceso a arreglos multidimensionales
# Interpretando código y traduciendo a assembly RISC-V
# Luis Espino 2024

import ply.lex as lex
import ply.yacc as yacc

data = ''
text = ''

# Definición de tokens
reserved = {
    'print': 'PRINT'
}

tokens = [
    'ID',
    'NUM',
    'LBRACKET',
    'RBRACKET',
    'LPAREN',
    'RPAREN',
    'COMMA',
    'ASSIGN',
] + list(reserved.values())

# Reglas de expresiones regulares para tokens simples
t_LBRACKET = r'\['
t_RBRACKET = r'\]'
t_LPAREN = r'\('
t_RPAREN = r'\)'
t_COMMA = r','
t_ASSIGN = r'='

# Definición de token ID
def t_ID(t):
    r'[a-zA-Z_][a-zA-Z0-9_]*'
    t.type = reserved.get(t.value, 'ID')  # Check for reserved words
    return t

# Definición de token NUM
def t_NUM(t):
    r'\d+'
    t.value = int(t.value)
    return t

# Ignorar espacios, tabulaciones y saltos de línea
t_ignore = ' \t\n'

# Manejo de errores
def t_error(t):
    print("Carácter ilegal '%s'" % t.value[0])
    t.lexer.skip(1)

# Construcción del lexer
lexer = lex.lex()

# Función auxiliar para calcular las dimensiones de una lista
def get_dimensions(lst):
    if isinstance(lst, list):
        if lst:
            return [len(lst)] + get_dimensions(lst[0])
        else:
            return [0]
    else:
        return []

# Diccionario para almacenar las variables asignadas
variables = {}

# Reglas de producción
def p_statements(p):
    '''statements : statement
                  | statements statement'''
    if len(p) == 2:
        p[0] = [p[1]]
    else:
        p[0] = p[1] + [p[2]]

def p_statement_assign(p):
    'statement : ID ASSIGN expression'
    global data
    dimensions = get_dimensions(p[3])
    print(f'{p[1]} = {p[3]} (dimensiones: {dimensions})')
    variables[p[1]] = p[3]
    value = flatten_array(variables.get(p[1]))
    data += p[1] + ': .word '
    nums = ', '.join(map(str, value))
    if (nums==''):
        data += '0\n'
    else:
        data += nums + '\n'
    if (len(dimensions) == 1):
        data += p[1] + '_cols: .word '+str(dimensions[0])+'\n'
    elif (len(dimensions) == 2):
        data += p[1] + '_rows: .word '+str(dimensions[0])+'\n'
        data += p[1] + '_cols: .word '+str(dimensions[1])+'\n'    
    elif (len(dimensions) == 3):
        data += p[1] + '_face: .word '+str(dimensions[0])+'\n'
        data += p[1] + '_rows: .word '+str(dimensions[1])+'\n'
        data += p[1] + '_cols: .word '+str(dimensions[1])+'\n'    
    p[0] = ('assign', p[1], p[3], dimensions)

def p_statement_print(p):
    'statement : PRINT LPAREN ID indices RPAREN'  # Nueva producción para imprimir valores
    global text
    value = variables.get(p[3])  # Obtener el valor del arreglo
    dimensions = get_dimensions(value)
    if value:
        for index in p[4]:
            value = value[index]
        print(f'{p[3]}{p[4]} = {value}')
        if (len(dimensions) == 1): # [COL]
            text += '\t'+'li t0, '+str(index)+'\n'
            text += '\t'+'slli t0, t0, 2\n'
            text += '\t'+'la t1, '+p[3]+'\n'
            text += '\t'+'add t2, t1, t0\n'
            text += '\t'+'lw a0, 0(t2)\n'
            text += '\t'+'li a7, 1\n'
            text += '\t'+'ecall\n'
            text += '\t'+'li a0, 10\n'
            text += '\t'+'li a7, 11\n'
            text += '\t'+'ecall\n\n'
        elif (len(dimensions) == 2): #[ROW][COL]
            text += '\t'+'li t0, '+str(p[4][0])+'\n'
            text += '\t'+'slli t0, t0, 2\n'
            text += '\t'+'li t1, '+str(p[4][1])+'\n'
            text += '\t'+'slli t1, t1, 2\n'
            text += '\t'+'lw t2, '+p[3]+'_cols\n'
            text += '\t'+'mul t3, t0, t2\n'
            text += '\t'+'add t3, t3, t1\n'
            text += '\t'+'la t4, '+p[3]+'\n'
            text += '\t'+'add t4, t4, t3\n'
            text += '\t'+'lw a0, 0(t4)\n'
            text += '\t'+'li a7, 1\n'
            text += '\t'+'ecall\n'
            text += '\t'+'li a0, 10\n'
            text += '\t'+'li a7, 11\n'
            text += '\t'+'ecall\n\n'
        elif (len(dimensions) == 3): #[FACE][ROW][COL]
            text += '\t'+'li t0, '+str(p[4][0])+'\n'
            text += '\t'+'slli t0, t0, 2\n'
            text += '\t'+'li t1, '+str(p[4][1])+'\n'
            text += '\t'+'slli t1, t1, 2\n'
            text += '\t'+'li t2, '+str(p[4][2])+'\n'
            text += '\t'+'slli t2, t2, 2\n'
            text += '\t'+'lw t3, '+p[3]+'_rows\n'
            text += '\t'+'mul t4, t3, t0\n'
            text += '\t'+'add t4, t4, t1\n'
            text += '\t'+'lw t5, '+p[3]+'_cols\n'
            text += '\t'+'mul t6, t5, t4\n'
            text += '\t'+'add t6, t6, t2\n'            
            text += '\t'+'la t0, '+p[3]+'\n'
            text += '\t'+'add t0, t0, t6\n'
            text += '\t'+'lw a0, 0(t0)\n'
            text += '\t'+'li a7, 1\n'
            text += '\t'+'ecall\n'
            text += '\t'+'li a0, 10\n'
            text += '\t'+'li a7, 11\n'
            text += '\t'+'ecall\n\n'
    else:
        print(f'Error: El arreglo {p[3]} no está definido')

def p_indices_single(p):
    '''indices : LBRACKET NUM RBRACKET'''
    p[0] = [p[2]]

def p_indices_multiple(p):
    '''indices : LBRACKET NUM RBRACKET  indices'''
    p[0] = [p[2]] + p[4]

def p_expression_empty_list(p):
    'expression : LBRACKET RBRACKET'
    p[0] = []

def p_expression_list(p):
    'expression : LBRACKET list_elements RBRACKET'
    p[0] = p[2]

def p_list_elements_single(p):
    'list_elements : value'
    p[0] = [p[1]]

def p_list_elements_multiple(p):
    'list_elements : list_elements COMMA value'
    p[0] = p[1] + [p[3]]

def p_value_num(p):
    'value : NUM'
    p[0] = p[1]

def p_value_list(p):
    'value : expression'
    p[0] = p[1]

def p_error(p):
    print("Error de sintaxis en '%s'" % p.value)

def flatten_array(arr):
    result = []
    def flatten_helper(arr):
        for item in arr:
            if isinstance(item, list):
                flatten_helper(item)
            else:
                result.append(item)
    flatten_helper(arr)
    return result

# Construcción del parser
parser = yacc.yacc()

# Ejemplo de entrada
code_input = '''
a = []
b = [1, 2, 3]
c = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
d = [[[1, 2], [3, 4]], [[5, 6],[7, 8]]]
print(b[1])
print(c[2][2])
print(d[1][0][0])
'''
text = '.text\n.globl _start\n_start:\n'
data = '.data\n'

# Análisis de la entrada
print('=== INTERPRETER ===')
result = parser.parse(code_input)
print('\n\n')
print('=== COMPILER ===')
code_output = data+'\n'+text 
print(code_output)
