package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(colorize("=== Welcome to Token Lifecycle Lab ===", colorCyan))
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

			action, err := readActionChoice(reader, actions)
			if err != nil {
				fmt.Println(colorize(err.Error(), colorRed))
				continue
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
				fmt.Println(colorize("=== Token Lifecycle Lab. The End ===", colorCyan))
				return
			}

			fmt.Println()
			break
		}
	}
}
