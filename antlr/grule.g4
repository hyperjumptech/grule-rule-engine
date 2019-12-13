grammar grule;


// PARSER HERE

root
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
    : THEN  assignExpressions
    ;

assignExpressions
    : assignExpression+
    ;

assignExpression
    : assignment SEMICOLON
    | methodCall SEMICOLON
    | functionCall SEMICOLON
    ;

assignment
    : variable ASSIGN expression
    ;

expression
    : expression logicalOperator expression
    | LR_BRACKET expression logicalOperator expression RR_BRACKET
    | predicate
    ;

predicate
    : expressionAtom comparisonOperator expressionAtom
    | expressionAtom
    ;

expressionAtom
    : constant
    | variable
    | left=expressionAtom mathOperator right=expressionAtom
    | LR_BRACKET left=expressionAtom mathOperator right=expressionAtom RR_BRACKET
    | functionCall
    | methodCall
    ;

methodCall
    : DOTTEDNAME '(' functionArgs? ')'
    ;

functionCall
    : SIMPLENAME '(' functionArgs? ')'
    ;

functionArgs
    : (constant | variable | functionCall | methodCall | expression)
    (
      ','
      (constant | variable | functionCall | methodCall | expression)
    )*
    ;

logicalOperator
    : AND | OR
    ;

variable
    : SIMPLENAME | DOTTEDNAME
    ;

mathOperator
    : MUL | DIV | PLUS | MINUS
    ;

comparisonOperator
    : GT | LT | GTE | LTE | EQUALS | NOTEQUALS
    ;

constant
    : stringLiteral
    | decimalLiteral
    | '-' decimalLiteral
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
DOTTEDNAME                  : SIMPLENAME ( DOT SIMPLENAME )+ ;

PLUS                        : '+' ;
MINUS                       : '-' ;
DIV                         : '/' ;
MUL                         : '*' ;

EQUALS                      : '==' ;
ASSIGN                      : '=' ;
GT                          : '>' ;
LT                          : '<' ;
GTE                         : '>=' ;
LTE                         : '<=' ;
NOTEQUALS                   : '!=' ;

SEMICOLON                   : ';' ;
LR_BRACE                    : '{';
RR_BRACE                    : '}';
LR_BRACKET                  : '(';
RR_BRACKET                  : ')';
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