.data
a: .word 0
a_cols: .word 0
b: .word 1, 2, 3
b_cols: .word 3
c: .word 1, 2, 3, 4, 5, 6, 7, 8, 9
c_rows: .word 3
c_cols: .word 3
d: .word 1, 2, 3, 4, 5, 6, 7, 8
d_face: .word 2
d_rows: .word 2
d_cols: .word 2

.text
.globl _start
_start:
	li t0, 1
	slli t0, t0, 2
	la t1, b
	add t2, t1, t0
	lw a0, 0(t2)
	li a7, 1
	ecall
	li a0, 10
	li a7, 11
	ecall

	li t0, 2
	slli t0, t0, 2
	li t1, 2
	slli t1, t1, 2
	lw t2, c_cols
	mul t3, t0, t2
	add t3, t3, t1
	la t4, c
	add t4, t4, t3
	lw a0, 0(t4)
	li a7, 1
	ecall
	li a0, 10
	li a7, 11
	ecall

	li t0, 1
	slli t0, t0, 2
	li t1, 0
	slli t1, t1, 2
	li t2, 0
	slli t2, t2, 2
	lw t3, d_rows
	mul t4, t3, t0
	add t4, t4, t1
	lw t5, d_cols
	mul t6, t5, t4
	add t6, t6, t2
	la t0, d
	add t0, t0, t6
	lw a0, 0(t0)
	li a7, 1
	ecall
	li a0, 10
	li a7, 11
	ecall

