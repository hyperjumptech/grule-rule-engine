grammar grulev2;


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
    : variable
    | functionCall
    ;

arrayMapSelector
    : LS_BRACKET expression RS_BRACKET
    ;

functionCall
    : SIMPLENAME '(' argumentList? ')'
    ;

argumentList
    :  expression ( ',' expression )*
    ;

variable
    : SIMPLENAME
    | constant
    | variable DOT functionCall
    | variable DOT SIMPLENAME
    | variable arrayMapSelector
    ;


constant
    : stringLiteral
    | decimalLiteral
    | booleanLiteral
    | realLiteral
    | NOT? NULL_LITERAL
    ;

decimalLiteral
    : MINUS? DECIMAL_LITERAL
    ;

realLiteral
    : MINUS? REAL_LITERAL
    ;

stringLiteral
    : DQUOTA_STRING | SQUOTA_STRING
    ;

booleanLiteral
    : TRUE | FALSE;

// LEXER HERE

fragment DEC_DIGIT          : [0-9];
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
fragment EXPONENT_NUM_PART  : 'E' '-'? DEC_DIGIT+;

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

SIMPLENAME                  : [a-zA-Z] [a-zA-Z0-9]* ;

PLUS                        : '+' ;
MINUS                       : '-' ;
DIV                         : '/' ;
MUL                         : '*' ;
MOD                         : '%' ;

EQUALS                      : '==' ;
ASSIGN                      : '=' ;
GT                          : '>' ;
LT                          : '<' ;
GTE                         : '>=' ;
LTE                         : '<=' ;
NOTEQUALS                   : '!=' ;

BITAND                      : '&';
BITOR                       : '|';

SEMICOLON                   : ';' ;
LR_BRACE                    : '{';
RR_BRACE                    : '}';
LR_BRACKET                  : '(';
RR_BRACKET                  : ')';
LS_BRACKET                  : '[';
RS_BRACKET                  : ']';
DOT                         : '.' ;
DQUOTA_STRING               : '"' ( '\\'. | '""' | ~('"'| '\\') )* '"';
SQUOTA_STRING               : '\'' ('\\'. | '\'\'' | ~('\'' | '\\'))* '\'';

DECIMAL_LITERAL             : DEC_DIGIT+;

REAL_LITERAL                : (DEC_DIGIT+)? '.' DEC_DIGIT+
                            | DEC_DIGIT+ '.' EXPONENT_NUM_PART
                            | (DEC_DIGIT+)? '.' (DEC_DIGIT+ EXPONENT_NUM_PART)
                            | DEC_DIGIT+ EXPONENT_NUM_PART
                            ;

SPACE                       : [ \t\r\n]+ {l.Skip()}
                            ;

COMMENT                     : '/*' .*? '*/' {l.Skip()}
                            ;

LINE_COMMENT                : '//' ~[\r\n]* {l.Skip()}
                            ;