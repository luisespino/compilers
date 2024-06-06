

let input = 
`lw t1, 10
lw t2, 10
add t3, t1, t2
mv t0, t3
`
document.getElementById("x").value = input

function check() {
    let x = "";
    x = document.getElementById("x").value;
    document.getElementById("log").innerHTML = "";
    try {
        document.getElementById("log").innerHTML = JSON.stringify(PEG.parse(x));
     } catch (err) {
        document.getElementById("log").innerHTML = err;
    }
}
