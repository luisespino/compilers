#!/bin/bash
bison -d parser.y
flex lexer.l
gcc -o run parser.tab.c lex.yy.c main.c node.c -lfl
./run
rm parser.tab.c parser.tab.h lex.yy.c run
