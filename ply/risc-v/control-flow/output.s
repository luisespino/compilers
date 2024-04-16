.data
x: .word 10
y: .word 10

.text
.globl _start
_start:
	lw t0, x
	li t1, 100
	blt t0, t1, L1
	j L2
L2:
	lw t0, x
	li t1, 200
	bgt t0, t1, L3
	j L4
L3:
	lw t0, x
	lw t1, y
	bne t0, t1, L5
	j L6
L1:
L5:
	la t0, x
	li t1, 0
	sw t1, 0(t0)
L4:
L6:
	lw a0, x
	li a7, 1
	ecall
	li a0, 10
	li a7, 11
	ecall

	li a7, 10
	ecall

