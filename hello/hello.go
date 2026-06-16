package main

import (
	//	"example.com/greetings"
	"encoding/json"
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

type testMarshalStruct struct {
	Name string
	name string
}

type testStruct struct {
	name  string
	age   int
	print func() string
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

/*
  Funny thing about this: if you have a struct with two fields that are the same except 
	for capitalization, the json package will only unmarshal into the capitalized one. 
	So in this case, testMarshal.Name will be "capital" and testMarshal.name will be "" (empty string). 
	This is because the json package only looks at exported fields 
	(those that start with a capital letter) when unmarshaling JSON data.
*/
func testMarshal() {
	var testMarshal testMarshalStruct
	err := json.Unmarshal([]byte("{\"name\": \"asdfname\",\"Name\":\"capital\"}"), &testMarshal)
	if err != nil {
		fmt.Println("failed to unmarshal")
		log.Fatalln(err)
	}
	fmt.Println(testMarshal.Name)
	fmt.Println(testMarshal.name)
}

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

//	TestMapStructure()
	setupDb()

	writeDevice("device1", "http://example.com/device1")
	devices := readDevices()

	fmt.Printf("\n%s\n", devices[0].URL)
	fmt.Printf("%s\n", devices[0].ID)
	fmt.Printf("%s\n", devices[0].Organization)
}