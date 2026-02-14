package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Problems")
	// fmt.Println("Number: ", EnglishWordToInt("One Hundred Twenty Three"))
	fmt.Println("WORD: ", IntToEnglishWord(123456))
	fmt.Println("WORD: ", IntToEnglishWord(2234567890))
}

/**
* NOTE: Constraints: Max value is 2^31 - 1
 */
func EnglishWordToInt(numberPhrase string) int {

	if numberPhrase == "Zero" {
		return 0
	}

	split := strings.Split(numberPhrase, " ")

	lowNumMap := map[string]int{"One": 1, "Two": 2, "Three": 3, "Four": 4, "Five": 5, "Six": 6, "Seven": 7, "Eight": 8, "Nine": 9, "Ten": 10, "Eleven": 11, "Twelve": 12, "Thirteen": 13, "Fourteen": 14, "Fifteen": 15, "Sixteen": 16, "Seventeen": 17, "Eighteen": 18, "Nineteen": 19}
	tensNumMap := map[string]int{"Twenty": 20, "Thirty": 30, "Forty": 40, "Fifty": 50, "Sixty": 60, "Seventy": 70, "Eighty": 80, "Ninety": 90}
	hundred := "Hundred"
	bigNumMap := map[string]int{"Thousand": 1000, "Million": 1000000, "Billion": 1000000000}

	var result = 0
	var currentNumber = 0
	for idx, word := range split {
		low, foundLow := lowNumMap[word]
		if foundLow {
			currentNumber += low
		}

		if word == hundred {
			currentNumber *= 100
		}

		tens, foundTens := tensNumMap[word]
		if foundTens {
			currentNumber += tens
		}

		big, foundBig := bigNumMap[word]
		if foundBig {
			currentNumber *= big
			result += currentNumber
			currentNumber = 0
		}

		if idx == len(split)-1 {
			return result + currentNumber
		}
	}
	return result
}

func IntToEnglishWord(x int) string {
	lowNums := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	// tensNums := []string{"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	// hundred := "Hundred"
	// bigNumMap := []string{"Thousand", "Million", "Billion"}

	fmt.Println(x / 1000000000)
	fmt.Println(1001230000 / 1000000000)

	lowNumberMod := func(x, mod int) string {
		var idx = (x / mod) - 1
		if idx > len(lowNums) {
			fmt.Errorf("Invalid Mod Amount: %d. Choose smaller number.", mod)
		}
		return lowNums[idx]
	}
	bigNumberMod := func(x int) string {
		if x/100 > 0 {

		}
		return ""
	}

	var result = ""
	if x/1000000000 > 0 {
		result = lowNumberMod(x, 1000000000) + " Billion"
	}
	if x/1000000 > 0 {

	}

	return result
}
