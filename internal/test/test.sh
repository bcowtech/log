#!/bin/bash
ls *_test.go | xargs -n1 go test -v $1