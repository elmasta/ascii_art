package asciiArt

import (
	"strconv"
	"os"
	"os/exec"
)

func TerminalSize() int {
	cmd := exec.Command("stty", "size") //retrieve the size of the terminal
	cmd.Stdin = os.Stdin
  	out, _ := cmd.Output()
	// clean the out to only get the width in int form
	space := false
	temp := ""
	for _, v := range out {
		if space && v != '\n'{
			temp += string(v)
		} else if v == ' ' {
			space = true
		}
	}
	width, _ := strconv.Atoi(temp)
	return width
}
