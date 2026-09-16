package main

import "fmt"

import "strings"

func wordCount(s string) int {
	return len(strings.Fields(s))
}

import "strings"

func wordCount(s string) int {
	return len(strings.Fields(s))
}

func main() {
	fmt.Println(statusMessage())
}

func statusMessage() string {
	return "canary: unfiltered [canary-10b]"
}

func statusMessageFiltered() string {
	return "canary: filtered [canary-19]"
}
