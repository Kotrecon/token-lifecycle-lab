package main

import "fmt"

const (
	colorReset  = "\033[0m"
	colorCyan   = "\033[36m"
	colorYellow = "\033[33m"
	colorOrange = "\033[38;5;208m"
	colorRed    = "\033[31m"
)

func colorize(text string, color string) string {
	return color + text + colorReset
}

func printTokenCard(token Token, stageNumber int, totalStages int) {
	var stateColor = colorReset

	switch token.State {
	case Expired:
		stateColor = colorOrange
	case Revoked:
		stateColor = colorRed
	}

	fmt.Println(colorize(
		fmt.Sprintf("=== Stage %d/%d ===", stageNumber, totalStages),
		colorCyan,
	))

	fmt.Println("Token ID:", token.ID)
	fmt.Println("Token state:", colorize(string(token.State), stateColor))
	fmt.Println("Location:", token.State.Location())
	fmt.Println("Can client use it now:", token.State.ClientUsage())
	fmt.Println("What will API check:", token.State.APIChecks())
	fmt.Println("Next possible event:", token.State.NextEvents())
	fmt.Println("Что происходит:", token.State.Description())
}
