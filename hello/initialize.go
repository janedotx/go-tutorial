package main

type TestStructForFun struct {
	Count string
	Results []string
}

// an empty struct
func HowToReturnEmptyStruct() TestStructForFun {
	return TestStructForFun{}
}