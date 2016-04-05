//
// An example of how to process a csv files
//
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== csv ===")
	ReadData()
	WriteArray()
}

func WriteArray() {
	fmt.Println("Write: ===")
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func ReadData() {
	fmt.Println("Read: ===")

	// The first line is the title line
	const in = `
"Group","Title","Username","Password","URL","Notes"
"Root","Peterson Family Tree","joncrlsn","more-secret-than-you-know","http://www.example.com/phpgedview/",""
"Root","Twitter","jon_carlson","supersecret","http://twitter.com",""
`
	r := csv.NewReader(strings.NewReader(in))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}
