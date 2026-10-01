package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func printScenarioCategoryMenu() {
	fmt.Println(colorize("Choose a scenario category:", colorCyan))
	fmt.Printf(
		"1. %s\n",
		LifecycleSessionCategory.Label(),
	)
	fmt.Printf(
		"2. %s\n",
		APIValidationCategory.Label(),
	)
	fmt.Printf(
		"3. %s\n",
		LeakageStorageCategory.Label(),
	)
	fmt.Println("0. Back to main menu")
}

func readScenarioCategoryOrBackChoice(
	reader *bufio.Reader,
) (ScenarioCategory, bool, error) {
	fmt.Print(colorize("Choose an option: ", colorYellow))

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, false, err
	}

	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0, false, fmt.Errorf("enter an option number")
	}

	if choice == 0 {
		return 0, true, nil
	}

	switch choice {
	case int(LifecycleSessionCategory):
		return LifecycleSessionCategory, false, nil
	case int(APIValidationCategory):
		return APIValidationCategory, false, nil
	case int(LeakageStorageCategory):
		return LeakageStorageCategory, false, nil
	default:
		return 0, false, fmt.Errorf("choose 0 or a number from 1 to 3")
	}
}
