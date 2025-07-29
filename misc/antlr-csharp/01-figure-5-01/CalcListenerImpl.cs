using Antlr4.Runtime.Misc;
using System;
using System.Collections.Generic;

public class CalcListenerImpl : CalcBaseListener
{
    private Stack<double> stack = new Stack<double>();

    public override void ExitSumres(CalcParser.SumresContext context)
    {
        double right = stack.Pop();
        double left = stack.Pop();
        double result = context.op.Text == "+" ? left + right : left - right;
        stack.Push(result);
    }


    public override void ExitMuldiv(CalcParser.MuldivContext context)
    {
        double right = stack.Pop();
        double left = stack.Pop();
        double result = context.op.Text == "*" ? left * right : left / right;
        stack.Push(result);
    }

    public override void ExitNum(CalcParser.NumContext context)
    {
        double num = double.Parse(context.GetText());
        stack.Push(num);
    }

    public double GetResult()
    {
        return stack.Count > 0 ? stack.Pop() : 0;
    }
}
