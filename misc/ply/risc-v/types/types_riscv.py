import ply.lex as lex
import ply.yacc as yacc
import struct

data = ''
text = ''

def float_to_hex(f):
    return hex(struct.unpack('<I', struct.pack('<f', f))[0])

# Definición de tokens
reserved = {
    'print': 'PRINT',
    'true': 'TRUE',
    'false': 'FALSE',
    'int': 'INT',
    'float': 'FLOAT',
    'string': 'STRING',
    'char': 'CHAR',
    'boolean': 'BOOLEAN'
}

tokens = [
    'NUMBER',
    'ID',
    'STR',
    'EQUALS',
    'PLUS',
    'MINUS',
    'TIMES',
    'DIVIDE',
    'LPAREN',
    'RPAREN'
] + list(reserved.values()) 

# Expresiones regulares para tokens simples
t_EQUALS = r'='
t_PLUS = r'\+'
t_MINUS = r'-'
t_TIMES = r'\*'
t_DIVIDE = r'/'
t_LPAREN = r'\('
t_RPAREN = r'\)'

# Ignorar espacios, tabulaciones y nueva línea
t_ignore = ' \t\n'

# Definición de tokens más complejos
def t_NUMBER(t):
    r'\d+\.\d+|\d+'
    t.value = float(t.value) if '.' in t.value else int(t.value)
    return t

def t_ID(t):
    r'[a-zA-Z_][a-zA-Z0-9_]*'
    t.type = reserved.get(t.value, 'ID')  # Reemplazar con palabra reservada si es el caso
    return t

def t_STR(t):
    r'\"([^\\\n]|(\\.))*?\"'
    t.value = t.value[1:-1]  # Eliminar comillas dobles
    return t

def t_CHAR(t):
    r'\'.\''
    t.value = t.value[1]  # Obtener el carácter sin comillas
    return t


# Manejo de errores léxicos
def t_error(t):
    print("Carácter ilegal '%s'" % t.value[0])
    t.lexer.skip(1)

# Construcción del analizador léxico
lexer = lex.lex()

# Definición de diccionario para almacenar variables, tipos y valores
variable_dict = {}

# Reglas de sintaxis

def p_statements(p):
    '''statements : statements statement
                  | statement'''
    if len(p) == 3:
        p[0] = p[1] + [p[2]]
    else:
        p[0] = [p[1]]

        
def p_statement(p):
    '''statement : variable_declaration
                 | assignment
                 | print_statement'''
    p[0] = p[1]

def p_variable_declaration(p):
    '''variable_declaration : data_type ID EQUALS expression
                             | data_type ID'''
    global data
    if len(p) == 5:
        variable_dict[p[2]] = {'type': p[1], 'value': p[4]}
        if (p[1]=='int'):
            data += p[2]+': .word '+str(p[4])+'\n'
        elif (p[1]=='float'):            
            data += p[2]+': .word '+float_to_hex(p[4])+'\n'
        elif (p[1]=='string'):
            data += p[2]+': .asciz \"'+p[4]+'\"\n'
        elif (p[1]=='char'):
            data += p[2]+': .byte '+str(ord(p[4]))+'\n'
        elif (p[1]=='boolean'):
            if (p[4]=='true'):
                data += p[2]+': .word 1\n'
            else:
                data += p[2]+': .word 0\n'
    else:
        variable_dict[p[2]] = {'type': p[1]}
        if (p[1]=='int'):
            data += p[2]+': .word 0\n'
        elif (p[1]=='float'):            
            data += p[2]+': .word 0x0\n'
        elif (p[1]=='string'):
            data += p[2]+': .asciz \"\"\n'
        elif (p[1]=='char'):
            data += p[2]+': .byte 0\n'
        elif (p[1]=='boolean'):
            data += p[2]+': .word 0\n'

    print('variable_declaration', p[1], p[2], p[4] if len(p) == 5 else None)
    
    p[0] = ('variable_declaration', p[1], p[2], p[4] if len(p) == 5 else None)

