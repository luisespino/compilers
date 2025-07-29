using Antlr4.Runtime.Misc;
using System;
using System.Collections.Generic;

public class GrammarListenerImpl : GrammarBaseListener
{
    private Stack<double> stack = new Stack<double>();

    public override void ExitSum(GrammarParser.SumContext context)
    {
        double right = stack.Pop();
        double left = stack.Pop();
        double result = left + right;
        stack.Push(result);
    }


    public override void ExitMul(GrammarParser.MulContext context)
    {
        double right = stack.Pop();
        double left = stack.Pop();
        double result = left * right;
        stack.Push(result);
    }

    public override void ExitDigit(GrammarParser.DigitContext context)
    {
        double num = double.Parse(context.GetText());
        stack.Push(num);
    }

    public double GetResult()
    {
        return stack.Count > 0 ? stack.Pop() : 0;
    }
}
