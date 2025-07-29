using System;
using System.Collections.Generic;
using System.Linq;
using Antlr4.Runtime;
using Antlr4.Runtime.Tree;

public class Attr
{
    public string Addr;
    public string Code;
    public string True;
    public string False;
}

public class Variable
{
    public string Name;
    public string Type;
    public string Value;
}

public class Translator : GrammarBaseVisitor<Attr>
{
    public string Data = ".global _start\n\n.data\n";
    public string Text = "\n.text\n_start:\n";
    Dictionary<string, Variable> variables = new Dictionary<string, Variable>();
    int Label = 1;
    
    public override Attr VisitProg(GrammarParser.ProgContext context)
    {
        foreach (var stmtContext in context.stmt())
        {
            Text += Visit(stmtContext)?.Code ?? "";
        }
        Text += "\tmov X8, #93\n\tsvc #0\n";
        return new Attr { Addr = null, Code = Data + Text, True = null, False = null};
    }

    public override Attr VisitDecl(GrammarParser.DeclContext context)
    {
        string variable = Visit(context.var()).Addr;
        string type = context.type.Text;
        Attr expr = Visit(context.expr());

        variables[variable] = new Variable 
        {
            Name = variable, 
            Type = type,
            Value = expr.Addr // assuming it is a DIGIT
        };

        Data += variable + ": .word " + expr.Addr + "\n";
        
        return new Attr { Addr = null, Code = expr.Code, True = null, False = null };
    }
    
    public override Attr VisitAssg(GrammarParser.AssgContext context)
    {
        string variable = Visit(context.var()).Addr;
        Attr expr = Visit(context.expr());

        variables[variable] = new Variable 
        {
            Name = variable,
            Type = "int",
            Value = expr.Addr // assuming it is a DIGIT
        };

        expr.Code += "\t"+"ldr X0, ="+variable+"\n";
        expr.Code += "\t"+"mov X1, #"+expr.Addr+"\n";
        expr.Code += "\t"+"str X1, [X0]\n\n";
        
        return new Attr { Addr = null, Code = expr.Code, True = null, False = null };
    }

    public override Attr VisitIf(GrammarParser.IfContext context)
    {
    	Attr cond = Visit(context.cond());
    	//?? new Attr { Addr = null, Code = null, True = null, False = null};
    	string scode = "";
        foreach (var stmtContext in context.stmt())
        {
            scode += Visit(stmtContext).Code;
            //scode += Visit(stmtContext)?.Code ?? "";
        }
        //string code = (cond.Code ?? "") + (cond.True ?? "") + scode + (cond.False ?? "");
        cond.Code += cond.True + scode + cond.False;
        return new Attr { Addr = null, Code = cond.Code, True = null, False = null };
    }

/*
    public override Attr VisitWhile(GrammarParser.WhileContext context)
    {
    	return new Attr { Addr = null, Code = null, True = null, False = null};
    }
*/
    public override Attr VisitCond(GrammarParser.CondContext context)
    {
        return Visit(context.@bool());
    }
    
    public override Attr VisitOrExpr(GrammarParser.OrExprContext context)
    {
        Attr b1 = Visit(context.and());
        Attr b2 = Visit(context.or());
        
        b1.Code += b1.False + b2.Code;
        string t = b1.True + b2.True;
        
        return new Attr { Addr = null, Code = b1.Code, True = t, False = b2.False };
    }
    
    public override Attr VisitAndExpr(GrammarParser.AndExprContext context)
    {
        Attr b1 = Visit(context.not());
        Attr b2 = Visit(context.and());
        
        b1.Code += b1.True + b2.Code;
        string f = b1.False + b2.False;
        
        return new Attr { Addr = null, Code = b1.Code, True = b2.True, False = f };
    }

    public override Attr VisitRel(GrammarParser.RelContext context)
    {
    	Attr e0 = Visit(context.expr(0));
    	Attr e1 = Visit(context.expr(1));
    	
    	int op = context.oprel().Start.Type;
    	
    	string t = "L" + Label++;
    	string f = "L" + Label++;
    	e0.Code += e1.Code;
    	
    	if (variables.ContainsKey(e0.Addr))
	    e0.Code += "\t"+"ldr X0, ="+e0.Addr+"\n";
	else
	    e0.Code += "\t"+"mov X0, #"+e0.Addr+"\n";
 
     	if (variables.ContainsKey(e1.Addr))
	    e0.Code += "\t"+"ldr X1, ="+e1.Addr+"\n";
	else
	    e0.Code += "\t"+"mov X1, #"+e1.Addr+"\n";
	
    	switch (op)
    	{
            case GrammarParser.EQ: // '=='
                e0.Code += "\tcmp X0, X1\n";
                e0.Code += $"\tb.eq {t}\n";
                break;

            case GrammarParser.NEQ: // '!='
                e0.Code += "\tcmp X0, X1\n";
                e0.Code += $"\tb.ne {t}\n";
                break;

            case GrammarParser.LT: // '<'
                e0.Code += "\tcmp X0, X1\n";
                e0.Code += $"\tb.lt {t}\n";
                break;

            case GrammarParser.LE: // '<='
                e0.Code += "\tcmp X0, X1\n";
                e0.Code += $"\tb.le {t}\n";
                break;

            case GrammarParser.GT: // '>'
                e0.Code += "\tcmp X0, X1\n";
                e0.Code += $"\tb.gt {t}\n";
                break;

            case GrammarParser.GE: // '>='
                e0.Code += "\tcmp X0, X1\n";
                e0.Code += $"\tb.ge {t}\n";
                break; 
        }
        
        e0.Code += $"\tj {f}\n";
        t += "\n";
        f += "\n";
        
        return new Attr { Addr = null, Code = e0.Code, True = t, False = f };
    }
    
    public override Attr VisitNotExpr(GrammarParser.NotExprContext context)
    {
        Attr b = Visit(context.not());
        
        return new Attr { Addr = null, Code = b.Code, True = b.False, False = b.True };
    }
    
    public override Attr VisitVar(GrammarParser.VarContext context)
    {
        string name = context.NAME().GetText();
        return new Attr { Addr = name, Code = null, True = null, False = null };
    }

    public override Attr VisitNum(GrammarParser.NumContext context)
    {
        string digit = context.DIGIT().GetText();
        return new Attr { Addr = digit, Code = null, True = null, False = null };
    }
}


