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
	if len(aStr) < 2 {
		return 0
	}
	if aStr == bStr {
		return 1
	}
	aBigram := SplitBigram(aStr)
	bBigram := SplitBigram(bStr)
	intersection := 0
	for aKey := range aBigram {
		if bBigram[aKey] == true {
			intersection++
		}
	}
	return float64(intersection) * 2 / float64(len(aBigram)+len(bBigram))
}

// SplitBigram split the word into set of bigrams
func SplitBigram(str string) map[DiceBigram]bool {
	var last rune
	wordList := strings.Fields(str)
	str = strings.Join(wordList, " ")
	bigramMap := make(map[DiceBigram]bool)
	for idx, ch := range str {
		if idx > 0 {
			bigramMap[DiceBigram{last, ch}] = true
		}
		last = ch
	}
	return bigramMap
}
