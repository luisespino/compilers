# Example of using ANTLR for C# on Arch Linux

The main requirements are .NET, OpenJDK, and ANTLR4:

## Install .NET on Arch Linux

Upgrade the system:

```
sudo pacman -Syu
```

Install dependencies:

```
sudo pacman -Sy \
    glibc \
    gcc \
    krb5 \
    icu \
    openssl \
    libc++ \
    zlib
```

Download the installation script:

```
wget https://dot.net/v1/dotnet-install.sh -O dotnet-install.sh
```

Grant execute permission to the script:

```
chmod +x ./dotnet-install.sh
```

Run the script to install the latest version:
```
./dotnet-install.sh --version latest
```

Install the .NET Runtime:
```
./dotnet-install.sh --version latest --runtime aspnetcore
```

You can install a specific major version:
```
./dotnet-install.sh --channel 9.0
```

Set DOTNET_ROOT environment variable:
```
export DOTNET_ROOT=$HOME/.dotnet
```

Update the PATH environment variable:
```
export PATH=$PATH:$DOTNET_ROOT:$DOTNET_ROOT/tools
```

Check the installed version:
```
dotnet --version
```

![Alt text](https://github.com/luisespino/compilers/blob/master/antlr-csharp/01-calc/screen1.png?raw=true "version")

## Check or install OpenJDK


Check Java version:

```
java -version
```

If OpenJDK is not installed, install it:
```
sudo pacman -S jdk-openjdk
```

![Alt text](https://github.com/luisespino/compilers/blob/master/antlr-csharp/01-calc/screen2.png?raw=true "version")


## Install ANTLR4

Install ANTLR4 via the following command:
```
sudo pacman -S antlr4
```
## Create a C# project using ANTLR

Create a new project:
```
dotnet new console -n Calc
cd Calc
```
Add the ANTLR dependency:
```
dotnet add package Antlr4.Runtime.Standard
```
Write the grammar [Calc.g4](https://github.com/luisespino/compilers/blob/master/antlr-csharp/01-calc/Calc.g4), for a simple calculator:
```
grammar Calc;

// tokens
NUM : [0-9]+;
WS  : [ \r\n\t]+ -> skip;

// rules
e : e op=('+'|'-') t    # Sumres
  | t                   # PassT
  ;

t : t op=('*'|'/') f    # Muldiv
  | f                   # PassF
  ;

f : '(' e ')'           # PassE
  | ID                  # Id 
  | NUM                 # Num
  ;

```

Generate the C# files with ANTLR:
```
antlr4 -Dlanguage=CSharp Calc.g4
```

The following files will be generated:
- CalcLexer.cs
- CalParser.cs
- CalcListener.cs
- CalcBaseListener.cs

Create the [CalcListenerImpl.cs](https://github.com/luisespino/compilers/blob/master/antlr-csharp/01-calc/CalcListenerImpl.cs) file:
```
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
```

Modify the [Program.cs](https://github.com/luisespino/compilers/blob/master/antlr-csharp/01-calc/Program.cs) file:
```
using Antlr4.Runtime;
using System;

class Program
{
    static void Main(string[] args)
    {
        var listener = new CalcListenerImpl();

        while (true)
        {
            Console.Write("Enter a expression or exit): ");
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
            Console.WriteLine("Resultado: " + result);
        }
    }
}
```

Ready! Now build and run the project:
```
dotnet build
dotnet run
````
Now, the process is complete, and your calculator will evaluate expressions entered by the user!

![Alt text](https://github.com/luisespino/compilers/blob/master/antlr-csharp/01-calc/screen2.png?raw=true "version")
