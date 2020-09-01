# ANTLR

## Making ANTLR alias.

```bash
alias antlr='java -jar ~/Laboratory/Antlr/antlr-4.7-complete.jar'
```

## Executing ANTLR

```bash
antlr -Dlanguage=Go -o parser -package grulev2 -lib . -listener -visitor grulev2.g4
```