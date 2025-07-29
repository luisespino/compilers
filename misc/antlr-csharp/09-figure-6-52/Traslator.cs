using System;
using System.Collections.Generic;
using System.Linq;
using Antlr4.Runtime;
using Antlr4.Runtime.Tree;

public class Variable
{
    public string Name;
    public string Value;
}

public class Arg
{
    public string Name;
    public string Reg;
}


public class Translator : GrammarBaseVisitor<string>
{
    public string Data = ".global _start\n\n.data\n";
    public string Text = "\n.text\n_start:\n";
    public string Func = "";
    Dictionary<string, Variable> variables = new Dictionary<string, Variable>();
    Dictionary<string, Arg> args = null;
    
    public override string VisitProg(GrammarParser.ProgContext context)
    {
        foreach (var stmtContext in context.stmt())
            Visit(stmtContext);
            
        Text += "\tmov X8, #93\n\tsvc #0\n";
        
        return Data + Text + Func;
    }

    public override string VisitFun(GrammarParser.FunContext context)
    {
        string id = Visit(context.id());
        Func += id + ":\n";
        
        args = new Dictionary<string, Arg>();
        if (context.args() != null)
            Visit(context.args());        
        
        string ret = Visit(context.ret());
        if (args.ContainsKey(ret))
	    Func += "\t"+"mov X0, "+args[ret].Reg+"\n";
	else
	    Func += "\t"+"mov X0, #"+ret+"\n";
        Func += "\t"+"ret\n";
        
        return null;
    }
    
    public override string VisitArgs(GrammarParser.ArgsContext context)
    {
        int argNum = 0;    
    	string argID = null; 
        foreach (var idContext in context.id())
        {
            argID = Visit(idContext);
            args[argID] = new Arg 
            { 
                Name = argID,
                Reg = "X"+ argNum++
            };
        }
    	return null;
    }

    public override string VisitRet(GrammarParser.RetContext context)
    {
    	return Visit(context.expr());
    }
    
    public override string VisitExpr(GrammarParser.ExprContext context)
    {
        if (context.id() != null)
            return VisitId(context.id());
   	return VisitNum(context.num());
    }
    
    public override string VisitId(GrammarParser.IdContext context)
    {
        return context.IDENTIFIER().GetText();
    }

    public override string VisitNum(GrammarParser.NumContext context)
    {
        return context.NUMBER().GetText();
    }
    
    public override string VisitCall(GrammarParser.CallContext context)
    {
        string id = Visit(context.id());
        
        if (context.argsend() != null)
            Visit(context.argsend());
        
        Text += "\t"+"bl "+id+"\n\n";
                
        return null;
    }
    
    public override string VisitArgsend(GrammarParser.ArgsendContext context)
    {
        int argNum = 0;
    	string num = null;
        foreach (var numContext in context.num())
        {
            num = Visit(numContext);
            Text += "\t"+"mov X"+(argNum++)+", #"+num+"\n";    
        }
    	return null;        
    }

}


