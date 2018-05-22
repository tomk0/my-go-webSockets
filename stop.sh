#!/bin/bash

input="pid.term"

pgrep -f node >> $input
pgrep -f npm >> $input
pgrep -f goServer >> $input

while IFS= read -r var
do
  kill $var
done < "$input"

rm $input
