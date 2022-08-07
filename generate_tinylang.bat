SET ANTLR4=.\antlr-4.10.1-complete.jar
rem %ANTLR4% -Dlanguage=Go -no-listener -visitor -o tinylang\parser -Xlog Tinylang.g4
%ANTLR4% -Dlanguage=Go -o tinylang\parser Tinylang.g4