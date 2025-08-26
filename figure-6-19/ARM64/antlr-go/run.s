.global _start

.data
a: .word 1
b: .word 2
t1: .word 0
t2: .word 0
c: .word 0
t3: .word 0
t4: .word 0
d: .word 0

.text
_start:
	ldr X0, =t1
	ldr X1, =b
	ldr X2, [X1]
	mov X3, #2
	mul X2, X2, X3
	str X2, [X0]

	ldr X0, =t2
	ldr X1, =a
	ldr X2, [X1]
	ldr X3, =t1
	ldr X4, [X3]
	add X2, X2, X4
	str X2, [X0]

	ldr X0, =c
	ldr X1, =t2
	ldr X2, [X1]
	str X2, [X0]

	ldr X0, =t3
	ldr X1, =c
	ldr X2, [X1]
	add X2, X2, #1
	str X2, [X0]

	ldr X0, =t4
	ldr X1, =t3
	ldr X2, [X1]
	mov X3, #3
	mul X2, X2, X3
	str X2, [X0]

	ldr X0, =d
	ldr X1, =t4
	ldr X2, [X1]
	str X2, [X0]

	mov X8, #93
	svc #0
