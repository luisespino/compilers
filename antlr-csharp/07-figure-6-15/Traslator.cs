using System;
using System.Collections.Generic;
using System.Linq;
using Antlr4.Runtime;
using Antlr4.Runtime.Tree;

public class Attr
{
    public string str;
    public Object[] obj;
}

public class Variable
{
    public string Name;
    public object[] Expression;
    public string Type;
}

public class Translator : GrammarBaseVisitor<Attr>
{
    public string Data = ".global _start\n\n.data\n";
    public string Text = "\n.text\n_start:\n";
    Dictionary<string, Variable> variables = new Dictionary<string, Variable>();

    static int[] GetDimensions(object[] lst)
    {
        if (lst is object[] innerArray)
        {
            if (innerArray.Length > 0)
            {
                return new[] { innerArray.Length }.Concat(GetDimensions(innerArray[0] as object[])).ToArray();
            }
            else
            {
                return new[] { 0 };
            }
        }
        else
        {
            return new int[] { };
        }
    }
    
    static string[] FlattenArray(object[] arr)
    {
        var result = new List<string>();  
        FlattenHelper(arr, result);  
        return result.ToArray(); 
    }

    static void FlattenHelper(object[] arr, List<string> result)
    {
        foreach (var item in arr)
        {
            if (item is object[] innerArray) 
                FlattenHelper(innerArray, result); 
            else if (item is string number) 
                result.Add(number); 
        }
    }
    
    public override Attr VisitProgram(GrammarParser.ProgramContext context)
    {
        foreach (var stmtContext in context.stmt())
        {
            Visit(stmtContext);
        }
        Text += "\tmov X8, #93\n\tsvc #0\n";
        return new Attr { str = Data+Text, obj = null };
    }

    public override Attr VisitDeclExpr(GrammarParser.DeclExprContext context)
    {
        string variable = Visit(context.var()).str;
        string type = context.type.Text;
        object[] expression = Visit(context.expr()).obj;

        variables[variable] = new Variable 
        {
            Name = variable, 
            Expression = expression, 
            Type = type
        };

        int[] dimensions = GetDimensions(expression);
        string[] value = FlattenArray(expression);
        Data += variable + ": .word ";
        string nums = string.Join(", ", value);
        if (string.IsNullOrEmpty(nums))
        {
            Data += "0\n";
        }
        else
        {
            Data += nums + "\n";
        }

        if (dimensions.Length == 1)
        {
            Data += variable + "_cols: .word " + dimensions[0] + "\n";
        }
        else if (dimensions.Length == 2)
        {
            Data += variable + "_rows: .word " + dimensions[0] + "\n";
            Data += variable + "_cols: .word " + dimensions[1] + "\n";
        }
        else if (dimensions.Length == 3)
        {
            Data += variable + "_face: .word " + dimensions[0] + "\n";
            Data += variable + "_rows: .word " + dimensions[1] + "\n";
            Data += variable + "_cols: .word " + dimensions[2] + "\n";
        }
        
        return new Attr { str = null, obj = null };
    }

    public override Attr VisitVarName(GrammarParser.VarNameContext context)
    {
        string name = context.NAME().GetText();
        return new Attr { str = name, obj = null };
    }

    public override Attr VisitExprEmpty(GrammarParser.ExprEmptyContext context)
    {
        return new Attr { str = null, obj = new object[] { } };
    }

    public override Attr VisitExprValues(GrammarParser.ExprValuesContext context)
    {
        object[] list = Visit(context.list()).obj;
        return new Attr { str = null, obj = list };
    }

    public override Attr VisitValues(GrammarParser.ValuesContext context)
    {
        Attr val = Visit(context.val());
        Object[] list = Visit(context.list()).obj;
        Object[] value = null;
        if (val.str != null)
            value = new [] {val.str};
        else
            value = new [] {val.obj};
        return new Attr { str = null, obj = value.Concat(list).ToArray() };
    }

    public override Attr VisitValue(GrammarParser.ValueContext context)
    {
        Attr val = Visit(context.val());
        if (val.str != null)
            return new Attr { str = null, obj = new [] {val.str} };
        return new Attr { str = null, obj = new [] {val.obj} };
    }

    public override Attr VisitValDigit(GrammarParser.ValDigitContext context)
    {
        string digit = context.DIGIT().GetText();
        return new Attr { str = digit, obj = null};
    }

    public override Attr VisitValExpr(GrammarParser.ValExprContext context)
    {
        Object[] expr = Visit(context.expr()).obj;
        return new Attr { str = null, obj = expr};
    }
}
