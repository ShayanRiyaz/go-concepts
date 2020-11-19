#!/bin/bash
[ ! -d "exampledata" ] && mkdir -p "exampledata"

echo ------------
echo Generate example data
echo ------------

for num in $(seq 1 $1)
do
    echo "THIS is an example file $num" >> "exampledata/example$num.txt"
done

echo ------------
echo Generated example data
echo ------------
