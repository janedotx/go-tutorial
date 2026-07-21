package main

import "fmt"

// pointers that point to nothing are nil,
// because golang promises that *ptr<T> always
// returns something of type T and it can't do
// that if the pointer points to nothing
func TestNilPointer() {
	var s *string
	if (s == nil) {
		fmt.Println("that string points to nil")
	}
}