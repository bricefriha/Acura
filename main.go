package main

import (
	"fmt"
	"os"
	"os/exec"
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
	cmd := exec.Command("rsync", "-avz",
		"--include="+projName+"/pubspec.yaml",
		"--include="+projName+"/pubspec.lock",
		"--include="+projName+"/lib/***",
		"--include="+projName+"/assets/***",
		"--include="+projName+"/ios/***",
		"--include=lib/config/*",
		"--exclude="+projName+"/*",
		"../",
		fmt.Sprintf("%s:~/Documents/%s", remote, projName),
	)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing rsync:", err)
		return
	}
	fmt.Println("Result:", output)
}
