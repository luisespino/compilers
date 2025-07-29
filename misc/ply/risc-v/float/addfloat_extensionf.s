.data
float1: .float 1.5
float2: .float 1.0
result: .float 0

.text
.globl main

main:
    flw f1, float1, t3
    flw f2, float2, t3
    
    fadd.s f3, f1, f2
    
    fsw f3, result, t3 

    jal ra, print_result
    
    li a7, 10
    ecall

print_result:
    flw fa0, result, t3
    li a7, 2
    ecall
    ret
