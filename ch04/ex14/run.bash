#!/bin/bash

rm out.html
go run main.go repo:golang/go is:open json decoder > out.html