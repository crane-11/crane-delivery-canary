package main

import "fmt"

func main() {
	fmt.Println(statusMessage())
}

func statusMessage() string {
	return "canary: unfiltered [canary-10b]"
}

func statusMessageFiltered() string {
	return "canary: filtered [canary-19]"
}
