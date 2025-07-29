using Antlr4.Runtime;
using System;

class Program
{
    static void Main(string[] args)
    {
        while (true)
        {
            var listener = new GrammarListenerImpl();
            Console.Write("Enter a expression or exit): ");
            string input = Console.ReadLine();
            
            if (input.ToLower() == "exit")
            {
                break;
            }
            
            AntlrInputStream inputStream = new AntlrInputStream(input);
            
            GrammarLexer lexer = new GrammarLexer(inputStream);
            
            CommonTokenStream commonTokenStream = new CommonTokenStream(lexer);
            
            GrammarParser parser = new GrammarParser(commonTokenStream);
            
            parser.AddParseListener(listener);
            
            parser.e();
            
            Tree tree = new Tree
            {
                root = listener.GetResult()
            };
            
            Console.WriteLine("Inorder:");
            tree.Inorder(tree.root);
            Console.WriteLine("\n\nPostorder:");
            tree.Postorder(tree.root);
            Console.WriteLine("\n\nDot:");
            Console.WriteLine(tree.dot + "}");
        }
    }
}

