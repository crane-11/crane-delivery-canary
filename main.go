package main

import "fmt"

func main() {
	fmt.Println(statusMessage())
}

func statusMessage() string {
	return "canary: unfiltered [canary-6]"
}
