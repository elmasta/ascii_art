package main

// welcome to the realm of magic

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"asciiArt"
)

func main() {
	// If we have no args we exit
	if len(os.Args) < 2 {
		fmt.Println("No argument detected, please open readme.md to know how to use this program")
		os.Exit(0)
	}

	// Retrieve the flags
	oFlag := flag.String("output", "", "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	flag.Parse()

	// Create a file if we have the output flag
	toFile := false
	fToCreate := "../examples/" + *oFlag
	f, _ := os.Create(fToCreate)
	if len(string(*oFlag)) > 0 {
		toFile = true
		defer f.Close()
	} else {
		f.Close()
	}

	// check banner, standard is default
	file := "../standard.txt"
	word := ""
	if len(os.Args) == 2 {
		word = os.Args[1]
	} else {
		word = os.Args[len(os.Args)-1]
	}
	
	// We retrieve the banner
	bytes, _ := ioutil.ReadFile(file)
	data := strings.Split(string(bytes), "\n")

	// We launch the printer
	asciiArt.StrTreatment(toFile, f, word, data)
}
