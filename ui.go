package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const (
	colorReset      = "\033[0m"
	colorCyan       = "\033[36m"
	colorBrightCyan = "\033[96m"
	colorYellow     = "\033[33m"
	colorOrange     = "\033[38;5;208m"
	colorRed        = "\033[31m"
	colorGray       = "\033[90m"
)

func colorize(text string, color string) string {
	return color + text + colorReset
}

func printMainMenu() {
	fmt.Println(colorize("Main menu:", colorCyan))
	fmt.Println("1. Interactive Lifecycle Lab")
	fmt.Println("2. Scenario Simulator")
	fmt.Println("3. Exit")
}

func readAppMode(reader *bufio.Reader) (AppMode, error) {
	fmt.Print(colorize("Choose an option: ", colorYellow))

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0, fmt.Errorf("enter an option number")
	}

	switch choice {
	case int(LifecycleLabMode):
		return LifecycleLabMode, nil
	case int(ScenarioSimulatorMode):
		return ScenarioSimulatorMode, nil
	case int(ExitMode):
		return ExitMode, nil
	default:
		return 0, fmt.Errorf("choose a number from 1 to 3")
	}
}

func printApplicationWelcome() {
	fmt.Println(colorize("=== Token Lifecycle Lab ===", colorCyan))
	fmt.Println()
}

func printApplicationEnd() {
	fmt.Println(colorize("=== Token Lifecycle Lab. The End ===", colorCyan))
}
