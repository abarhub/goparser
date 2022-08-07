// tinylang.g4
grammar Tinylang;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;
COMMA: ',';
TYPE_VOID: 'void';
TYPE_INT: 'int';
TYPE_STRING: 'string';
IDENT: [a-zA-Z][a-zA-Z0-9_]*;
PARENTHESIS_OPEN: '(';
PARENTHESIS_CLOSE: ')';
CURLY_BRACKET_OPEN: '{';
CURLY_BRACKET_CLOSE: '}';
SEMICOLON: ';';
EQUALS: '=';


// Rules
start :  function                  # StartFunction
        ;


function: type IDENT PARENTHESIS_OPEN PARENTHESIS_CLOSE CURLY_BRACKET_OPEN   instruction_list CURLY_BRACKET_CLOSE  # Funct
        ;

instruction_list: ( instruction SEMICOLON )*    # ListInstruction
        ;

instruction: id=IDENT EQUALS exp=expression                        # InstrAffect
        |  t=type id=IDENT  EQUALS exp=expression                    # InstrDeclare
        ;

type : TYPE_VOID       # TypeVoid
        | TYPE_INT     # TYPEINT
        | TYPE_STRING  # TypeString
        ;

expression
   : expression op=('*'|'/') expression # MulDiv
   | expression op=('+'|'-') expression # AddSub
   | NUMBER                             # Number
   | '(' expression ')'                 # Parenthesis
   ;
