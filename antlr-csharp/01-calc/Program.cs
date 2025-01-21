using Antlr4.Runtime;
using System;

class Program
{
    static void Main(string[] args)
    {
        var listener = new CalcListenerImpl();
        Console.WriteLine("Enter a expression or exit: ");
        while (true)
        {
            Console.Write("> ");
            string input = Console.ReadLine();

            if (input.ToLower() == "exit")
            {
                break;
            }

            AntlrInputStream inputStream = new AntlrInputStream(input);

            CalcLexer lexer = new CalcLexer(inputStream);

            CommonTokenStream commonTokenStream = new CommonTokenStream(lexer);

            CalcParser parser = new CalcParser(commonTokenStream);

            parser.AddParseListener(listener);

            parser.e(); 
            
            double result = listener.GetResult();
            Console.WriteLine(result);
        }
    }
}
