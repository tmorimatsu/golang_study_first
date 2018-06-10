#!/bin/bash

bash -c "rm out.txt"
bash -c "rm -r contents"
bash -c "mkdir contents"
url="https://www.airbnb.com"
for i in {0..5}
do
    tmp=$url$i
    cfname=${tmp////_}
    ret=`bash -c "go build main.go && ./main contents/$cfname.txt $url"`
    echo $ret
done

# contents/i.txtの内容比較
echo "if nothing is displayed after this, all files are the same"
for i in {0..5}
do
    for n in {0..5}
    do
        tmp=$url$i
        cfname1=${tmp////_}
        tmp=$url$n
        cfname2=${tmp////_}
        if [ "$i" -lt "$n" ] 
        then
            bash -c "diff -q contents/$cfname1.txt contents/$cfname2.txt"
        fi
    done
done

# todo: 複数のurlに対応(時間があれば)
