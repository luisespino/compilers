using System;
using System.Collections.Generic;
using System.Linq;
using Antlr4.Runtime;
using Antlr4.Runtime.Tree;

public class Attr
{
    public string AddrNUM;
    public string AddrID;
    public string Code;
}

public class SymbolTable
{
    private HashSet<string> IDS;
    
    public SymbolTable()
    {
        IDS = new HashSet<string>();
    }
    
    public void Insert(string name)
    {
        IDS.Add(name);
    }
    
    public bool Exists(string name)
    {
        return IDS.Contains(name);
    }
}

public class Translator : GrammarBaseVisitor<Attr>
{
    public int TID = 0;
    public string Data = ".global _start\n\n.data\n";
    public string Code = "\n.text\n_start:\n";
    
    SymbolTable symbolTable = new SymbolTable();

    public bool isNUM(Attr attr)
    {
        if (attr.AddrNUM != null && attr.AddrID == null) return true;
        else return false;
    }
    
    public override Attr VisitProgram(GrammarParser.ProgramContext context)
    {
        foreach (var sContext in context.s())
        {
            Visit(sContext);
        }
        Code += "\n\tmov X8, #93\n\tsvc #0\n";
        return new Attr { AddrNUM = null, AddrID = null, Code = Data+Code };
    }

    public override Attr VisitAssign(GrammarParser.AssignContext context)
    {
        string id = context.ID().GetText();
        Attr expr = Visit(context.e());
        
        
        if (isNUM(expr))
        {
            if (symbolTable.Exists(id))
            {
                Code += "\t"+"ldr X0, ="+id+"\n";
                Code += "\t"+"mov X1, #"+expr.AddrNUM+"\n";
                Code += "\t"+"str X1, [X0]\n\n";
            }
            else
            {
                Data += id+": .word "+expr.AddrNUM+"\n";
            }            
        }
        else
        {
            if (symbolTable.Exists(id))
            {
                Code += "\t"+"ldr X0, ="+id+"\n";
                Code += "\t"+"ldr X1, ="+expr.AddrID+"\n";
                Code += "\t"+"ldr X2, [X1]\n";                
                Code += "\t"+"str X2, [X0]\n\n";
            }
            else
            {
                Data += id+": .word 0\n";
                
                Code += "\t"+"ldr X0, ="+id+"\n";
                Code += "\t"+"ldr X1, ="+expr.AddrID+"\n";
                Code += "\t"+"ldr X2, [X1]\n";                
                Code += "\t"+"str X2, [X0]\n\n";                
            }
        }
        
        return new Attr { AddrNUM = null, AddrID = null };
    }

    public override Attr VisitSumres(GrammarParser.SumresContext context)
    {
        Attr left = Visit(context.e());
        Attr right = Visit(context.t());
        
        TID++;
        string addr = "t" + TID;
        Data += addr+": .word 0\n";
        
        Code += "\t"+"ldr X0, ="+addr+"\n";
	        
        string op = (context.op.Text == "+")?"add":"sub";
        if (isNUM(left)) {
            Code += "\t"+"ldr X1, [X0]\n";
            if (isNUM(right)) {
                Code += "\t"+"add X1, X1, #"+left.AddrNUM+"\n";
                Code += "\t"+op+" X1, X1, #"+right.AddrNUM+"\n";
                Code += "\t"+"str X1, [X0]\n";
            }
            else
            {
                Code += "\t"+"add X1, X1, #"+left.AddrNUM+"\n";
                Code += "\t"+"ldr X2, ="+right.AddrID+"\n";
                Code += "\t"+"ldr X3, [X2]\n";
                Code += "\t"+op+" X1, X3, #"+right.AddrNUM+"\n";
                Code += "\t"+"str X1, [X0]\n";
            }
        }
        else
        {
            if (isNUM(right)) {
                Code += "\t"+"ldr X2, ="+left.AddrID+"\n";
                Code += "\t"+"ldr X3, [X2]\n";
                Code += "\t"+op+" X3, X3, #"+right.AddrNUM+"\n";
                Code += "\t"+"str X3, [X0]\n";
            }
            else
            {
                Code += "\t"+"ldr X2, ="+left.AddrID+"\n";
                Code += "\t"+"ldr X3, [X2]\n";
                Code += "\t"+"ldr X4, ="+right.AddrID+"\n";
                Code += "\t"+"ldr X5, [X4]\n";
                Code += "\t"+op+" X3, X3, X5\n";
                Code += "\t"+"str X3, [X0]\n";
            }        
        }
        Code += "\n";
        return new Attr { AddrNUM = null, AddrID = addr };
    }

    public override Attr VisitMuldiv(GrammarParser.MuldivContext context)
    {
        Attr left = Visit(context.t());
        Attr right = Visit(context.f());
        
        TID++;
        string addr = "t" + TID;
        Data += addr+": .word 0\n";
        
        Code += "\t"+"ldr X0, ="+addr+"\n";
        
        string op = (context.op.Text == "*")?"mul":"sdiv";
        if (isNUM(left)) {
	    Code += "\t"+"ldr X1, [X0]\n";
            if (isNUM(right)) {
                Code += "\t"+"add X1, X1, #"+left.AddrNUM+"\n";
                Code += "\t"+"mov X2, #"+right.AddrNUM+"\n";
                Code += "\t"+op+" X1, X1, X2\n";
                Code += "\t"+"str X1, [X0]\n";
            }
            else
            {
                Code += "\t"+"add X1, X1, #"+left.AddrNUM+"\n";
                Code += "\t"+"ldr X2, ="+right.AddrID+"\n";
                Code += "\t"+"ldr X3, [X2]\n";
                Code += "\t"+"mov X4, #"+right.AddrNUM+"\n";
                Code += "\t"+op+" X1, X3, X4\n";
                Code += "\t"+"str X1, [X0]\n";
            }
        }
        else
        {
            if (isNUM(right)) {
                Code += "\t"+"ldr X2, ="+left.AddrID+"\n";
                Code += "\t"+"ldr X3, [X2]\n";
                Code += "\t"+"mov X4, #"+right.AddrNUM+"\n";
                Code += "\t"+op+" X3, X3, X4\n";
                Code += "\t"+"str X3, [X0]\n";
            }
            else
            {
                Code += "\t"+"ldr X2, ="+left.AddrID+"\n";
                Code += "\t"+"ldr X3, [X2]\n";
                Code += "\t"+"ldr X4, ="+right.AddrID+"\n";
                Code += "\t"+"ldr X5, [X4]\n";
                Code += "\t"+op+" X3, X3, X5\n";
                Code += "\t"+"str X3, [X0]\n";
            }        
        }
        Code += "\n";
        return new Attr { AddrNUM = null, AddrID = addr };
    }

    public override Attr VisitPassE(GrammarParser.PassEContext context)
    {
        return Visit(context.e());
    }

    public override Attr VisitId(GrammarParser.IdContext context)
    {
        string id = context.ID().GetText();
        return new Attr { AddrNUM = null, AddrID = id };
    }

    public override Attr VisitNum(GrammarParser.NumContext context)
    {
    	string num = context.NUM().GetText(); 
        return new Attr { AddrNUM = num, AddrID = null };
    }
}

