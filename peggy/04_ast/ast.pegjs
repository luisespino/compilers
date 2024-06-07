{
    class node{
  		constructor( value, left=null, right=null ) {
      	    this.value = value;
     		this.left = left;
      		this.right = right;
    	}
    }
    
    function generateDot(root) {
        let dot = "graph G {\n";
        let counter = { count: 0 };

        function traverse(node) {
            let current = counter.count++;
            dot += `  ${current} [label="${node.value}"];\n`;
            if (node.left) {
                let left = traverse(node.left);
                dot += `  ${current} -- ${left};\n`;
            }
            if (node.right) {
                let right = traverse(node.right);
                dot += `  ${current} -- ${right};\n`;
            }
            return current;
        }
        traverse(root);
        dot += "}\n";
        return dot;
    }
}

s = root:e { return generateDot(root);}

e = left:t "+" right:e { return new node("+",left,right); }
    / t

t = left:f "*" right:t { return new node("*",left,right); }
    / f

f = _ num:[0-9]+ _ { return new node(num); }

_ = [ \t\n\r]*