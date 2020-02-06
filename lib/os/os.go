package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Println(os.Args[0], os.Args[1])
	command := exec.Command("./lib/os/run.sh")

	var out bytes.Buffer
	command.Stdout = &out
	err := command.Run()

	if err != nil {
		log.Println(err)
	}

	fmt.Println(out.String())
}
