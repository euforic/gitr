package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	subCmd := os.Args[1:]

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		if _, err := os.Stat("./" + file.Name() + "/.git"); os.IsNotExist(err) {
			continue
		}

		fmt.Printf("Running git command on: %s\n", file.Name())

		cmd := append([]string{"-C", "./" + file.Name()}, subCmd...)
		out, err := exec.Command("git", cmd...).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", out)
	}
}
