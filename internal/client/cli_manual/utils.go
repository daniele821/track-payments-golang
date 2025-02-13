package cli_manual

import "time"

func matchEveryLenght(str, match string) bool {
	if len(str) > len(match) {
		return false
	}
	return len(str) >= 1 && str == match[:len(str)]
}

func matchEveryLenghtFromAnyWords(str string, matches []string) bool {
	for _, match := range matches {
		if matchEveryLenght(str, match) {
			return true
		}
	}
	return false
}

func splitter(data []string) (splitted [][]string) {
	tmpArray := []string{}
	for _, elem := range append(data, "@") {
		if elem == "@" {
			splitted = append(splitted, tmpArray)
			tmpArray = []string{}
		} else {
			tmpArray = append(tmpArray, elem)
		}
	}
	return splitted
}

func getDateAndTime() (dateStr, timeStr string) {
	return time.Now().Format("2006/01/02"), time.Now().Format("15:04")
}
