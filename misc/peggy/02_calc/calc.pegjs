s = e

e = left:t "+" right:e { return left + right; }
    / t

t = left:f "*" right:t { return left * right; }
    / f

f = _ num:[0-9]+ _ { return parseInt(num.join(""), 10); }

_ = [ \t\n\r]*
