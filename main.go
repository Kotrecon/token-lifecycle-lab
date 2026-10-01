package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(colorize("=== Token Lifecycle Lab ===", colorCyan))
	fmt.Println()

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
			runScenarioSimulator()

		case ExitMode:
			fmt.Println(colorize("=== Token Lifecycle Lab. The End ===", colorCyan))
			return
		}
	}
}
