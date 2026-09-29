package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Issued      TokenState = "Issued"
	Delivered   TokenState = "Delivered"
	Stored      TokenState = "Stored"
	Transmitted TokenState = "Transmitted"
	Used        TokenState = "Used"
	Rotated     TokenState = "Rotated"
	Expired     TokenState = "Expired"
	Revoked     TokenState = "Revoked"
)

var lifecycleStages = []TokenState{
	Issued,
	Delivered,
	Stored,
	Transmitted,
	Used,
	Rotated,
	Expired,
}

func main() {
	token := Token{
		ID:    "demo-token-001",
		State: lifecycleStages[0],
	}

	var reader = bufio.NewReader(os.Stdin)

	fmt.Println(colorize("=== Welcome to Token Lifecycle Lab ===", colorCyan))
	fmt.Println()

	printTokenCard(token, 1, len(lifecycleStages))

	for index := 1; index < len(lifecycleStages); index++ {
		fmt.Println()
		fmt.Println(colorize("Press Enter to continue...", colorYellow))

		_, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(colorize("Input error: "+err.Error(), colorRed))
			return
		}

		token.State = lifecycleStages[index]

		fmt.Println()
		printTokenCard(token, index+1, len(lifecycleStages))
	}

	fmt.Println()
	fmt.Println("Lifecycle completed.")
	fmt.Println("The token expired and can no longer be accepted by the API.")

	fmt.Println(colorize("=== Token Lifecycle Lab. The End ===", colorCyan))
}
