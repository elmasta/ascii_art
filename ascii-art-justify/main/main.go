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
	cFlag := flag.String("color", "white", "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> something")
	jFlag := flag.String("align", "left", "Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
	flag.Parse()

	// Create a file if we have the output flag
	toFile := false
	f, _ := os.Create(*oFlag)
	if len(string(*oFlag)) > 0 {
		toFile = true
		defer f.Close()
	} else {
		f.Close()
	}
	found := false
    flag.Visit(func(f *flag.Flag) {
        if f.Name == "output" {
            found = true
        } else if f.Name == "color" {
			found = true
		} else if f.Name == "align" {
			found = true
		}
    })

	// check color
	colorFlagPos := 0
	if len(string(*cFlag)) > 0 {
		for i, v := range os.Args {
			if v == "--color="+*cFlag {
				colorFlagPos = i
			}
		}
	}

	// check banner, standard is default
	file := "standard.txt"
	hasFile := false
	switch os.Args[len(os.Args)-1] {
	case "standard":
		hasFile = true
	case "shadow":
		file = "shadow.txt"
		hasFile = true
	case "thinkertoy":
		file = "thinkertoy.txt"
		hasFile = true
	}
	
	if hasFile {
		if found && len(os.Args) > 4 {
			fmt.Println("Usage: go run .  [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right  something  standard")
			os.Exit(0)
		}
	} else {
		if found && len(os.Args) > 3 {
			fmt.Println("Usage: go run .  [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right  something  standard")
			os.Exit(0)
		}
	}
	
	ind := -1 // ind is the index of the starting part we have to color
	count := -1 // count is the number of letters we have to color
	word := "" // init the word we are working with
	startInd := 0
	if hasFile {
		word = os.Args[len(os.Args)-2]
		count = len(os.Args[len(os.Args)-2])
	} else {
		word = os.Args[len(os.Args)-1]
		count = len(os.Args[len(os.Args)-1])
	}
	if os.Args[colorFlagPos+1] != string(*oFlag) {
		count = len(os.Args[colorFlagPos+1])
		ind = strings.Index(word, os.Args[colorFlagPos+1])
	} else {
		ind = 0
	}
	startInd = ind
	
	// We retrieve the banner
	byte, _ := ioutil.ReadFile(file)
	data := strings.Split(string(byte), "\n")

	// We retrieve the size of the terminal
	width := asciiArt.TerminalSize()
	sList, lineSize, rest := asciiArt.PrepJustify(word, data, width, file)

	// We launch the printer
	asciiArt.StrTreatment(string(*jFlag), string(*cFlag), sList, lineSize, toFile, f, width, startInd, ind, count, word, data, rest)
}
