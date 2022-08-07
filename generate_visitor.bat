SET ANTLR4=.\antlr-4.10.1-complete.jar
%ANTLR4% -Dlanguage=Go -no-listener -visitor -o visitor/parser Calc.g4