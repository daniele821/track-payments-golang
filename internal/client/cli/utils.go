package cli

import (
	"errors"
	"strconv"
	"strings"
	"time"
	"unicode"
)

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

func fillDataIfEmpty(data, instead string) string {
	if len(data) <= 1 {
		return instead
	}
	return data
}

func fillDataIfEmptyOpt(data, instead *string) *string {
	if data == nil || len(*data) <= 1 {
		return instead
	}
	return data
}

func parsePrice(priceStr string) (int, error) {
	priceParts := strings.Split(priceStr, ".")
	if len(priceParts) > 2 {
		return 0, errors.New("invalid price: too many dots (max 1 allowed)")
	}
	priceInt, err := strconv.Atoi(priceParts[0])
	if err != nil {
		return 0, err
	}
	if len(priceParts) == 1 {
		return priceInt * 100, nil
	}
	for _, rune := range priceParts[1] {
		if !unicode.IsDigit(rune) {
			return 0, errors.New("invalid price: after the dot, only digits are allowed")
		}
	}
	priceDec, err := strconv.Atoi((priceParts[1] + "00")[:2])
	if err != nil {
		return 0, err
	}
	return priceInt*100 + priceDec, nil
}

func strPrice(price int) string {
	priceStr := strconv.Itoa(price)
	if len(priceStr) < 3 {
		priceStr = strings.Repeat("0", 3-len(priceStr)) + priceStr
	}
	return priceStr[:len(priceStr)-2] + "." + priceStr[len(priceStr)-2:] + "€"
}

func getDateAndTime() (dateStr, timeStr string) {
	return time.Now().Format("2006/01/02"), time.Now().Format("15:04")
}
