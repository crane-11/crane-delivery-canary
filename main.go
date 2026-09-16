package main

import "fmt"

func main() {
	fmt.Println(statusMessage() + " drill-scenario-tripwire")
}

const StatusDrill = "drill-scenario-a"

func StatusDrillLabel() string {
	return statusMessage() + " " + StatusDrill
}

func statusMessage() string {
	return "canary: unfiltered [canary-10b]"
}

func statusMessageFiltered() string {
	return "canary: filtered [canary-19]"
}

func BaseLabel() string {
	return "canary: unfiltered [drill-boundary]"
}
