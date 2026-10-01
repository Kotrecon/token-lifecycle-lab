package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func printKeyValue(key string, value string) {
	fmt.Printf("%s %s\n", colorize(key+":", colorYellow), value)
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

	printKeyValue("Token ID", token.ID)
	printKeyValue(
		"Token state",
		colorize(string(token.State), stateColor),
	)
	printKeyValue("Location", token.State.Location())
	printKeyValue("Can client use it now", token.State.ClientUsage())
	printKeyValue("What will API check", token.State.APIChecks())
	printKeyValue("Next possible event", token.State.NextEvents())
	printKeyValue("Что происходит", token.State.Description())
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

	fmt.Println("0. Back to main menu")
}

func readActionOrExitChoice(
	reader *bufio.Reader,
	actions []LifecycleAction,
) (LifecycleAction, bool, error) {
	fmt.Print(colorize("Choose an action: ", colorYellow))

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", false, err
	}

	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return "", false, fmt.Errorf("enter an action number")
	}

	if choice == 0 {
		return "", true, nil
	}

	if choice < 1 || choice > len(actions) {
		return "", false, fmt.Errorf(
			"choose 0 to return to main menu or a number from 1 to %d",
			len(actions),
		)
	}

	return actions[choice-1], false, nil
}

func readRestartChoice(reader *bufio.Reader) (bool, error) {
	fmt.Println()
	fmt.Println("What would you like to do?")
	fmt.Println("1. Start a new demo run")
	fmt.Println("2. Back to main menu")
	fmt.Print(colorize("Choose an option: ", colorYellow))

	input, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return false, fmt.Errorf(
			"enter 1 to start a new demo run or 2 to return to main menu",
		)
	}

	switch choice {
	case 1:
		return true, nil
	case 2:
		return false, nil
	default:
		return false, fmt.Errorf(
			"choose 1 to start a new demo run or 2 to return to main menu",
		)
	}
}
