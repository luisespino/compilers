{
	let t = 0;
	let code = ".global _start\n";
    code += "\n_start:\n";    
}

s = res:e
	{
    	code += '\n\t'+'li a0, 0';
    	code += '\n\t'+'li a7, 93';
   		code += '\n\t'+'ecall\n';
       	return code;
    }

e = left:t "+" right:e 
	{
    	t += 4;
    	code += '\n\t'+'li t3, '+left;
        code += '\n\t'+'lw t1, 0(t3)';
    	code += '\n\t'+'li t3, '+right;
    	code += '\n\t'+'lw t2, 0(t3)';
    	code += '\n\t'+'add t0, t1, t2';
    	code += '\n\t'+'li t3, '+t;
    	code += '\n\t'+'sw t0, 0(t3)';        
        return t;
    }
    / t

t = left:f "*" right:t 
	{
    	t += 4;
    	code += '\n\t'+'li t3, '+left;
        code += '\n\t'+'lw t1, 0(t3)';
    	code += '\n\t'+'li t3, '+right;
    	code += '\n\t'+'lw t2, 0(t3)';
    	code += '\n\t'+'mul t0, t1, t2';
    	code += '\n\t'+'li t3, '+t;
    	code += '\n\t'+'sw t0, 0(t3)';        
        return t;
    }
    / f

f = _ "(" mid:e ")" { return mid; }
	/ _ num:[0-9]+ _ 
    {
    	t += 4;
    	code += '\n\t'+'li t0, '+num;
    	code += '\n\t'+'li t3, '+t;
        code += '\n\t'+'sw t0, 0(t3)';      
        return t;    
    }

_ = [ \t\n\r]*