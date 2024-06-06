// Luis Espino 
// 2024

program = instructions

instructions
	= instruction *
    
instruction
	// arithmetic 
	= _ op:"add "i rd:reg comma rs1:reg comma rs2:reg end
 	/ _ op:"sub "i rd:reg comma rs1:reg comma rs2:reg end
	/ _ op:"addi "i rd:reg comma rs1:reg comma imm:imm end

	// logic
	/ _ op:"and "i rd:reg comma rs1:reg comma rs2:reg end
	/ _ op:"or "i rd:reg comma rs1:reg comma rs2:reg end
	/ _ op:"xor "i rd:reg comma rs1:reg comma rs2:reg end
	
    	// load/store
	/ _ op:"lw "i rd:reg comma imm:imm end
   	/ _ op:"st "i rd:reg comma imm:imm end

	// pseudo
	/ _ op:"mv "i rd:reg comma rs:reg end
    	/ _ op:"nop"i _ end
    	/ end
     
reg "register"
	= _ ("zero"i/"x0"i/"r0"i)
    	/ _ ("ra"i/"x1"i/"r1"i)
    	/ _ ("sp"i/"x2"i/"r2"i)
    	/ _ ("gp"i/"x3"i/"r3"i)
    	/ _ ("tp"i/"x4"i/"r4"i)
    	/ _ ("t0"i/"x5"i/"r5"i)
    	/ _ ("t1"i/"x6"i/"r6"i)
    	/ _ ("t2"i/"x7"i/"r7"i)
    	/ _ ("s0"i/"x8"i/"r8"i/"fp"i)
        / _ ("s1"i/"x9"i/"r9"i)
        / _ ("a0"i/"x10"i/"r10"i)
    	/ _ ("a1"i/"x11"i/"r11"i)
    	/ _ ("a2"i/"x12"i/"r12"i)
    	/ _ ("a3"i/"x13"i/"r13"i)
    	/ _ ("a4"i/"x14"i/"r14"i)
    	/ _ ("a5"i/"x15"i/"r15"i)
    	/ _ ("a6"i/"x16"i/"r16"i)
    	/ _ ("a7"i/"x17"i/"r17"i)
    	/ _ ("s2"i/"x18"i/"r18"i)
    	/ _ ("s3"i/"x19"i/"r19"i)
    	/ _ ("s4"i/"x20"i/"r20"i)
    	/ _ ("s5"i/"x21"i/"r21"i)
    	/ _ ("s6"i/"x22"i/"r22"i)
    	/ _ ("s7"i/"x23"i/"r23"i)
    	/ _ ("s8"i/"x24"i/"r24"i)
    	/ _ ("s9"i/"x25"i/"r25"i)
    	/ _ ("s10"i/"x26"i/"r26"i)
    	/ _ ("s11"i/"x27"i/"r27"i)
    	/ _ ("t3"i/"x28"i/"r28"i)
    	/ _ ("t4"i/"x29"i/"r29"i)
    	/ _ ("t5"i/"x30"i/"r30"i)
    	/ _ ("t6"i/"x31"i/"r31"i)

imm "integer"
	= _ i:[0-9]+

comma "comma"
	= _ "," 

end "newline"
	= _ "\n"
 
_ "ignored"
	= [ \t]*
