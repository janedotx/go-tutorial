package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)



type MapStructureInner struct {
	Name string `mapstructure: name`
	Pet	string `mapstructure: pet`
	ScaryField string `mapstructure: scary_field`

}

type MapStructureOuter struct {
	Category string `mapstructure: category`
	Inner MapStructureInner `mapstructure: inner`
}
/*
 {
   "category": "person",
   "inner": {
     "name": "Alice",
     "pet": "cat"
   }
 }	
*/

func TestMapStructure() {
	var m = make(map[string]interface{})
	m["category"] = "person"
	m["inner"] = map[string]string{"name": "Alice", "pet": "cat"}
	m["scary_field"] = map[string]string{"name": "scary"}

	var result MapStructureOuter
	err := mapstructure.Decode(m, &result)
	if err != nil {
		println("error decoding mapstructure:", err.Error())
		return
	}

	fmt.Printf("%#v", result)

}