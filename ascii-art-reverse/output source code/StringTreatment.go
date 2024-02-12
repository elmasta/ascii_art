package asciiArt

import (
	"fmt"
	"os"
)

func PrintStr(strTable []string, toFile bool, f *os.File) { // print to the console or add to file
	for i, v := range strTable {
		if toFile {
			if i == len(strTable) -1 {
				f.WriteString(v + "\n$")
			} else {
				f.WriteString(v + "\n")
			}
		} else {
			fmt.Println(v)
		}
	}
}

func AddToTable(v byte, strTable []string, data []string) []string {
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
		for _, subV := range data[(int(v)-32)*9+count+1] {
			if subV != 13 {
				temp += string(subV)
			}
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

func StrTreatment(toFile bool, f *os.File, word string, data []string) {
	strTable := []string{} // We prepare an empty table of string that will be used to store the strings to print
	for i := 0; i < len(word); i++ { // We loop for each character in the argument
		v := word[i]
		if v == '\\' && len(word)-1 != i { // If we find a \ & the index is not on the last character
			if word[i+1] == 'n' { // If the next character is an n
				if len(strTable) > 0 {
					PrintStr(strTable, toFile, f) // If we find a \n and the string table has something ot be printed
				} else if toFile {
					f.WriteString("\n") // If the string table is empty and we wright to a file
				} else {
					fmt.Println() // Otherwirse we just print an empty line
				}
				// We reset the table of string
				strTable = []string{}
				i++
			} else {
				// We call the function addToTable to add each part of a word in the table of string
				strTable = AddToTable(v, strTable, data)
			}
		} else {
			// We call the function addToTable to add each part of a word in the table of string
			strTable = AddToTable(v, strTable, data)
		}
		
	}
	//add the $ at the end of lines
	for i, v := range strTable {
		strTable[i] = v + "$"
	}
	if len(strTable[0]) > 0 {
		// if we're at the end of the word and there's something to print we print it
		PrintStr(strTable, toFile, f)
	}
}
