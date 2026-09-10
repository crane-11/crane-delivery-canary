package main

import "fmt"

// statusCanary is the exact status marker emitted by this build.
//
// The literal value is load-bearing: acceptance testing asserts on it
// byte-for-byte, so it must never be reformatted, trimmed, re-cased or
// otherwise altered. Keeping it in one named constant gives the marker a
// single source of truth instead of an inline literal.
const statusCanary = "canary: unfiltered [canary-7]"

func main() {
	fmt.Println(statusMessage())
}

// statusMessage returns the canary status marker verbatim.
func statusMessage() string {
	return statusCanary
}
