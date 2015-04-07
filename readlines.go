package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("--- readlines ---")

	fileName := "foo.deleteme.txt"

	lines := []string{"hi mom", "hi dad", "this is Jon"}
	if err := writeLinesArray(lines, fileName); err != nil {
		log.Fatalf("writeLines: %s\n", err)
	}

	fmt.Println("Reading line by line:")
	c := readLinesChannel(fileName)
	for line := range c {
		fmt.Printf("  Line: %s\n", line)
	}

	fmt.Println("Reading lines into an array:")
	lines, err := readLinesArray(fileName)
	if err != nil {
		log.Fatalf("readLines: %s\n", err)
	}
	for i, line := range lines {
		fmt.Printf("  Line: %d %s\n", i, line)
	}

}

/*
 * Reads a file line by line into an array
 *
 * 	lines, err := readLinesArray(fileName)
 *	if err != nil {
 *		log.Fatalf("readLines: %s\n", err)
 *	}
 *	for i, line := range lines {
 *		fmt.Printf("  Line: %d %s\n", i, line)
 *	}
 */
func readLinesArray(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

/*
 * Writes the lines to the given file.
 */
func writeLinesArray(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

/*
 * Reads a file line by line into a channel
 *
 * c := readLinesChannel(fileName)
 * for line := range c {
 *   fmt.Printf("  Line: %s\n", line)
 * }
 */
func readLinesChannel(fileName string) <-chan string {
	c := make(chan string)
	file, err := os.Open(fileName)
	if err != nil {
		log.Panic(err)
	}
	go func() {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			c <- scanner.Text()
		}
		close(c)
	}()
	return c
}
