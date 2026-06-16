package main

import (
	"fmt"
	"os"
)

func ReadFiles() {
	dat, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("error reading file:", err)
		os.Exit(1)
	}
	fmt.Println(string(dat))
}

func WriteFiles() {
	f, err := os.Create("test.txt")

	if err != nil {
		fmt.Println("error creating file:", err)
		os.Exit(1)
	}
	defer f.Close()

}