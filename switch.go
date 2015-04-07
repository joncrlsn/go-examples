package main

import "fmt"
import "os"
import "log"

var port int

func main() {
	fmt.Println("Hello!")
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "chooseone":
			fmt.Println("You chose one")
		case "concat":
			fmt.Println("You chose concat")
		case "diff":
			fmt.Println("You chose diff")
		default:
			fmt.Println("Invalid argument:", os.Args[1])
		}
	}
	fmt.Println("Done!")
}

func check(err error, action string) {
	if err != nil {
		log.Fatalf("Error %s: %v\n", action, err)
	}
}
