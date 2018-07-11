#!/bin/bash

# main.goが相対pathのため別ディレクトリから叩くと動かない
bash -c "go run main.go github.com/tmorimatsu/golang_study_first/blob/master/README.md"
