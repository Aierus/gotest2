#!/bin/bash

for i in $(seq 1 1 8)
do
    ./GoQuiz2 $((10**i))
done
