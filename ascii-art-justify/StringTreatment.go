package asciiArt

import (
	"fmt"
	"os"
)

func PrintStr(strTable []string, toFile bool, f *os.File) { // print to the console or add to file
	for _, v := range strTable {
		if toFile {
			f.WriteString(v + "\n")
		} else {
			fmt.Println(v)
		}
	}
}

func Color(v string, flag string) string { // add ANSI escape code to a set of characters
	colors := map[string]string {
		"reset": "\033[0m",
		"red":   "\033[31m",
		"green": "\033[32m",
		"yellow": "\033[33m",
		"blue": "\033[34m",
		"purple": "\033[35m",
		"cyan": "\033[36m",
		"white": "\033[97m",
	}
	return  colors[flag] + v  + colors["reset"]
}

func AddToTable(v byte, strTable []string, data []string, toChange bool, flag string) []string {
	count := 0
	first := false
	if len(strTable) == 0 {
		first = true
	}
	for count < 8 {
		/*
			Each character is 8 lines.
			We loop 8 times to retrieve the parts of the character.
			The formula to retrieve the part of the character in the file is:

			value ascii of the character - 32 (the first printable character is space at 32) * 9 + counter + 1.

			We add what we retrieve in the string table in each string.
			We return the string table once we're done.
		*/
		temp := ""
		for _, v := range data[(int(v)-32)*9+count+1] {
			if v != 13 {
				temp += string(v)
			}
		}
		if toChange {
			temp = Color(temp, flag)
		}
		if first {
			strTable = append(strTable, temp)
		} else {
			strTable[count] += temp
		}
		count++
	}
	return strTable
}

func StrTreatment(jFlag string, cFlag string, sList [][]string, lineSize []int, toFile bool, f *os.File, width int, startInd int, ind int, count int, word string, data []string, rest []int) {
	strTable := []string{} // We prepare an empty table of string that will be used to store the strings to print
	tempI := 0
	if jFlag == "justify" {
		for i, v := range sList {
			for subI, subV := range v {
				for subSubI := range subV {
					strTable = AddToTable(subV[subSubI], strTable, data, false, "white")
				}
				if subI+1 < len(v) {
					count := 0
					for count < 8 {
						tempS := lineSize[i]-1
						for tempS >= 0 {
							strTable[count] += " "
							tempS--
						}
						if rest[tempI] > 0 {
							strTable[count] += " "
						}
						count++
					}
					if rest[tempI] > 0 {
						rest[tempI]--
					}
				}
			}
			if len(v[0]) > 0 {
				tempI++
			}
			PrintStr(strTable, toFile, f)
			strTable = []string{}
		}
	} else {
		for i := 0; i < len(word); i++ { // We loop for each character in the argument
			v := word[i]
			if v == '\\' && len(word)-1 != i { // If we find a \ & the index is not on the last character
				if word[i+1] == 'n' { // If the next character is an n
					if len(strTable) > 0 {
						nmbOfSpace := width - len(strTable[0])
						switch jFlag {
						case "center":
							nmbOfSpace = nmbOfSpace/2
							for i := range strTable {
								for count := nmbOfSpace; count >= 0; count-- {
									strTable[i] = " " + strTable[i]
								}
							}
						case "right":
							for i := range strTable {
								for count := nmbOfSpace-1; count >= 0; count-- {
									strTable[i] = " " + strTable[i]
								}
							}
						}
						PrintStr(strTable, toFile, f) // If we find a \n and the string table has something ot be printed
					} else if toFile {
						f.WriteString("\n") // If the string table is empty and we wright to a file
					} else {
						fmt.Println() // Otherwirse we just print an empty line
					}
					// We reset the table of string
					strTable = []string{}
					i++
					if startInd == 0 {
						ind += 2
					}
				} else {
					// We call the function addToTable to add each part of a word in the table of string
					if ind == i && count > 0 {
						strTable = AddToTable(v, strTable, data, true, cFlag)
						ind++
						count--
					} else {
						strTable = AddToTable(v, strTable, data, false, cFlag)
					}
				}
			} else {
				// We call the function addToTable to add each part of a word in the table of string
				if ind == i && count > 0 {
					strTable = AddToTable(v, strTable, data, true, cFlag)
					ind++
					count--
				} else {
					strTable = AddToTable(v, strTable, data, false, cFlag)
				}
			}
		}
		if len(strTable[0]) > 0 {
			// if we're at the end of the word and there's something to print we print it
			nmbOfSpace := width - len(strTable[0])
			switch jFlag {
			case "center":
				nmbOfSpace = nmbOfSpace/2
				for i := range strTable {
					for count := nmbOfSpace; count >= 0; count-- {
						strTable[i] = " " + strTable[i]
					}
				}
			case "right":
				for i := range strTable {
					for count := nmbOfSpace-1; count >= 0; count-- {
						strTable[i] = " " + strTable[i]
					}
				}
			}
			PrintStr(strTable, toFile, f)
		}
	}
}
