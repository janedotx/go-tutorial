package main

import "fmt"

func TestWhileLoops () {
	i := false
	j := 0
	for ; !(i); {
		if (j < 10) {
			j++;
		} else {
			i = true
			break
		}
		fmt.Println(j)
		fmt.Println("\n")
	} 
}