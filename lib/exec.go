package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	command := exec.Command("ls")
	var out bytes.Buffer
	command.Stdout = &out
	err := command.Run()

	if err != nil {

		log.Println(err)
	}

	fmt.Println(out.String())
}
