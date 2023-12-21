package main

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumberals = []RomanNumeral{
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}
var result strings.Builder

func ConvertToRoman(arabic int) string {
	for _, numeral := range allRomanNumberals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}
