# GNU Make >= 4.3

## Configuration (See ../ANTLR.md)

OUTPUT_DIR = parser/grulev3

ANTLR = antlr4
ANTLRFLAGS = -Dlanguage=Go -o '${OUTPUT_DIR}' -package '$(notdir ${OUTPUT_DIR})'

GRAMMAR = grulev3.g4

## Rules

PREFIX = ${OUTPUT_DIR}/$(basename ${GRAMMAR})
tokens = $(addsuffix .interp,$(1)) $(addsuffix .tokens,$(1))

TOKENS = $(call tokens,${PREFIX}Lexer) $(call tokens,${PREFIX})
LEXER = ${PREFIX}_lexer.go
PARSER = ${PREFIX}_parser.go ${PREFIX}_listener.go ${PREFIX}_base_listener.go ${PREFIX}_visitor.go ${PREFIX}_base_visitor.go

.PHONY: clean all
.DELETE_ON_ERROR:

all: ${TOKENS} ${LEXER} ${PARSER}

${TOKENS} ${LEXER} ${PARSER} &: ${GRAMMAR}
	$(ANTLR) ${ANTLRFLAGS} -visitor $^

clean:
	rm -f ${TOKENS} ${LEXER} ${PARSER}
