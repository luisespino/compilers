using Antlr4.Runtime;
using System;

class Program
{
    static void Main(string[] args)
    {
        var listener = new GrammarListenerImpl();
        
        Console.WriteLine("Enter a expression with (+/*) or exit): ");
            

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

            CommonTokenStream commonTokenStream = new CommonTokenStream(lexer);

            GrammarParser parser = new GrammarParser(commonTokenStream);

            parser.AddParseListener(listener);

            parser.e(); 
            
            double result = listener.GetResult();
            Console.WriteLine("Resultado: " + result);
        }
    }
}
