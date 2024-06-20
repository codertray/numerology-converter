package conversions

import (
	"fmt"

	nm "example.com/numerology/numbers"

	"strconv"
	"strings"
)

func Format(str string) []string {
	noSpaces := strings.ReplaceAll(str, " ", "")
	allUpper := strings.ToUpper(noSpaces)
	allchars := strings.TrimSuffix(allUpper, "\f")
	noLines := strings.TrimSuffix(allchars, "\n")
	splitStr := strings.Split(noLines, "")
	return splitStr
}

func Calculate(chars []string) int {
	num := 0
	for _, value := range chars {
		digit := nm.Numbers[value]
		num += digit
	}
	return num
}

func Merger(val int) int {
	str := strconv.Itoa(val)
	splitStr := strings.Split(str, "")
	fmt.Println(splitStr)
	num := 0
	for _, val := range splitStr {
		digit, _ := strconv.Atoi(val)
		num += digit
	}
	return num
}
