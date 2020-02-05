#!/bin/bash

function usage {
	echo "new-day.sh 05"
	exit 1
}

if [ -z "$1" ]; then
	usage 
fi

DAY=$1

dir="day${DAY}"
mkdir $dir
pushd $dir

echo 'package main

import "fmt"

func Run() string {
	return ""
}

func main() {
	fmt.Println(Run())
}
' > main.go

echo 'package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	a := assert.New(t)

	a.Equal("", Run())
}
' > main_test.go


