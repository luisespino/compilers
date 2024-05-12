let x;

function S() {
    document.getElementById("log").innerHTML = x + ' = ' + E();
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

x = "2+3*5";
S();
