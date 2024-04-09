.data
i1: .word 0
i2: .word 1
f1: .word 0x0
f2: .word 0x40200000
s1: .asciz ""
s2: .asciz "Hola mundo"
c: .byte 99
b1: .word 0
b2: .word 1

.text
.globl _start
_start:
	la t0, i1
	li t1, 2
	sw t1, 0(t0)
	la t0, b1
	li t1, 0
	sw t1, 0(t0)
	la t0, s1
	li t1, 104
	sb t1, 0(t0)
	addi t0, t0, 1
	li t1, 111
	sb t1, 0(t0)
	addi t0, t0, 1
	li t1, 108
	sb t1, 0(t0)
	addi t0, t0, 1
	li t1, 97
	sb t1, 0(t0)
	addi t0, t0, 1
	li t1, 0
	sb t1, 0(t0)

	lw a0, i1
	li a7, 1
	ecall
	li a0, 10
	li a7, 11
	ecall

	lw a0, f2
	li a7, 34
	ecall
	li a0, 10
	li a7, 11
	ecall

	lw a0, b1
	li a7, 1
	ecall
	li a0, 10
	li a7, 11
	ecall

	la a0, s1
	li a7, 4
	ecall
	li a0, 10
	li a7, 11
	ecall

