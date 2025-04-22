using System;
using Antlr4.Runtime;

class Program
{
    public static void Main(string[] args)
    {
        while (true)
        {
            Console.Write("Enter filename or exit: ");
            string fileName = Console.ReadLine();

            if (fileName.ToLower() == "exit")
            {
                break;
            }

            if (File.Exists(fileName))
            {
                try
                {
                    string fileContent = File.ReadAllText(fileName);

                    Translator translator = new Translator();
                    
                    // input stream
                    AntlrInputStream inputStream = new AntlrInputStream(fileContent);
                    GrammarLexer lexer = new GrammarLexer(inputStream);
                    CommonTokenStream tokens = new CommonTokenStream(lexer);
                    GrammarParser parser = new GrammarParser(tokens);

                    // get syntax tree root
                    GrammarParser.ProgContext tree = parser.prog();
                    string result = translator.Visit(tree).Code;

                    Console.WriteLine("ARM64 code:\n" + result);
                }
                catch (Exception ex)
                {
                    Console.WriteLine($"Reading file error: {ex.Message}");
    		    Console.WriteLine("StackTrace: " + ex.StackTrace);
                }
            }
            else
            {
                Console.WriteLine("File does not exist.");
            }
        }
    }
}
