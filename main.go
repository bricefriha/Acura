package main

import (
	"fmt"
	"os"
)

func main() {
	var remote string
	framework := "flutter"

	for i := 0; i < len(os.Args); i++ {
		arg := os.Args[i]
		if len(arg) > 0 {
			switch arg {
			case "-r":
				if i+1 < len(os.Args) {
					remote = string(os.Args[i+1])
				}
			case "-f":
				if i+1 < len(os.Args) {
					framework = string(os.Args[i+1])
				}

			}

		}
	}
	fmt.Println("Remote:", remote)
	fmt.Println("Framework:", framework)
}