def p_assignment(p):
    '''assignment : ID EQUALS expression'''
    global text
    variable_dict[p[1]]['value'] = p[3]
    print('assignment', p[1], p[3])
    if (variable_dict[p[1]]['type']=='int'):
        text +='\t'+'la t0, '+p[1]+'\n'
        text +='\t'+'li t1, '+str(p[3])+'\n'
        text +='\t'+'sw t1, 0(t0)\n'
    elif (variable_dict[p[1]]['type']=='float'):
        text +='\t'+'la t0, '+p[1]+'\n'
        text +='\t'+'li t1, '+float_to_hex(p[3])+'\n'
        text +='\t'+'sw t1, 0(t0)\n'
    elif (variable_dict[p[1]]['type']=='string'):
        text +='\t'+'la t0, '+p[1]+'\n'
        for caracter in p[3]:
            text +='\t'+'li t1, '+str(ord(caracter))+'\n'
            text +='\t'+'sb t1, 0(t0)\n'
            text +='\t'+'addi t0, t0, 1\n'
        text +='\t'+'li t1, 0\n'
        text +='\t'+'sb t1, 0(t0)\n'
    elif (variable_dict[p[1]]['type']=='char'):
        text +='\t'+'la t0, '+p[1]+'\n'
        text +='\t'+'li t1, '+str(ord(p[3]))+'\n'
        text +='\t'+'sw t1, 0(t0)\n'
    elif (variable_dict[p[1]]['type']=='boolean'):
        text +='\t'+'la t0, '+p[1]+'\n'
        if (p[3]=='true'):
            text +='\t'+'li t1, 1\n'
        else:
            text +='\t'+'li t1, 0\n'
        text +='\t'+'sw t1, 0(t0)\n'            
    p[0] = ('assignment', p[1], p[3])

def p_expression(p):
    '''expression : NUMBER
                  | STR
                  | CHAR
                  | TRUE
                  | FALSE
                  | ID'''
    p[0] = p[1]

def p_data_type(p):
    '''data_type : INT
                 | FLOAT
                 | BOOLEAN
                 | STRING
                 | CHAR'''
    p[0] = p[1]

def p_print_statement(p):
    '''print_statement : PRINT LPAREN ID RPAREN'''
    global text
    if p[3] in variable_dict:
        print(f"print {p[3]}: {variable_dict[p[3]]['type']} = {variable_dict[p[3]].get('value', 'No asignado')}")
        text += '\n'
        if (variable_dict[p[3]]['type']=='int'):
            text +='\t'+'lw a0, '+p[3]+'\n'
            text +='\t'+'li a7, 1\n'
            text +='\t'+'ecall\n'
        elif (variable_dict[p[3]]['type']=='float'):
            text +='\t'+'lw a0, '+p[3]+'\n'
            text +='\t'+'li a7, 34\n'
            text +='\t'+'ecall\n'
        elif (variable_dict[p[3]]['type']=='string'):
            text +='\t'+'la a0, '+p[3]+'\n'
            text +='\t'+'li a7, 4\n'
            text +='\t'+'ecall\n'
        elif (variable_dict[p[3]]['type']=='char'):
            text +='\t'+'lw a0, '+p[3]+'\n'
            text +='\t'+'li a7, 11\n'
            text +='\t'+'ecall\n'
        elif (variable_dict[p[3]]['type']=='boolean'):
            text +='\t'+'lw a0, '+p[3]+'\n'
            text +='\t'+'li a7, 1\n'
            text +='\t'+'ecall\n'
        text +='\t'+'li a0, 10\n'
        text +='\t'+'li a7, 11\n'
        text +='\t'+'ecall\n'
        
    else:
        print(f"print variable '{p[3]}' no definida.")

# Manejo de errores sintácticos
def p_error(p):
    print("Error de sintaxis en '%s'" % p.value)

# Construcción del analizador sintáctico
parser = yacc.yacc()

# Ejemplo de entrada
input_text = '''
int i1
int i2 = 1
float f1
float f2 = 2.5
string s1
string s2 = "Hola mundo"
char c = 'c'
boolean b1
boolean b2 = true
i1 = 2
b1 = false
s1 = "hola"
print(i1)
print(f2)
print(b1)
print(s1)
'''

text = '.text\n.globl _start\n_start:\n'
data = '.data\n'

# Análisis de la entrada
print('=== INTERPRETER ===')
result = parser.parse(input_text)
print('\n\n')
print('=== COMPILER ===')
code_output = data+'\n'+text 
print(code_output)


