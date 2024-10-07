{
    const variables = {};
    let t = 0;
    let data = ".global _start\n\n.data\n";
    let code = "\n.text\n_start:";
}

start
    = statement* {
  	    code += '\n\t'+'li a0, 0';
        code += '\n\t'+'li a7, 93';
        code += '\n\t'+'ecall\n';
  	    return data+code;
    }

statement
    = declaration
    / assignment
    / _ // Permite líneas en blanco

declaration
    = type:"int" _ name:[a-zA-Z]+ _ "=" _ value:intValue __ ";" {
        const varName = name.join("");
        const varValue = value;
        variables[varName] = { name: varName, value: varValue, type: type };
        data += varName+": .word "+varValue+"\n";
    }
    / type:"float" _ name:[a-zA-Z]+ _ "=" _ value:floatValue __ ";" {
        const varName = name.join("");
        const varValue = value;
        variables[varName] = { name: varName, value: varValue, type: type };
        data += varName+": .float "+varValue+"\n";
    }

assignment
    = name:variable _ "=" _ expression:operation __ ";" {
        const result = expression;
        let varName = variables[name].name; 
        if (result == 't0') {
      	    code += '\n\t'+'sw t0, '+varName;
        } else {
            code += '\n\t'+'la t0, '+varName;      
      	    code += '\n\t'+'fsw f0, (t0)';
        }
    }

operation
    = left:variable _ "+" _ right:variable {
        let leftType = variables[left].type;
        let rightType = variables[right].type;
        let leftName = variables[left].name;
        let rightName = variables[right].name;
        let register = 't0';

        if (leftType == "int") {
      	    if (rightType == "int") {
        	    code += '\n\t'+'lw t1, '+leftName;
                code += '\n\t'+'lw t2, '+rightName;
                code += '\n\t'+'add t0, t1, t2';
            } else {
        	    code += '\n\t'+'lw t0, '+leftName;
        	    code += '\n\tfcvt.s.w f1, t0';
                code += '\n\t'+'la t0, '+rightName;
        	    code += '\n\t'+'flw f2, (t0)';
                code += '\n\tfadd.s f0, f1, f2';
                register = 'f0';
            }
        } else {
		    if (rightType == "int") {
        	    // complete!
            } else {
        	    // complete!
            }
        }
        return register;
    }
    / value:(intValue / floatValue) {
      return value;
    }

variable
    = name:[a-zA-Z]+ {
        return name.join("");
    }

intValue
    = [0-9]+ {
        return parseInt(text());
    }

floatValue
    = [0-9]+ "." [0-9]+ {
        return parseFloat(text());
    }

_ "whitespace"
    = [ \t\n\r]+ // Permite al menos cualquier cantidad de espacios en blanco
 
__ "whitespace"
    = [ \t\n\r]* // Permite cualquier cantidad de espacios en blanco
