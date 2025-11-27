package main

import (
	"fmt"
	"os"
)

func main() {
	var remote string
	framework := "flutter"
	argsCount := len(os.Args)

	for i := 0; i < argsCount; i++ {
		arg := os.Args[i]

		if len(arg) > 0 {
			switch arg {
			case "-r":
				if i+1 < argsCount {
					remote = string(os.Args[i+1])
				}
			case "-f":
				if i+1 < argsCount {
					framework = string(os.Args[i+1])
				}

			}

		}
	}
	fmt.Println("Remote:", remote)
	fmt.Println("Framework:", framework)
}
