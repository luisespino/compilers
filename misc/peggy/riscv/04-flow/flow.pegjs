{
	// global
	let addr = 0;
	let label = 1;
    let symbol = {};
	let data = '.global _start\n\n.data\n';
	let code = '\n.text\n_start:\n';

	// synthesized attributes 
	class syn {
    	constructor(addr, code, ltrue, lfalse) {
        	this.addr  = addr;   	// int or name
        	this.code  = code;		// code for riscv
        	this.true  = ltrue;  	// true label
        	this.false = lfalse;	// false label
        }
	}

}



prog
	= s:sentence *{
    	let allCode = data+code;
        for (const sentence of s) {
            allCode += sentence.code;
        }
        allCode += '\n\t'+'li a0, 0';
        allCode += '\n\t'+'li a7, 93';
        allCode += '\n\t'+'ecall\n';
        return allCode;
    }

sentence
	= s:decl
    / s:if 
    / s:while
    / s:assg
    / __
    
decl
	= 'int' _ name:name _ '=' _ expr:expr _ {
    	symbol[name.addr] = { name: name.addr, value: expr.addr };
        data += name.addr+': .word '+expr.addr+'\n';
        return new syn('decl', expr.code, '', '');
    }

assg
	= name:name _ '=' _ expr:expr _ {
    	symbol[name.addr] = expr.addr;
        expr.code += '\t'+'la t0, '+ name.addr +'\n'
    	expr.code += '\t'+'li t1, '+ expr.addr +'\n'
    	expr.code += '\t'+'sw t1, 0(t0)\n'
    	return new syn('assg', expr.code, '', '') 
	}

if
	= 'if' _ '(' _ b:bool _ ')' _ s:sentence* _ 'endif' _ {
     	let scode = '';
        for (const sentence of s) {
            scode += sentence.code;
        }   
    	b.code += b.true + scode + b.false;
        return new syn ('if', b.code, '', '');
    }

while
	= 'while' _ '(' _ b:bool _ ')' _ s:sentence* _ 'endwhile' _ {
     	let scode = '';
        for (const sentence of s) {
            scode += sentence.code;
        }   
        let begin = 'L'+label;
       	b.code = begin+':\n' + b.code
    	b.code += b.true + scode
    	b.code += '\t'+'j '+begin+'\n'
    	b.code += b.false
    	return new syn('while', b.code, '', '');
    }
 
expr
    = '(' _ expr _ ')' _
    / name
    / int


bool
    = or_expr

or_expr
    = b1:and_expr _ '||' _ b2:or_expr {
        b1.code += b1.false + b2.code;
        let truel = b1.true + b2.true;
        return new syn('or', b1.code, truel, b2.false);
    }
    / and_expr

and_expr
    = b1:term _ '&&' _ b2:and_expr {
        b1.code += b1.true + b2.code;
        let falsel = b1.false + b2.false;
        return new syn('and', b1.code, b2.true, falsel);
    }
    / term

term
    = '!' _ b:term {
        return new syn('not', b.code, b.false, b.true);
    }
    / comparison

comparison
    = e1:expr _ op:("==" / "!=" / "<" / "<=" / ">" / ">=") _ e2:expr  {
        let truel = 'L' + label++;
        let falsel = 'L' + label++;
        e1.code += e2.code;

        if (e1.addr in symbol) {
            e1.code += '\t' + 'lw t0, ' + e1.addr + '\n';
        } else {
            e1.code += '\t' + 'li t0, ' + e1.addr + '\n';
        }

        if (e2.addr in symbol) {
            e1.code += '\t' + 'lw t1, ' + e2.addr + '\n';
        } else {
            e1.code += '\t' + 'li t1, ' + e2.addr + '\n';
        }

        switch (op) {
            case '==':
                e1.code += '\t' + 'beq t0, t1, ' + truel + '\n';
                break;
            case '!=':
                e1.code += '\t' + 'bne t0, t1, ' + truel + '\n';
                break;
            case '<':
                e1.code += '\t' + 'blt t0, t1, ' + truel + '\n';
                break;
            case '<=':
                e1.code += '\t' + 'sub t2, t1, t0\n';
                e1.code += '\t' + 'blez t2, ' + truel + '\n';
                break;
            case '>':
                e1.code += '\t' + 'bgt t0, t1, ' + truel + '\n';
                break;
            case '>=':
                e1.code += '\t' + 'bge t0, t1, ' + truel + '\n';
                
                break;
        }
        
        e1.code += '\t' + 'j ' + falsel + '\n';
        truel += ':\n';
        falsel += ':\n';
        return new syn('op', e1.code, truel, falsel);
    }


    
name
    = name:[a-zA-Z]+ {
        return new syn(name.join(''), '', '', '');
    }

int
    = [0-9]+ {
    	return new syn(parseInt(text()), '', '', '');
    }
    
_ "whitespace"
    = [ \t\n\r]*
   
__ "whitespace"
    = [ \t\n\r]+