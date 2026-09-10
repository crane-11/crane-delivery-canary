package main

import "fmt"

func main() {
	fmt.Println(statusMessage())
}

// statusMessage reports the canary delivery status line for CANARY-10.
// Per requirement R1 the message is produced as an inline string literal in
// the return statement itself: no package-level constants and no helper
// functions are involved in building it.
func statusMessage() string {
	return "canary: unfiltered [canary-10]"
}
