package main

import (
	"os/exec"
	"os"
	"log"
)


var cmdChain = []*exec.Cmd{
	exec.Command("./lib/synonyms"),
	exec.Command("./lib/sprinkle",
		`-f=transforms.ini`),
	exec.Command("./lib/coolify"),
	exec.Command("./lib/domainify"),
	exec.Command("./lib/available"),
}

func main() {
	cmdChain[0].Stdin = os.Stdin
	cmdChain[len(cmdChain)-1].Stdout = os.Stdout
	cmdChain[0].Stderr = os.Stderr

	for i := 0; i < len(cmdChain)-1; i++ {
		thisCmd := cmdChain[i]
		nextCmd := cmdChain[i+1]
		stdout, err := thisCmd.StdoutPipe()
		if err != nil {
			log.Fatalln(err)
		}
		nextCmd.Stdin = stdout
		nextCmd.Stderr = os.Stderr
	}

	for _, cmd := range cmdChain {
		if err := cmd.Start(); err != nil {
			log.Fatalln(err)
		} else {
			defer cmd.Process.Kill()
		}
	}

	for _, cmd := range cmdChain {
		if err := cmd.Wait(); err != nil {
			log.Fatalln(err)
		}
	}
}