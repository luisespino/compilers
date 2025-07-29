{
	let t = 0;
	let code = ".global _start\n";
    code += "\n_start:\n";    
}

Sentence
  = Expression
	{
    	code += '\n\t'+'li a0, 0';
    	code += '\n\t'+'li a7, 93';
   		code += '\n\t'+'ecall\n';
       	return code;
    }
    
Expression
  = head:Term tail:(_ ("+" / "-") _ Term)* {
      return tail.reduce(function(result, element) {
        if (element[1] === "+")
		{
    		t += 4;
    		code += '\n\t'+'li t3, '+result;
        	code += '\n\t'+'lw t1, 0(t3)';
    		code += '\n\t'+'li t3, '+element[3];
    		code += '\n\t'+'lw t2, 0(t3)';
    		code += '\n\t'+'add t0, t1, t2';
    		code += '\n\t'+'li t3, '+t;
    		code += '\n\t'+'sw t0, 0(t3)';        
        	return t;
    	}
      }, head);
    }

Term
  = head:Factor tail:(_ ("*" / "/") _ Factor)* {
      return tail.reduce(function(result, element) {
        if (element[1] === "*")
		{
    		t += 4;
    		code += '\n\t'+'li t3, '+result;
        	code += '\n\t'+'lw t1, 0(t3)';
    		code += '\n\t'+'li t3, '+element[3];
    		code += '\n\t'+'lw t2, 0(t3)';
    		code += '\n\t'+'mul t0, t1, t2';
    		code += '\n\t'+'li t3, '+t;
    		code += '\n\t'+'sw t0, 0(t3)';        
        	return t;
    	}
      }, head);
    }

Factor
  = "(" _ expr:Expression _ ")" { return expr; }
  / Integer

Integer "integer"
  = _ num:[0-9]+ _ 
    {
    	t += 4;
    	code += '\n\t'+'li t0, '+num.join("");
    	code += '\n\t'+'li t3, '+t;
        code += '\n\t'+'sw t0, 0(t3)';      
        return t;    
    }
  
_ "whitespace"
  = [ \t\n\r]*