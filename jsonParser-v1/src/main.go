package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Job     string  `json:"job"`
	Address Address `json:"address"`
}

func parseJSON(this js.Value, p []js.Value) interface{} {
	jsonString := p[0].String()

	var person Person
	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		return js.ValueOf(fmt.Sprintf("Error parsing JSON: %v", err))
	}

	result := fmt.Sprintf(
		"Name: %s\nAge: %d\nJob: %s\nAddress:\n  Street: %s\n  City: %s\n  State: %s",
		person.Name, person.Age, person.Job, person.Address.Street, person.Address.City, person.Address.State,
	)

	return js.ValueOf(result)
}

func main() {
	fmt.Println("webAssembly with GoLang ")
	c := make(chan struct{}, 0)
	js.Global().Set("parseJSON", js.FuncOf(parseJSON))
	<-c
}
