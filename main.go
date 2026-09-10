package main

import "fmt"

func main() {
	fmt.Println(statusMessage())
}

// statusMessage reports the canary delivery status line for CANARY-10.
func statusMessage() string {
	return "canary: unfiltered [canary-10]"
}
