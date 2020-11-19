#!/bin/bash

[! -d "example_words.txt"] && touch "example_words .txt"

RANDOM= $$
for i in $@
do 
    cat "$@*$RANDOM " >> "example.txt"
done