let x;

function S() {
    document.getElementById("log").innerHTML = E();
}

function E() {
    let e = T();
    while (x.charAt(0) === '+') {
        x = x.substring(1);
        let t = T();
        e = e + t;
    }
    return e;
}

function T() {
    let t = F();
    while (x.charAt(0) === '*') {
        x = x.substring(1);
        let f = F();
        t = t * f;
    }
    return t;
}

function F() {
    if (x.charAt(0) >= '0' && x.charAt(0) <= '9') {
        let num = parseInt(x.charAt(0));
        x = x.substring(1);
        return num;
    }
}

document.getElementById("x").addEventListener("keydown", function(e) {
    if (e.key === "Enter") {
        x = document.getElementById("x").value;
        S();
    }
});
