grammar grulev3;

// PARSER HERE
grl
    : ruleEntry* EOF
    ;

ruleEntry
    : RULE ruleName ruleDescription? salience? LR_BRACE whenScope thenScope RR_BRACE
    ;

salience
    : SALIENCE decimalLiteral
    ;

ruleName
    : SIMPLENAME
    ;

ruleDescription
    : DQUOTA_STRING | SQUOTA_STRING
    ;

whenScope
    : WHEN  expression
    ;

thenScope
    : THEN  thenExpressionList
    ;

thenExpressionList
    : thenExpression+
    ;

thenExpression
    : assignment SEMICOLON
    | functionCall SEMICOLON
    | variable SEMICOLON
    ;

assignment
    : variable ASSIGN expression
    ;

expression
    : expression mulDivOperators expression
    | expression addMinusOperators expression
    | expression comparisonOperator expression
    | expression andLogicOperator expression
    | expression orLogicOperator expression
    | LR_BRACKET expression RR_BRACKET
    | expressionAtom
    ;

mulDivOperators
    : MUL | DIV | MOD
    ;

addMinusOperators
    : PLUS | MINUS | BITAND | BITOR
    ;

comparisonOperator
    : GT | LT | GTE | LTE | EQUALS | NOTEQUALS
    ;

andLogicOperator
    : AND
    ;

orLogicOperator
    : OR
    ;

expressionAtom
    : constant
    | functionCall
    | variable
    ;

constant
    : stringLiteral
    | decimalLiteral
    | booleanLiteral
    | realLiteral
    ;

variable
    : SIMPLENAME
    | variable MEMBERVARIABLE
    | variable arrayMapSelector
    | variable METHOD_SIGNATURE  argumentList? RR_BRACKET
    ;

arrayMapSelector
    : LS_BRACKET expression RS_BRACKET
    ;

functionCall
    : FUNCTION_SIGNATURE  argumentList? RR_BRACKET
    ;

argumentList
    :  expression ( ',' expression )*
    ;

realLiteral
    : MINUS? FLOAT_LIT
    ;

decimalLiteral
    : MINUS? INT_LIT
    ;

stringLiteral
    : DQUOTA_STRING | SQUOTA_STRING
    ;

booleanLiteral
    : TRUE | FALSE;

// LEXER HERE
fragment A                  : [aA] ;
fragment B                  : [bB] ;
fragment C                  : [cC] ;
fragment D                  : [dD] ;
fragment E                  : [eE] ;
fragment F                  : [fF] ;
fragment G                  : [gG] ;
fragment H                  : [hH] ;
fragment I                  : [iI] ;
fragment J                  : [jJ] ;
fragment K                  : [kK] ;
fragment L                  : [lL] ;
fragment M                  : [mM] ;
fragment N                  : [nN] ;
fragment O                  : [oO] ;
fragment P                  : [pP] ;
fragment Q                  : [qQ] ;
fragment R                  : [rR] ;
fragment S                  : [sS] ;
fragment T                  : [tT] ;
fragment U                  : [uU] ;
fragment V                  : [vV] ;
fragment W                  : [wW] ;
fragment X                  : [xX] ;
fragment Y                  : [yY] ;
fragment Z                  : [zZ] ;

fragment UNDER_SCORE        : '_';
fragment DEC_DIGIT          : [0-9];
fragment HEX_DIGIT          : [0-9a-fA-F];

PLUS                        : '+' ;
MINUS                       : '-' ;
DIV                         : '/' ;
MUL                         : '*' ;
MOD                         : '%' ;
DOT                         : '.' ;
SEMICOLON                   : ';' ;

LR_BRACE                    : '{';
RR_BRACE                    : '}';
LR_BRACKET                  : '(';
RR_BRACKET                  : ')';
LS_BRACKET                  : '[';
RS_BRACKET                  : ']';

RULE                        : R U L E  ;
WHEN                        : W H E N ;
THEN                        : T H E N ;
AND                         : '&&' ;
OR                          : '||' ;
TRUE                        : T R U E ;
FALSE                       : F A L S E ;
NULL_LITERAL                : N U L L ;
NOT                         : N O T ;
SALIENCE                    : S A L I E N C E ;

EQUALS                      : '==' ;
ASSIGN                      : '=' ;
GT                          : '>' ;
LT                          : '<' ;
GTE                         : '>=' ;
LTE                         : '<=' ;
NOTEQUALS                   : '!=' ;

BITAND                      : '&';
BITOR                       : '|';

METHOD_SIGNATURE            : DOT FUNCTION_SIGNATURE;
FUNCTION_SIGNATURE          : SIMPLENAME LR_BRACKET;
MEMBERVARIABLE              : DOT SIMPLENAME ;
SIMPLENAME                  : [a-zA-Z] [a-zA-Z0-9]*;

DQUOTA_STRING               : '"' ( '\\'. | '""' | ~('"'| '\\') )* '"';
SQUOTA_STRING               : '\'' ('\\'. | '\'\'' | ~('\'' | '\\'))* '\'';

FLOAT_LIT                   : DECIMAL_FLOAT_LIT | HEX_FLOAT_LIT ;
DECIMAL_FLOAT_LIT           : DEC_LIT DOT DEC_LIT? DECIMAL_EXPONENT?
                            | DEC_LIT DECIMAL_EXPONENT
                            | DOT DEC_LIT DECIMAL_EXPONENT?
                            ;
DECIMAL_EXPONENT            : E (PLUS|MINUS)? DEC_DIGITS;

HEX_FLOAT_LIT               : '0' X HEX_MANTISA HEX_EXPONENT
                            ;
fragment HEX_MANTISA        : UNDER_SCORE? HEX_DIGITS DOT HEX_DIGITS?
                            | UNDER_SCORE? HEX_DIGITS
                            | DOT HEX_DIGITS
                            ;
HEX_EXPONENT                : P (PLUS|MINUS)? DEC_DIGITS
                            ;

INT_LIT                     : DEC_LIT | HEX_LIT ;
DEC_LIT                     : '0'
                            | [1-9] (UNDER_SCORE? DEC_DIGITS)*
                            ;

HEX_LIT                     : '0' X HEX_DIGITS;
HEX_DIGITS                  : HEX_DIGIT+;
DEC_DIGITS                  : DEC_DIGIT+;


// IGNORED TOKENS
SPACE                       : [ \t\r\n]+ {l.Skip()};
COMMENT                     : '/*' .*? '*/' {l.Skip()};
LINE_COMMENT                : '//' ~[\r\n]* {l.Skip()};