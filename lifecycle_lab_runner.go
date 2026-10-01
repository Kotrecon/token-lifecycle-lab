package main

import (
	"bufio"
	"fmt"
)

func runLifecycleLab(reader *bufio.Reader) {
	fmt.Println(colorize("=== Interactive Lifecycle Lab ===", colorCyan))
	fmt.Println()

	for {
		token := Token{
			ID:    "demo-token-001",
			State: Issued,
		}

		for {
			printTokenCard(token)

			actions := AvailableActions(token.State)
			if len(actions) == 0 {
				fmt.Println(colorize("Token reached a terminal state.", colorCyan))
				break
			}

			printActionsMenu(actions)

			action, backToMainMenu, err := readActionOrExitChoice(reader, actions)
			if err != nil {
				fmt.Println(colorize(err.Error(), colorRed))
				continue
			}

			if backToMainMenu {
				fmt.Println()
				fmt.Println(
					colorize("=== Leaving Interactive Lifecycle Lab ===", colorCyan),
				)
				return
			}

			err = ApplyAction(&token, action)
			if err != nil {
				fmt.Println(colorize(err.Error(), colorRed))
				continue
			}

			fmt.Println()
		}

		for {
			startNewRun, err := readRestartChoice(reader)
			if err != nil {
				fmt.Println(colorize(err.Error(), colorRed))
				continue
			}

			if !startNewRun {
				fmt.Println()
				fmt.Println(
					colorize("=== Leaving Interactive Lifecycle Lab ===", colorCyan),
				)
				return
			}

			fmt.Println()
			break
		}
	}
}
