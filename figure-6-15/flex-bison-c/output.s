.global _start

.data
a: .word 0
b: .word 0, 0, 0, 0
f1: .double 0.0
f2: .double 0.0, 0.0, 0.0, 0.0

.text
_start:
	mov x8, #93
	svc #0
