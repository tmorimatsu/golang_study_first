#!/bin/bash

rm ans.txt
go run ../gopl/findlinks/fetch/fetch.go https://golang.org | go run ../gopl/findlinks/findlinks1.go > ans.txt
go run ../gopl/findlinks/fetch/fetch.go https://golang.org | go run main.go > practice.txt