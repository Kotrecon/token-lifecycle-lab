package main

import (
	"bufio"
	"fmt"
)

func runScenarioSimulator(reader *bufio.Reader) {
	fmt.Println(colorize("=== Scenario Simulator ===", colorCyan))
	fmt.Println()

	for {
		printScenarioCategoryMenu()

		category, backToMainMenu, err :=
			readScenarioCategoryOrBackChoice(reader)
		if err != nil {
			fmt.Println(colorize(err.Error(), colorRed))
			fmt.Println()
			continue
		}

		if backToMainMenu {
			fmt.Println()
			fmt.Println(
				colorize("=== Leaving Scenario Simulator ===", colorCyan),
			)
			fmt.Println()
			return
		}

		fmt.Println()
		fmt.Printf("%s is being prepared.\n", category.Label())
		fmt.Println()
	}
}
