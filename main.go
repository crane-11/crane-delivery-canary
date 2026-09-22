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

const DrillOverlayProbe = "overlay-probe"

func DrillOverlayProbeLabel() string {
	return DrillOverlayProbe
}

const DrillVersionMarker = "canary-drill-marker-v1"

func DrillVersionLabel() string {
	return "version: " + DrillVersionMarker
}

const DrillSweepMarker = "canary-sweep-marker-v2"

func DrillSweepLabel() string {
	return "sweep: " + DrillSweepMarker
}

func longer(a, b string) string {
	if len(a) >= len(b) {
		return a
	}
	return b
}
