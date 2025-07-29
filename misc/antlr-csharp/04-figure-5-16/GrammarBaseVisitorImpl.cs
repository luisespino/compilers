using System;

public class GrammarBaseVisitorImpl : GrammarBaseVisitor<string>
{
    private string baseType;

    public override string VisitT(GrammarParser.TContext context)
    {
    	baseType = Visit(context.b()); 
        string dimensions = Visit(context.c()); 
        return dimensions;
    }

    public override string VisitB(GrammarParser.BContext context)
    {
        return context.GetText() == "int" ? "integer" : "float";
    }

    public override string VisitC(GrammarParser.CContext context)
    {
        if (context.c() != null)
        {
            string currentDimension = context.n().NUM().GetText();
            string nestedDimensions = Visit(context.c());
            return $"array({currentDimension}, {nestedDimensions})";
        }
        else
        {
            return baseType;
        }
    }

    public override string VisitN(GrammarParser.NContext context)
    {
        return context.NUM().GetText();
    }
}

