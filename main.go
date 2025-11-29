package main

import (
	flutterService "Acura/services"
	"fmt"
	"os"
)

func main() {
	var remote string
	framework := "flutter"
	argsCount := len(os.Args)
	projName := "my_flutter_project" // fake for testing

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
	flutterService.ExportProject(projName, remote)
}
