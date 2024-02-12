package Reverse

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// renvoi l'ascii-art sous forme d'un tableau de 8 strings
func Format(data string) []string {
	toReturn := strings.Split(data, "$\n")
	toReturn = toReturn[:8]
	return toReturn
}

// converti l'ascii-art en son équivalent en string
func StringReader(data []string) string {
	toReturn := ""
	for count := 0; count < len(data[0]); count++ { // parcours les colonnes de l'ascii-art
		if data[0][count] == ' ' {
			nmbOfSpace := 0
			for i := 0; i < 8; i++ { // vérifie si toute la colonne est composée d'espace
				if data[i][count] == ' ' {
					nmbOfSpace++
				}
			}
			if nmbOfSpace == 8 {
				if count == 0 { // repère l'espace
					toReturn += " "
					for i := 0; i < 8; i++ { // supprime l'espace
						data[i] = data[i][5:]
					}
				} else {
					temp := Compare(data, count) // Compare la lettre avec les lettres du fichier standard
					if temp == "Letter Not Found" {
						return "Error : Lettre non reconnue"
					}
					toReturn += temp
				}
				for i := 0; i < 8; i++ { // supprime la lettre récupérée de l'ascii-art
					data[i] = data[i][count+1:]
				}
				count = -1
			}
		}
	}
	return toReturn
}

func Compare(data []string, count int) string {
	var cleanData []string
	for _, line := range data { // Extrait juste la lettre de l'ascii-art à comparer
		cleanData = append(cleanData, line[:count+1])
	}
	f_content, err := ioutil.ReadFile("../standard.txt")
	if err != nil {
		fmt.Println("Error : standard.txt not found.")
		os.Exit(0)
	}
	lines := strings.Split(string(f_content), "\n")
	subCount := 0
	nmbOfAscii := 0
	for i, line := range lines { // repère le numéro de la lettre dans le fichier standard
		if len(line) == 0 {
			nmbOfAscii++
		}
		if line == cleanData[0] { // la première ligne des deux lettres
			for j, data_line := range cleanData {
				if data_line == lines[i+j] { // compare les 8 lignes des deux lettres
					subCount++
				}
			}
			if subCount == 8 {
				return string(rune(nmbOfAscii - 1 + 32))
			} else {
				subCount = 0
			}
		}
	}
	return "Letter Not Found"
}
