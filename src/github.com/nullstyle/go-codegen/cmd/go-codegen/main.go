package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	codegen "github.com/nullstyle/go-codegen"
)

var includes = flag.String(
	"include",
	".",
	"A colon separated list of directories to search for templates",
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	searchPath := []string{"."}
	if *includes != "" {
		for _, path := range strings.Split(*includes, ":") {
			path, err := filepath.Abs(path)
			if err != nil {
				log.Fatal(err)
			}

			stat, err := os.Stat(path)
			if err != nil {
				log.Fatal(err)
			}

			if !stat.IsDir() {
				log.Fatalf("not a directory: %s", path)
			}

			searchPath = append(searchPath, path)
		}
	}

	for _, arg := range args {
		arg, err := filepath.Abs(arg)
		if err != nil {
			log.Fatal(err)
		}

		err = codegen.Process(arg, searchPath)

		if err != nil {
			log.Fatalln(err)
		}
	}
}
