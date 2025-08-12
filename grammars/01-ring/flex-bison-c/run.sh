#!/bin/bash

# Compile
bison -d parser.y
gcc -o run parser.tab.c 

# Run
if [ -z "$1" ]; then
    ./run
else
    ./run "$1"
fi

# Clean
rm parser.tab.c parser.tab.h run