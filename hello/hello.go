package main

import (
	"fmt"
	"example.com/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	messages, err := greetings.Hellos([]string{"gladys", "sammy", "darin"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}