#!/bin/sh

# 引数は半角数字ならなんでもok
num="1145111"
echo $num
bash -c "go run main.go $num"