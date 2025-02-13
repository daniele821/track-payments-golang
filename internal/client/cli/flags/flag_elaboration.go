package flags

func (f FlagParsed) GetWordFlag(flag string, args int) ([]string, error) {
	return nil, nil
}

func (f FlagParsed) GetLetterFlag(flag string, args int) ([]string, error) {
	return nil, nil
}

func (f FlagParsed) Empty() bool {
	return len(f.flagArgs) == 0
}

func (f FlagParsed) GetNotEmptyErrorMsg() string {
	return ""
}
