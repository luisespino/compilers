{
	let variables = {};
	let data = ".global _start\n\n.data\n";
    let code = "\n.text\n_start:";
    
    function getDimensions(lst) {
    	if (Array.isArray(lst)) {
        	if (lst.length > 0) {
            	return [lst.length].concat(getDimensions(lst[0]));
        	} else {
            	return [0];
        	}
    	} else {
        	return [];
      	}
	}
    
    function flattenArray(arr) {
    	const result = [];

    	function flattenHelper(arr) {
        	for (const item of arr) {
            	if (Array.isArray(item)) {
                	flattenHelper(item);
            	} else {
                	result.push(item);
            	}
        	}
    	}

    	flattenHelper(arr);
    	return result;
	}
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
    / __ // Permite líneas en blanco
    
declaration
	= type:"int" _ variable:variable _ "=" _ expression:expression _ ";" {
    	variables[variable] = { name: variable, value: expression, type: type };
		let dimensions = getDimensions(expression);
		let value = flattenArray(expression);
		data += variable + ': .word ';
		let nums = value.join(', ');
    	if (nums == '') {
        	data += '0\n'
    	} else {
        	data += nums + '\n'
        }
    	if (dimensions.length == 1) {
        	data += variable + '_cols: .word '+String(dimensions[0])+'\n'
        } else if(dimensions.length == 2) {
        	data += variable + '_rows: .word '+String(dimensions[0])+'\n'
        	data += variable + '_cols: .word '+String(dimensions[1])+'\n'
        } else if(dimensions.length == 3) {
        	data += variable + '_face: .word '+String(dimensions[0])+'\n'
        	data += variable + '_rows: .word '+String(dimensions[1])+'\n'
        	data += variable + '_cols: .word '+String(dimensions[2])+'\n'            
        }
    }
    / type:"int" _ variable:variable _ "=" _ value:value _ ";" {
    	variables[variable] = { name: variable, value: value, type: type };
		data += variable+": .word "+value+"\n";
    }

variable
    = name:[a-zA-Z]+ {
        return name.join("");
    }

value
    = [0-9]+ {
        return parseInt(text());
    }
    / expression:expression {
    	return expression;
    }
    
expression
	= "[" _ "]" {
    	return [];
    }
    / "[" _ list:list _ "]" {
    	return list;
    }


list
    = value:value _ "," _ list:list {
    	return [value].concat(list);
    }
	/ value:value {
    	return [value];
    }
    
assignment
    = varDest:variable _ "=" _ varSrc:variable _ index:index _ ";" {
		let varDestName = variables[varDest].name;
        let varSrcName = variables[varSrc].name;
        let varSrcValue = variables[varSrc].value;
        let dimensions = getDimensions(varSrcValue);
		for (let i of index) {
    		varSrcValue = varSrcValue[i];
		}
		if (dimensions.length === 1) { // [COL]
    		code += '\tli t0, ' + index + '\n';
            code += '\tli t1, 4\n';
    		code += '\tmul t0, t0, t1\n';
    		code += '\tla t1, ' + varSrcName + '\n';
    		code += '\tadd t2, t1, t0\n';
            code += '\tla t0, ' + varDestName + '\n';
    		code += '\tlw t1, 0(t2)\n';
    		code += '\tsw t1, 0(t0)\n\n';
		} else if (dimensions.length === 2) { // [ROW][COL]
    		code += '\tli t0, ' + String(index[0]) + '\n';
            code += '\tli t1, 4\n';
    		code += '\tmul t0, t0, t1\n';
    		code += '\tli t1, ' + String(index[1]) + '\n';
            code += '\tli t2, 4\n';
    		code += '\tmul t1, t1, t2\n';
    		code += '\tlw t2, ' + varSrcName + '_cols\n';
    		code += '\tmul t3, t0, t2\n';
    		code += '\tadd t3, t3, t1\n';
    		code += '\tla t4, ' + varSrcName + '\n';
    		code += '\tadd t4, t4, t3\n';
            code += '\tla t0, ' + varDestName + '\n';
            code += '\tlw t1, 0(t4)\n';
    		code += '\tsw t1, 0(t0)\n\n';
		} else if (dimensions.length === 3) { // [FACE][ROW][COL]
    		code += '\tli t0, ' + String(index[0]) + '\n';
            code += '\tli t1, 4\n';
    		code += '\tmul t0, t0, t1\n';
    		code += '\tli t1, ' + String(index[1]) + '\n';
            code += '\tli t2, 4\n';
    		code += '\tmul t1, t1, t2\n';
    		code += '\tli t2, ' + String(index[2]) + '\n';
            code += '\tli t3, 4\n';
    		code += '\tmul t2, t2, t3\n';
    		code += '\tlw t3, ' + varSrcName + '_rows\n';
    		code += '\tmul t4, t3, t0\n';
    		code += '\tadd t4, t4, t1\n';
    		code += '\tlw t5, ' + varSrcName + '_cols\n';
    		code += '\tmul t6, t5, t4\n';
    		code += '\tadd t6, t6, t2\n';            
    		code += '\tla t0, ' + varSrcName + '\n';
    		code += '\tadd t0, t0, t6\n';
            code += '\tla t1, ' + varDestName + '\n';
            code += '\tlw t2, 0(t0)\n';
    		code += '\tsw t2, 0(t1)\n\n';
		}
    }
    
index
	= "[" _ value:value _ "]" _ index:index {
    	return [value].concat(index);
    }
	/ "[" _ value:value _ "]" {
    	return [value];
    }

__ "whitespace"
    = [ \t\n\r]+ // Permite al menos cualquier cantidad de espacios en blanco
 
_ "whitespace"
    = [ \t\n\r]* // Permite cualquier cantidad de espacios en blanco
