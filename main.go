package main

import "fmt"

func main() {
	fmt.Println(statusMessage())
}

// statusMessage reports the canary delivery status line for CANARY-10.
//
// Requirement R1 pins the exact value "canary: unfiltered [canary-10]" and
// requires it to appear as an inline string literal inside this return
// statement, so no package-level constant and no helper function participates
// in building the message. The companion assertion in main_test.go
// (TestAccept) compares against the same literal with whole-string equality.
func statusMessage() string {
	return "canary: unfiltered [canary-10]"
}
