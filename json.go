package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// https://www.socketloop.com/tutorials/golang-http-response-json-encoded-data
func main() {
	fmt.Println("=== json ===")
	marshal()
	unmarshal()
	encode()
	decode()
}

// marshal converts a JSON string to objects
func marshal() {
	fmt.Println("=== json.marshal ===")
	ryan := &Person{"Ryan", 25}
	wire, err := json.Marshal(ryan)
	check(err)
	fmt.Println(string(wire))
}

// unmarshal converts an object to a JSON string
func unmarshal() {
	fmt.Println("=== json.unmarshal ===")
	var jsonBlob = []byte(`[
		{"name": "Bill", "age": 109},
		{"name": "Bob",  "age": 5}
	]`)

	var persons []Person
	err := json.Unmarshal(jsonBlob, &persons)
	check(err)

	fmt.Printf("%+v\n", persons)
}

// encode writes objects to a JSON *stream*
func encode() {
	fmt.Println("=== json.encode ===")
	f, err := os.Create("deleteme.json")
	check(err)
	enc := json.NewEncoder(f)
	// or enc := json.NewEncoder(os.Stdout)
	p := Person{Name: "Joe", Age: 2}
	err = enc.Encode(&p)
	check(err)
	fmt.Printf("%s: %d\n", p.Name, p.Age)
	os.Remove("deleteme.json")
}

// decode reads a JSON *stream* into objects
func decode() {
	fmt.Println("=== json.decode ===")
	const jsonStream = `
		{"name": "Ed", "age": 55}
		{"name": "Ethan", "age": 33}
		{"name": "Elbert", "age": 111}
	`
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Person
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %d\n", m.Name, m.Age)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
