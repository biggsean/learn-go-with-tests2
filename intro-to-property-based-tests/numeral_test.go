package numeral

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
}

func TestConvertToRoman(t *testing.T) {
	for _, tt := range cases {
		tt := tt
		t.Run(fmt.Sprintf("%d gets converted to %q", tt.Arabic, tt.Roman), func(t *testing.T) {
			actual := ConvertToRoman(tt.Arabic)
			if actual != tt.Roman {
				t.Errorf("(%d): expected %s, actual %s", tt.Arabic, tt.Roman, actual)
			}

		})
	}
}
func TestConvertToArabic(t *testing.T) {
	for _, tt := range cases {
		tt := tt
		t.Run(fmt.Sprintf("%q gets converted to %d", tt.Roman, tt.Arabic), func(t *testing.T) {
			actual := ConvertToArabic(tt.Roman)
			if actual != tt.Arabic {
				t.Errorf("(%s): expected %d, actual %d", tt.Roman, tt.Arabic, actual)
			}

		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("Testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
