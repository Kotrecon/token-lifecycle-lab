package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

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

func printTokenCard(token Token) {
	var stateColor = colorReset

	switch token.State {
	case Expired:
		stateColor = colorOrange
	case Revoked:
		stateColor = colorRed
	}

	fmt.Println(colorize("=== Token Card ===", colorCyan))

	fmt.Println("Token ID:", token.ID)
	fmt.Println("Token state:", colorize(string(token.State), stateColor))
	fmt.Println("Location:", token.State.Location())
	fmt.Println("Can client use it now:", token.State.ClientUsage())
	fmt.Println("What will API check:", token.State.APIChecks())
	fmt.Println("Next possible event:", token.State.NextEvents())
	fmt.Println("Что происходит:", token.State.Description())
}

func printActionsMenu(actions []LifecycleAction) {
	if len(actions) == 0 {
		return
	}

	fmt.Println()
	fmt.Println(colorize("Available actions:", colorCyan))

	for index, action := range actions {
		fmt.Printf("%d. %s\n", index+1, action.Label())
	}
}

func readActionChoice(
	reader *bufio.Reader,
	actions []LifecycleAction,
) (LifecycleAction, error) {
	fmt.Print(colorize("Choose an action: ", colorYellow))

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return "", fmt.Errorf("enter an action number")
	}

	if choice < 1 || choice > len(actions) {
		return "", fmt.Errorf(
			"choose a number from 1 to %d",
			len(actions),
		)
	}

	return actions[choice-1], nil
}

func readRestartChoice(reader *bufio.Reader) (bool, error) {
	fmt.Println()
	fmt.Println("What would you like to do?")
	fmt.Println("1. Start a new demo run")
	fmt.Println("2. Exit")
	fmt.Print(colorize("Choose an option: ", colorYellow))

	input, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return false, fmt.Errorf("enter 1 to start a new demo run or 2 to exit")
	}

	switch choice {
	case 1:
		return true, nil
	case 2:
		return false, nil
	default:
		return false, fmt.Errorf("choose 1 to start a new demo run or 2 to exit")
	}
}
