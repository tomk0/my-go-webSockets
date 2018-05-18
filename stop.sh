#!/bin/bash

input="pid.term"

pgrep -f xterm >> $input

while IFS= read -r var
do
  kill $var
done < "$input"

rm $input
