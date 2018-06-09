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

efficient=`bash -c "go run efficient.go $s"`
echo "efficient: $efficient"
inefficient=`bash -c "go run inefficient.go $s"`
echo "inefficient: $inefficient"

diff=`expr $inefficient - $efficient`
ratio=`echo "scale=3; ($diff / $inefficient) * 100" | bc`
echo "$ratio% faster"
echo ""