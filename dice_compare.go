// Package dicecompare implementation of dice-coefficient in go
package dicecompare

import "strings"

// DiceBigram hold the bigram needed for calculation
type DiceBigram struct {
	first  rune
	second rune
}

// CompareList return string and coefficient from the list that have the highest-dice coefficient of the given string
func CompareList(str string, strList []string) (string, float64) {
	// idx list for highest coefficient
	maxIdx := -1
	// max coefficient from the list
	maxCoeff := 0.0
	for idx, s := range strList {
		coeff := Compare(str, s)
		if coeff > maxCoeff {
			maxCoeff = coeff
			maxIdx = idx
		}
	}
	if maxIdx != -1 {
		return strList[maxIdx], maxCoeff
	}
	return "", maxCoeff
}

// Compare return the dice coefficient of two given string
func Compare(aStr, bStr string) float64 {
	aStr = strings.ToLower(aStr)
	bStr = strings.ToLower(bStr)
	if aStr == bStr {
		return 1
	}
	aBigram := SplitBigramMultiple(aStr)
	bBigram := SplitBigramMultiple(bStr)
	intersection := 0
	for aKey := range aBigram {
		if bBigram[aKey] == true {
			intersection++
		}
	}
	return float64(intersection) / float64(len(bStr))
}

// SplitBigramMultiple split the word into set of bigrams and can accept words with whitespace
func SplitBigramMultiple(str string) map[DiceBigram]bool {
	bigramMap := make(map[DiceBigram]bool)
	wordList := strings.Fields(str)

	for _, word := range wordList {
		for bigramKey := range SplitBigram(word) {
			bigramMap[bigramKey] = true
		}
	}

	return bigramMap
}

// SplitBigram split the word into set of bigrams
func SplitBigram(str string) map[DiceBigram]bool {
	var last rune
	bigramMap := make(map[DiceBigram]bool)
	for idx, ch := range str {
		if idx > 0 {
			bigramMap[DiceBigram{last, ch}] = true
		}
		last = ch
	}
	return bigramMap
}
