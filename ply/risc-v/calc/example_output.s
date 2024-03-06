.globl _start
_start:

	li t0, 1
	li t3, 4
	sw t0, 0(t3)

	li t0, 2
	li t3, 8
	sw t0, 0(t3)

	li t3, 4
	lw t1, 0(t3)
	li t3, 8
	lw t2, 0(t3)
	add t0, t1, t2
	li t3, 12
	sw t0, 0(t3)

	li t0, 2
	li t3, 16
	sw t0, 0(t3)

	li t0, 3
	li t3, 20
	sw t0, 0(t3)

	li t3, 16
	lw t1, 0(t3)
	li t3, 20
	lw t2, 0(t3)
	add t0, t1, t2
	li t3, 24
	sw t0, 0(t3)

	li t3, 12
	lw t1, 0(t3)
	li t3, 24
	lw t2, 0(t3)
	mul t0, t1, t2
	li t3, 28
	sw t0, 0(t3)

	li t3, 28
	lw a0, 0(t3)
	li a7, 1
	ecall

	li a0, 0
	li a7, 93
	ecall
