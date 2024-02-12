package main

import (
	"fmt"
	"os"
	"flag"
	"io/ioutil"
    "Reverse"
)

func main() {
	if len(os.Args) != 2 {
        fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
    } else {
		fFlag := flag.String("reverse", "", "Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
		flag.Parse()
		f_path := "../examples/" + *fFlag
        f_content, err := ioutil.ReadFile(f_path)
        if err != nil {
            fmt.Println("Error : Nom de fichier incorrect")
            os.Exit(0)
        }
		toPrint := Reverse.StringReader(Reverse.Format(string(f_content)))
		fmt.Println(toPrint)
	}
}
