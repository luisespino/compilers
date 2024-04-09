.data
hex1: .word 0x3F800000		# 1.0
hex2: .word 0x3FC00000   	# 1.5
f_exp_coef: .word 0xFF
f_man_coef: .word 0x7FFFFF
f_man_norm: .word 0x800000
f_exp_hex1: .word 0x0
f_exp_hex2: .word 0x0
f_exp_diff: .word 0x0
f_man_hex1: .word 0x0
f_man_hex2: .word 0x0
f_man_sum: .word 0x0
f_sign_hex1: .word 0x0
f_sign_hex2: .word 0x0

.text
.globl _start
_start:
	    # load hex values
	    lw a0, hex1
    	lw a1, hex2
    	# call function
    	jal ra, addf
    	# print return value a0
    	li a7, 34
	    ecall
	    # exit
    	li a7, 10
    	ecall
    
addf:
	    # extract signs
    	la t0, f_sign_hex1
    	slli t1, a0, 31
    	andi t1, t1, 0x1
    	sw t1, 0(t0)
    	la t0, f_sign_hex2
    	slli t1, a1, 31
    	andi t1, t1, 0x1
    	sw t1, 0(t0)
    
    	# extract exponents
    	la t0, f_exp_hex1
    	lw t1, f_exp_coef
    	srli t2, a0, 23
    	and t2, t2, t1
    	addi t2, t2, -127
    	sw t2, 0(t0)
    	la t0, f_exp_hex2
    	lw t1, f_exp_coef
    	srli t2, a1, 23
    	and t2, t2, t1
    	addi t2, t2, -127
    	sw t2, 0(t0)
 
    	# extract mantissas
    	la t0, f_man_hex1
    	lw t1, f_man_coef
    	lw t2, f_man_norm
    	and t3, a0, t1
    	or t3, t3, t2
    	sw t3, 0(t0)
    	la t0, f_man_hex2
    	lw t1, f_man_coef
    	lw t2, f_man_norm
    	and t3, a1, t1
    	or t3, t3, t2
    	sw t3, 0(t0)
    
    	# adjust exponents
    	la t0, f_exp_diff
    	lw t1, f_exp_hex1
    	lw t2, f_exp_hex2
    	sub t2, t1, t2
    	sw t2, 0(t0)
    
    	lw t0, f_exp_diff
    	bgtz t0, f_exp_diff_gtz    

f_exp_diff_letz:
    	neg t0, t0
    	lw t1, f_man_hex1
    	la t2, f_man_hex1
    	srl t3, t1, t0
    	sw t3, 0(t2)
    	la t0, f_exp_hex1
    	lw t1, f_exp_hex2
    	sw t1, 0(t0)
    	j f_exp_diff_gtz_end

f_exp_diff_gtz:
    	lw t0, f_exp_diff
    	lw t1, f_man_hex2
    	la t2, f_man_hex2
    	srl t3, t1, t0
    	sw t3, 0(t2)
    	la t0, f_exp_hex2
    	lw t1, f_exp_hex1
    	sw t1, 0(t0)
    
f_exp_diff_gtz_end:
	    # adding mantissas
    	la t0, f_man_sum
    	lw t1, f_man_hex1
    	lw t2, f_man_hex2
    	add t2, t1, t2
    	sw t2, 0(t0)
    
    	# check carry
    	lw t0, f_man_sum
    	li t1, 0x1000000
    	and t2, t0, t1
    	bnez t2, sum_carry
    	j combine_result
    
sum_carry:
    	la t0, f_man_sum
    	lw t1, f_man_sum
    	srli t1, t1, 1
    	sw t1, 0(t0)
    	la t0, f_exp_hex1
    	lw t1, f_exp_hex1
    	addi t1, t1, 1
    	sw t1, 0(t0)
    
combine_result:
	    lw t0, f_sign_hex1
    	slli t0, t0, 31
    	lw t1, f_exp_hex1
    	addi t1, t1, 127
    	slli t1, t1, 23
    	or t0, t0, t1
    	lw t2, f_man_sum
    	lw t3, f_man_coef
    	and t2, t2, t3
    	or t0, t0, t2
    	mv a0, t0
    	ret
