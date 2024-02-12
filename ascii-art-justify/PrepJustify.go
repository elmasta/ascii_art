package asciiArt

func PrepJustify(word string, data []string, width int, banner string) ([][]string, []int, []int) {
	// we make a list of words to know how many words we have per line
	temp := ""
	words := []string{}
	sList := [][]string{}
	for i := 0; i < len(word); i++ {
		v := word[i]
		if v != ' ' && v != '\\' {
			temp += string(v)
		} else if v == '\\' && word[i+1] == 'n' {
			words = append(words, temp)
			sList = append(sList, words)
			words = []string{}
			temp = ""
			i++
		} else if len(temp) > 0 {
			words = append(words, temp)
			temp = ""
		}
	}
	if len(temp) > 0 {
		words = append(words, temp)
		sList = append(sList, words)
	}

	// we open the banner file to retrieve the size of each letter
	wSize := 0
	lineSize := []int{}
	rest := []int{}
	for _, v := range sList {
		for _, subv := range v {
			for _, subSubV := range subv {
				if banner == "thinkertoy.txt" {
					wSize += len(data[(int(subSubV)-32)*9+1])-1
				} else {
					wSize += len(data[(int(subSubV)-32)*9+1])
				}
			}
		}
		lineSize = append(lineSize, wSize)
		wSize = 0
	}

	// at this point the table lineSize store how wide the lines are
	for i, v := range lineSize {
		if len(sList[i]) > 1 {
			rest = append(rest, (width - v) % (len(sList[i]) - 1))
			lineSize[i] = (width - v) / (len(sList[i]) - 1)
		} else {
			lineSize[i] = (width - v)
		}
	}
	// at this point the table lineSize store how wide spaces between words needs to be for each lines
	// end justify preparation
	return sList, lineSize, rest
}
