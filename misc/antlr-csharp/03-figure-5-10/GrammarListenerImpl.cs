using Antlr4.Runtime.Misc;
using System;
using System.Collections.Generic;

public class Node
{
    public int id;
    public string label;
    public Node left;
    public Node right;
}
    
public class Tree
{
    public Node root;
    public string dot = "graph{\n";
    
    public void Inorder(Node tmp)
    {
        if (tmp != null)
        {
            Inorder(tmp.left);
            Console.Write(tmp.label + " ");
            Inorder(tmp.right);
        }
    }
    
    public void Postorder(Node tmp)
    {
        if (tmp != null)
        {
            Postorder(tmp.left);
            Postorder(tmp.right);
            Console.Write(tmp.label + " ");
            dot += tmp.id.ToString() + "[label=\"" + tmp.label + "\"];\n";
            if (tmp.left != null)
            {
                dot += tmp.id.ToString() + "--" + tmp.left.id.ToString() + ";\n";
            }
            if (tmp.right != null)
            {
                dot += tmp.id.ToString() + "--" + tmp.right.id.ToString() + ";\n";
            }
        }
    }
}

public class GrammarListenerImpl : GrammarBaseListener
{
    public int id = 0;
    
    private Stack<Node> stack = new Stack<Node>();
    
    public override void ExitSumres(GrammarParser.SumresContext context)
    {
        Node right = stack.Pop();
        Node left = stack.Pop();
        stack.Push(new Node
        { 
            id = id++,
            label = context.op.Text,
            right = right,
            left = left
        });
    }
    
    public override void ExitMuldiv(GrammarParser.MuldivContext context)
    {
        Node right = stack.Pop();
        Node left = stack.Pop();
        stack.Push(new Node
        {   
            id = id++,
            label = context.op.Text,
            right = right,
            left = left 
        });
    }
    
    public override void ExitNum(GrammarParser.NumContext context)
    {
        stack.Push(new Node
        {
            id = id++,
            label = context.GetText(),
            right = null,
            left = null
        });
    }
    
    public override void ExitId(GrammarParser.IdContext context)
    {
        stack.Push(new Node
        {
            id = id++,
            label = context.GetText(),
            right = null,
            left = null
        });
    }
    
    public Node GetResult()
    {
        return stack.Count > 0 ? stack.Pop() : null;
    }
}

