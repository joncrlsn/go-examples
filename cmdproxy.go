package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// main runs the bc Linux command and proxies stdin input to it... all in one method
func main() {
	cmd := exec.Command("pianobar")

	in, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	out, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	//
	// Capture standard error and print it
	//
	go func() {
		defer stderr.Close()
		errReader := bufio.NewReader(stderr)
		errScanner := bufio.NewScanner(errReader)
		for errScanner.Scan() {
			fmt.Println(errScanner.Text())
		}
	}()

	//
	// Capture standard input and pass it to the command
	//
	go func() {
		defer in.Close()
		consolereader := bufio.NewReader(os.Stdin)

		for {
			//fmt.Print("> ")
			//			inputText, err := consolereader.ReadString('\n') // this will wait for user input
			b, err := consolereader.ReadByte() // this will wait for user input
			if err != nil {
				if err != io.EOF {
					panic(err)
				}
			} else {
				_, err := in.Write([]byte{b})
				if err != nil {
					panic(err)
				}
				//fmt.Fprintln(in, inputText)
			}
		}
	}()

	//
	// Start the process
	//
	if err = cmd.Start(); err != nil {
		panic(err)
	}

	//
	// Capture standard output and print it
	//
	go func() {
		defer out.Close()
		reader := bufio.NewReader(out)
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	fmt.Println("Enter bc calculations (like 5*4) and press enter.  Enter 'quit' to exit")
	cmd.Wait()

}
