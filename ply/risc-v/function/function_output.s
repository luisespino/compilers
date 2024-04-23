.data
a: .word 1
b: .word 2

.text
.globl _start
_start:
	li a0, 1
	li a1, 2
	jal ra, sum
	li a7, 1
	ecall
	li a0, 10
	li a7, 11
	ecall

	li a7, 10
	ecall
sum:
	add t0, a0, a1
	mv a0, t0
	ret
