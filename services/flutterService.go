package flutterService

import (
	"fmt"
	"os/exec"
)

func ExportProject(projName string, remote string, dir string) {
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
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing rsync:", err)
		return
	}
	fmt.Println("Export successfull:", output)

}
