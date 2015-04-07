//
// Compare two string slices, outputting commands to make the second match the first.
// This design is in anticipation of the pgdiff program I am writing.
//
package main

import "fmt"

func main() {
	fmt.Println("--- diff ---")
	fmt.Println(" (Make slice 2 the same as slice 1)")

	s1 := []string{"a", "b", "c", "d", "e", "f", "g"}
	s2 := []string{"a", "a1", "b", "d"}
	doDiff(s1, s2)
	doDiff(s2, s1)

	s1 = []string{"b", "c", "d", "e", "f"}
	s2 = []string{"a", "b", "g"}
	doDiff(s1, s2)
	doDiff(s2, s1)
}

func doDiff(slice1 []string, slice2 []string) {
	fmt.Printf("------\n== slice1: %v\n", slice1)
	fmt.Printf("== slice2: %v\n", slice2)
	i1 := 0
	i2 := 0
	v1 := slice1[i1]
	v2 := slice2[i2]
	for i1 < len(slice1) || i2 < len(slice2) {
		if v1 == v2 {
			// compareValues(v1, v2) (look for non-identifying changes)
			fmt.Printf("Both have %s\n", v1) // add v1
			i1 = i1 + 1
			if i1 < len(slice1) {
				v1 = slice1[i1] // nextValue(db1)
			}
			i2 = i2 + 1
			if i2 < len(slice2) {
				v2 = slice2[i2] // nextValue(db2)
			}
		} else if v1 < v2 {
			// slice2 is missing a value slice1 has
			i1 = i1 + 1
			if i1 < len(slice1) {
				fmt.Printf("Add %s to slice 2\n", v1) // add v1
				v1 = slice1[i1]                       // nextValue(db1)
			} else {
				// slice 2 is longer than slice1
				fmt.Printf("Drop %s from slice 2\n", v2) // add v1
				i2 = i2 + 1
				if i2 < len(slice2) {
					v2 = slice2[i2] // nextValue(db1)
				}
			}
		} else if v1 > v2 {
			// db2 has an extra column that we don't want
			i2 = i2 + 1
			if i2 < len(slice2) {
				fmt.Printf("Drop %s from slice 2\n", v2) // drop v2
				v2 = slice2[i2]                          // nextValue(db2)
			} else {
				// slice 1 is longer than slice2
				fmt.Printf("Add %s to slice 2\n", v1) // add v1
				i1 = i1 + 1
				if i1 < len(slice1) {
					v1 = slice1[i1] // nextValue(db1)
				}
			}
		}
	}
}
