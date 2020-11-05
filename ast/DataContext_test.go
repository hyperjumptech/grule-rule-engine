package ast

import (
	"fmt"
)

type TestAStruct struct {
	BStruct *TestBStruct
}

type TestBStruct struct {
	CStruct *TestCStruct
}

type TestCStruct struct {
	Str string
	It  int
}

func (tcs *TestCStruct) EchoMethod(s string) {
	fmt.Println(s)
}

func (tcs *TestCStruct) EchoVariad(ss ...string) int {
	for _, s := range ss {
		fmt.Println(s)
	}
	return len(ss)
}
