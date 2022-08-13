// tinylang.g4
grammar Tinylang;

// Tokens
WHITESPACE: [ \r\n\t]+ -> skip;

LINE_COMMENT: '//' .*? '\r'? '\n' -> skip ;
COMMENT: '/*' .*? '*/'  -> skip ;

MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
MOD: '%';
PARENTHESIS_OPEN: '(';
PARENTHESIS_CLOSE: ')';
CURLY_BRACKET_OPEN: '{';
CURLY_BRACKET_CLOSE: '}';
SEMICOLON: ';';
EQUALS_TEST: '==';
NOT_EQUALS_TEST: '!=';
GREATER_THAN: '>';
GREATER_OR_EQUALS: '>=';
LESS_THAN: '<';
LESS_OR_EQUALS: '<=';
EQUALS: '=';
COMMA: ',';
AND_TEST: '&&';
OR_TEST: '||';
NOT_TEST: '!';

// number
NUMBER: [0-9]+;

// ident
TYPE_VOID: 'void';
TYPE_INT: 'int';
TYPE_STRING: 'string';
TYPE_BOOLEAN: 'boolean';
VALUE_TRUE: 'true';
VALUE_FALSE: 'false';
IDENT: [a-zA-Z][a-zA-Z0-9_]*;

STRING : '"' ~('"')* '"' ;

// Rules
start :  function                  # StartFunction
        ;


function: t=type name=IDENT PARENTHESIS_OPEN PARENTHESIS_CLOSE CURLY_BRACKET_OPEN   instrlist=instruction_list CURLY_BRACKET_CLOSE  # Funct
        ;

instruction_list: ( instruction SEMICOLON )*    # ListInstruction
        ;

instruction: id=IDENT EQUALS exp=expression                        # InstrAffect
        |  t=type id=IDENT  EQUALS exp=expression                    # InstrDeclare
        ;

type : TYPE_VOID       # TypeVoid
        | TYPE_INT     # TypeInt
        | TYPE_STRING  # TypeString
        | TYPE_BOOLEAN  # TypeBoolean
        ;

expression
   : expression op=('*'|'/'|'%') expression # MulDiv
   | expression op=('+'|'-') expression # AddSub
   | expression op=('>'|'>='|'<'|'<='|'=='|'!=') expression # Compare
   | expression op=('&&'|'||') expression # AndOr
   | op='!' expression                  # Not
   | e=NUMBER                             # Number
   | e=VALUE_TRUE                                 # True
   | e=VALUE_FALSE                                 # False
   | e=IDENT                                 # ExprIdent
   | e=STRING                                # ExprString
   | '(' expression ')'                 # Parenthesis
   ;
