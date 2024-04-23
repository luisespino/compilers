.text
.globl _start

_start:
	li a0, 5		# load param n = 5
	jal ra, fact		# call factorial
	li a7, 1		# load int flag
	ecall			# print a0
	li a7, 10		# load exit flag
	ecall			# exit

fact:
	bnez a0, nozero		# jump if n != 0
	li a0, 1		# n = 0 then fact = 1
	ret

nozero:
	addi sp, sp, -8		# reserve 2 bytes from stack
	sw ra, 4(sp)		# store ra in stack
	sw s0, 0(sp)		# store n in stack
	mv s0, a0		# store n in s0
	addi a0, a0, -1		# decreases n
	jal ra, factorial	# call factorial

	mul a0, a0, s0          # multiply n-1 by n
	lw s0, 0(sp)            # load n from stack
	lw ra, 4(sp)            # load ra frin stack
	addi sp, sp, 8          # free 2 bytes from stack
	ret                     
