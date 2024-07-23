let x = "";
let result = [];
let stack = [];
let error;

function shift() {
    if (x.length > 0) {    
        let c = x.charAt(0); 
        if (c >= '0' && c <= '9') {
            if (stack.length === 0) {
                stack.push('N');
                result.push(parseInt(c));
                x = x.substring(1);
            } else if (stack.length > 0 && (stack[stack.length-1] === '+' || stack[stack.length-1] === '*')) {
                stack.push('N');
                result.push(parseInt(c));
                x = x.substring(1);
            } else {
                document.getElementById("log").innerHTML = "Reject.";
                error = true;
            }
        } else if (c === '+') {
            if (stack.length > 0 && stack[stack.length-1] == 'E') {
                stack.push('+');
                x = x.substring(1);
            } else if (stack.length > 0 && (stack[stack.length-1] == '+' || stack[stack.length-1] == '*')) {
                document.getElementById("log").innerHTML = "Reject.";
                error = true;    
            }
        } else if (c == '*') {
            if (stack.length > 0 && stack[stack.length-1] == 'T') {
                stack.push('*');
                x = x.substring(1);
            } else if (stack.length > 0 && (stack[stack.length-1] == '+' || stack[stack.length-1] == '*')) {
                document.getElementById("log").innerHTML = "Reject.";
                error = true;    
            }
        } else {
            document.getElementById("log").innerHTML = "Reject.";
            error = true;
        }        
    }
}

function reduce() {
    if (stack.length >= 3 && stack[stack.length - 3] === 'T' && stack[stack.length - 2] === '*' && stack[stack.length - 1] === 'F') {
        stack.pop();
        stack.pop();
        stack.pop();
        stack.push('T');
        let a = result.pop()
        let b = result.pop()
        result.push(a*b);
    } else if (stack.length >= 3 && stack[stack.length - 3] === 'E' && stack[stack.length - 2] === '+' && stack[stack.length - 1] === 'T') {
        stack.pop();
        stack.pop();
        stack.pop();
        stack.push('E');
        let a = result.pop()
        let b = result.pop()
        result.push(a+b);
    } else if (stack.length > 0 && stack[stack.length -1 ] === 'E' && x.length === 0) {
        stack.pop();
        stack.push('S');
    } else if (stack.length > 0 && stack[stack.length -1 ] === 'T') {
        stack.pop();
        stack.push('E');
    } else if (stack.length > 0 && stack[stack.length -1 ] === 'F') {
        stack.pop();
        stack.push('T');
    } else if (stack.length > 0 && stack[stack.length -1 ] === 'N') {
        stack.pop();
        stack.push('F');
    } else if (stack.length === 1 && stack[0] === 'S' && x.length === 0) {
        stack.pop();
        document.getElementById("log").innerHTML = "Accept: " + result[0];
    } else if (stack.length > 0 && x.length === 0) {
        document.getElementById("log").innerHTML = "Reject.";
        error = true;
    }
    
}

function parse() {
    x = document.getElementById("x").value;
    result = [];
    stack = [];
    error = false;
    let cont = 0;
    while (!error && (x.length > 0 || stack.length > 0)) {
        shift();
        reduce();
        cont ++;
    }
    //document.getElementById("log").innerHTML += cont;
}

document.getElementById("x").addEventListener("keydown", function(e) {
    if (e.key === "Enter") {
        parse();
    }
});
