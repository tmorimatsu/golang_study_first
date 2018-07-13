#!/bin/sh

if [ -n "$1" ]
then
    args_len=$1
else
    args_len=10000
fi

echo ""
echo "args_len: $args_len"

s=`ruby -e "puts 'a ' * $args_len"`
bash -c "go run main.go $s"