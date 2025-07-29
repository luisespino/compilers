using System;
using Antlr4.Runtime;

class Program
{
    public static void Main(string[] args)
    {
        while (true)
        {
            Translator translator = new Translator();
            Console.Write("Enter statements or exit: ");
            string input = Console.ReadLine();
            // test_input = a = 1  b = 2  c = a + b * 2   d = (c + 1) * 3

            if (input.ToLower() == "exit")
            {
                break;
            }

            // input stream
            AntlrInputStream inputStream = new AntlrInputStream(input);
            GrammarLexer lexer = new GrammarLexer(inputStream);
            CommonTokenStream tokens = new CommonTokenStream(lexer);
            GrammarParser parser = new GrammarParser(tokens);

            // get syntax tree root
            GrammarParser.PContext tree = parser.p();
            string result = translator.Visit(tree).Code;

            Console.WriteLine("Intermediate code:\n" + result);
        }
    }
}

