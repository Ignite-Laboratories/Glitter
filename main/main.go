package main

type MyStruct struct {
	Test int
}

func MyMethod[T *MyStruct]() *T {
	var v T
	v.Test = 1
	return &v
}
