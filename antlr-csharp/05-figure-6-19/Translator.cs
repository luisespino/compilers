using System;
using System.Collections.Generic;
using System.Linq;
using Antlr4.Runtime;
using Antlr4.Runtime.Tree;

public class Attr
{
    public string Addr;
    public string Code;
}

public class Translator : GrammarBaseVisitor<Attr>
{
    public int TID = 0;

    public override Attr VisitProgram(GrammarParser.ProgramContext context)
    {
        string result = "";
        foreach (var sContext in context.s())
        {
            result += Visit(sContext).Code;
        }
        return new Attr { Addr = "", Code = result };
    }

    public override Attr VisitAssign(GrammarParser.AssignContext context)
    {
        string id = context.ID().GetText();
        Attr expr = Visit(context.e());
        string code = expr.Code;
        code += id + " = " + expr.Addr + "\n";
        return new Attr { Addr = "", Code = code };
    }

    public override Attr VisitSumres(GrammarParser.SumresContext context)
    {
        Attr left = Visit(context.e());
        Attr right = Visit(context.t());
        string op = " " + context.op.Text + " ";
        TID++;
        string addr = "t" + TID; 
        string code = left.Code + right.Code;
        code += addr + " = " + left.Addr + op + right.Addr + "\n";
        return new Attr { Addr = addr, Code = code };
    }

    public override Attr VisitMuldiv(GrammarParser.MuldivContext context)
    {
        Attr left = Visit(context.t());
        Attr right = Visit(context.f());
        string op = " " + context.op.Text + " ";
        TID++;
        string addr = "t" + TID; 
        string code = left.Code + right.Code;
        code += addr + " = " + left.Addr + op + right.Addr + "\n";
        return new Attr { Addr = addr, Code = code };
    }

    public override Attr VisitPassE(GrammarParser.PassEContext context)
    {
        return Visit(context.e());
    }

    public override Attr VisitId(GrammarParser.IdContext context)
    {
        string id = context.ID().GetText();
        return new Attr { Addr = id, Code = "" };
    }

    public override Attr VisitNum(GrammarParser.NumContext context)
    {
    	string num = context.NUM().GetText(); 
        return new Attr { Addr = num, Code = "" };
    }
}

