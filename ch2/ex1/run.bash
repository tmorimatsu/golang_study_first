#!/bin/sh

# 引数は半角数字ならなんでもok
bash -c "go run main.go 100 c"

# これはエラー
# bash -c "go run main.go -10 k"