// import model 
import { StartRules, SyntaxError, parse } from './pvm/parser.mjs'

let x;

document.getElementById("x").addEventListener("keydown", function(e) {
    if (e.key === "Enter") {
        x = document.getElementById("x").value;
        document.getElementById("log").innerHTML = parse(x);
    }
});
