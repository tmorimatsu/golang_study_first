#!/bin/bash

go run ./fetch/fetch.go https://golang.org | go run findlinks1.go