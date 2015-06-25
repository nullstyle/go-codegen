package main

import (
	"log"
	"os"

	codegen "github.com/nullstyle/go-codegen"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"."}
	}

	for _, arg := range args {
		err := codegen.Process(arg)

		if err != nil {
			log.Fatalln(err)
		}
	}
}
