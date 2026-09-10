package main

import "fmt"

func main() {
	fmt.Println(statusMessage())
}

// statusMessage reports the canary status line.
//
// Provenance: this is the R1 behavior preserved by the verified_revert recovery
// of merge a1378f144bb22bf0ccaeaf513ec069bf67613a65 (failed delivery run
// 591d0b2c-af77-44ef-af59-abce5746758d, post-merge acceptance). The exact
// string is an inline literal in the return statement below: no constant, no
// helper, no other indirection was introduced by the recovery.
func statusMessage() string {
	return "canary: unfiltered [canary-7]"
}
