using System;
using Antlr4.Runtime;

public class Program
{
    public static void Main(string[] args)
    {
        GrammarBaseVisitorImpl visitor = new GrammarBaseVisitorImpl();
        
        Console.WriteLine("Type a declaration like int[3] or 'exit':");
        while (true)
        {
            Console.Write("> ");
            string input = Console.ReadLine();

            if (input.ToLower() == "exit")
            {
                break;
            }

            AntlrInputStream inputStream = new AntlrInputStream(input);

            GrammarLexer lexer = new GrammarLexer(inputStream);
            CommonTokenStream tokenStream = new CommonTokenStream(lexer);

            GrammarParser parser = new GrammarParser(tokenStream);
            GrammarParser.TContext tree = parser.t();
            string result = visitor.Visit(tree);
            Console.WriteLine(result);
        }

    }
}

