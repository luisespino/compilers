.global _start

.data
x: .dword 0
y: .dword 0
z: .dword 0
a: .dword 0
a_cols: .dword 0
b: .dword 1, 2, 3
b_cols: .dword 3
c: .dword 1, 2, 3, 4, 5, 6, 7, 8, 9
c_rows: .dword 3
c_cols: .dword 3
d: .dword 1, 2, 3, 4, 5, 6, 7, 8
d_face: .dword 2
d_rows: .dword 2
d_cols: .dword 2

.text
_start:
	mov x0, #1
	mov x1, #8
	mul x0, x0, x1
	ldr x1, =b
	add x2, x1, x0
	ldr x0, =x
	ldr x1, [x2]
	str x1, [x0]

	mov x0, #2
	mov x1, #8
	mul x0, x0, x1
	mov x1, #2
	mov x2, #8
	mul x1, x1, x2
	ldr x2, =c_cols
	ldr x2, [x2]
	mul x3, x0, x2
	add x3, x3, x1
	ldr x4, =c
	add x4, x4, x3
	ldr x0, =y
	ldr x1, [x4]
	str x1, [x0]

	mov x0, #1
	mov x1, #8
	mul x0, x0, x1
	mov x1, #0
	mov x2, #8
	mul x1, x1, x2
	mov x2, #0
	mov x3, #8
	mul x2, x2, x3
	ldr x3, =d_rows
	ldr x3, [x3]
	mul x4, x3, x0
	add x4, x4, x1
	ldr x5, =d_cols
	ldr x5, [x5]
	mul x6, x5, x4
	add x6, x6, x2
	ldr x0, =d
	add x0, x0, x6
	ldr x1, =z
	ldr x2, [x0]
	str x2, [x1]

	mov x8, #93
	svc #0
