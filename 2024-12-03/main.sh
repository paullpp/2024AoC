#!/bin/bash

DATA="$(cat inputs.txt | grep -E -o 'mul\([0-9]+,[0-9]+\)')"


IFS=$'\n' parsed_data="$DATA"

res=0

for mul in $parsed_data; do
  num1="$(echo "$mul" | grep -E -o -m1 '[0-9]+' | head -n1)"
  num2="$(echo "$mul" | grep -E -o -m2 '[0-9]+' | tail -n1)"

  intermediate=$((num1 * num2))
  res=$((res + intermediate))
done

echo $res

