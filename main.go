package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"log"
)

const (
	Serve = "serve"
	Build = "build"
)

func Usage() {
	fmt.Printf("Options: \n")
	fmt.Printf("\t%s [PORT]	Creates a web server at the given port.\n", Serve)
	fmt.Printf("\t%s [PATH]	Builds the static website.", Build)
}

func StartWebServer(port int) {
	fmt.Printf("Starting web server at address 127.0.0.1:%d", port)
}

func GenerateStaticSite(path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("ERROR: Could not list the content of directory `%s`!\n", path)
	}
	for _, file := range entries {
		extension := strings.Split(file.Name(), ".")[1]
		if extension == "html" {
			completePath := path + "/" + file.Name()
			bytes, err := os.ReadFile(completePath)
			if err != nil {
				log.Fatalf("ERROR: could not read file `%s`!\n", file.Name())
			}
			fileContent := string(bytes[:])
			for _, char := range fileContent {
				fmt.Printf("%c", char)
			}
		}
	}
}

func main() {
	program := os.Args[0]
	
	if len(os.Args) < 2 {
		fmt.Printf("Usage: <%s> [OPTIONS]\n", program)
		Usage()
		os.Exit(1)
	}
	
	if os.Args[1] == Serve {
		if len(os.Args) < 3 {
			log.Fatalf("ERROR: `%s` command requires a port!\n", Serve)
		}
		port, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("ERROR: Could not parse `%s` as a port!\n", os.Args[2])
		}
		StartWebServer(port)
	}

	if os.Args[1] == Build {
		if len(os.Args) < 3 {
			log.Fatalf("ERROR: `%s` command requires a path!\n", Build)
		}
		GenerateStaticSite(os.Args[2])
	}
}