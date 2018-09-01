#!/usr/bin/env bash

rm test_result.html
go run main.go < test_raw.html > test_result.html
diff -q test_answer.html test_result.html
