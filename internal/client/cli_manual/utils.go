package cli_manual

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
