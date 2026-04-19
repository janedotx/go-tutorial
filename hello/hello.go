package main

import (
	//	"example.com/greetings"
	"fmt"
	"log"
	"net/http"
)

type person struct {
    name string
    age  int
}

func newPerson(name string) *person {
	p:= person{name: name}
	p.age = 42
	return &p
}

type testStruct struct {
	name string
	age  int
	print  func() string
}

func (t testStruct) printName() string {
	fmt.Println("this struct name is:",t.name)
	return "s"
}

func talk() string {
	return "this is the say func saying a name"
}

  // resp, _ := http.Get("http://www.google.com")
  // page, _ := io.ReadAll(resp.Body)

func testGET(url string) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("status code: ", resp.StatusCode)
}

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)


	//t := testStruct{name: "Seamus", age: 10, print: func() string { return "print func Seamus" }}
	
	t := testStruct{name: "Seamus", age: 10, print: talk}
	fmt.Println(t.print())

	// begin copied from https://gobyexample.com/structs
	// fmt.Println(person{"Bob", 20})

	// fmt.Println(person{name: "Alice", age: 30})

	// fmt.Println(person{name: "Fred"})
    
	// fmt.Println(&person{name: "Ann", age: 40})

	// fmt.Println(newPerson("Jon"))

	/*
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name, "who knocks")

	sp := &s
	fmt.Println("This is sp", sp)
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
			"Rex",
			true,
	}
	fmt.Println(dog)

	// messages, err := greetings.Hellos([]string{"gladys", "sammy", "darin"})

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(messages)
	// s := new testStruct() //{age: 10, name: "Seamus"}
	// fmt.Println(s)
	// fmt.Println(*s)
	fmt.Println("how you try me golang")
	*/

	// go say("asdf")
	// say("qer")
}