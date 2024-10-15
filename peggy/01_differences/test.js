function parse(param) {
    return s({in:param, out:null}).out;
}

function s(param) {
    return e(param);
}

function e(param) {
    let t1 = t(param);
    if (t1.in.charAt(0) === '+'){
        t1.in = t1.in.substring(1);
        let e1 = e(t1);
        return {in:e1.in, out: t1.out + e1.out};
    }
    return t(param);
}

function t(param) {
    let f1 = f(param);
    if (f1.in.charAt(0) === '*'){
        f1.in = f1.in.substring(1);
        let t1 = t(f1);
        return {in:t1.in, out: f1.out * t1.out};
    }
    return f(param);
}

function f(param) {
    let f1 = {in:param.in, out:param.out};
    if (f1.in.charAt(0) >= '0' && f1.in.charAt(0) <= '9') {
        let num = parseInt(f1.in.charAt(0));
        f1.in = f1.in.substring(1);
        return {in:f1.in, out:num};
    }
}

document.getElementById("x").addEventListener("keydown", function(e) {
    if (e.key === "Enter") {
        x = document.getElementById("x").value;
        document.getElementById("log").innerHTML = parse(x);
    }
});
