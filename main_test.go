package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestOK(t *testing.T) {}

func TestAccept(t *testing.T) {
	// RECOVERY-DRILL SIMULATOR - clearly-labeled clause, not a product assertion.
	// It fails this test when the git HEAD under test has two parents (a merge
	// commit), so the verified-revert / recovery path can be exercised against
	// merge-shaped history. On an ordinary single-parent HEAD it passes silently.
	// Detection consults only the HEAD of the checkout under test (no other
	// branch, tag or ref). If git is unavailable, or the checkout is not a git
	// work tree, the clause is inert and the string assertion below governs.
	//
	// Provenance: this clause and the whole-string assertion below are the R1/R2
	// behavior preserved by the verified_revert recovery of merge
	// a1378f144bb22bf0ccaeaf513ec069bf67613a65 (delivery run
	// 591d0b2c-af77-44ef-af59-abce5746758d). That recovery restores the
	// single-parent HEAD shape which makes this drill silent again; it adds no
	// constant, no helper, no assertion change and no other repository path.
	if out, err := exec.Command("git", "rev-list", "--parents", "-n", "1", "HEAD").Output(); err == nil {
		if fields := strings.Fields(string(out)); len(fields) >= 3 {
			t.Errorf("recovery-drill simulator: HEAD under test is a merge commit with %d parents (%s); failing TestAccept to exercise verified recovery", len(fields)-1, strings.TrimSpace(string(out)))
		}
	}

	if got := statusMessage(); got != "canary: unfiltered [canary-7]" {
		t.Fatalf("statusMessage() = %q, want exactly %q", got, "canary: unfiltered [canary-7]")
	}
}
