package main

import "os"

type NullFieldsTest struct {
	name string
} 

func TestNullFields() {
	nullTest:= NullFieldsTest{ name: "asdf"}

	nullTest.name = "qwer"
	println("this is the name field of nullTest:", nullTest.name)
}

type Bacon struct {
	name string
	attributes []NullFieldsTest
}

type BaconAnyAttrs struct {
	name string
	attributes []interface{}
}

func TestAssigningStructs2() {
	d := map[string]string{"name": "crispy"}
	c := []interface{}{d}
	b := BaconAnyAttrs{name: "bacon", attributes: c}

	b.attributes = []interface{}{NullFieldsTest{name: "crispy"}}
	// b2 := Bacon{name: "bacon", attributes: []NullFieldsTest{} }
	// b2.attributes = b.attributes.([]NullFieldsTest)
	var e interface{} = NullFieldsTest{name: "crispy"}
	n, ok := e.(NullFieldsTest)
	if !ok {
		println("type assertion failed")
		os.Exit(1)
	}
	println("this is the name field of n:", n.name)
}

func TestAssigningStructs1() {
	b := Bacon{name: "bacon", attributes: []NullFieldsTest{{name: "crispy"}, {name: "salty"}}}

	println("this is the name field of b:", b.name)

	var m1 = make(map[string]string)
	m1["name"] = "crispy"
	var m2 = make(map[string]string)
	m2["name"] = "salty"

	// var a = [2]map[string]string{m1, m2}

	var n1 = NullFieldsTest{name: "n1"}
	var narr = [1]NullFieldsTest{n1}

	// [:] is used to convert an array to a slice
	c := Bacon{name: "bacon", attributes: narr[:]}

	println("this is the name field of c:", c.name)
}