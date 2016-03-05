package main

import (
	"fmt"
	"encoding/json"
)

type NameStruct struct {
	Name string `json:"uuid"`
	BrokenName string `json:"brokenname"`
}

func (n NameStruct) brokenname() string {
	return n.Name + "blahlblahlb"
}

func (n NameStruct) MarshalJSON() ([]byte, error) {
	type Alias NameStruct
	return json.Marshal(&struct {
		BrokenName string `json:"brokenname"`
		Alias
	}{
		BrokenName:     n.brokenname(),
		Alias:    (Alias)(n),
	})
}

func main() {
	a := NameStruct{ Name: "john" }
	// create a new variable that implements a Marshaler, then it should have MarshalJSON. This helped me find that it should not be a function on a pointer struct, but just a function on a struct
	// var x json.Marshaler
	// x = a
	// run it directly
	// b, _ := a.MarshalJSON()
	b, _ := json.Marshal(a)
	output := string(b)
	fmt.Println(output)
}
