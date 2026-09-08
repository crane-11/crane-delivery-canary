package main

import "testing"

func TestOK(t *testing.T) {}

func TestAccept(t *testing.T) {
	if statusMessage() == "" {
		t.Fatal("empty status message")
	}
}
