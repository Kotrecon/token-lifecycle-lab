package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	printApplicationWelcome()

	for {
		printMainMenu()

		mode, err := readAppMode(reader)
		if err != nil {
			fmt.Println(colorize(err.Error(), colorRed))
			fmt.Println()
			continue
		}

		fmt.Println()

		switch mode {
		case LifecycleLabMode:
			runLifecycleLab(reader)

		case ScenarioSimulatorMode:
			runScenarioSimulator(reader)

		case ExitMode:
			printApplicationEnd()
			return
		}
	}
}
