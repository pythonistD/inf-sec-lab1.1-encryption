package encrypt

import (
	"fmt"
	"github.com/pythonistD/inf-sec-lab1.1-encryption/pkg/dto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeleteRepsFromKeyword(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "helo"},
		{"abcd  cds", "abcds"},
		{"diplomat", "diplomat"},
		{"abcd!    fe", "abcdfe"},
		{"КАК ДЫМ ОТЕЧЕСТВА НАМ СЛАДОК И ПРИЯТЕН", "КАДЫМОТЕЧСВНЛИПРЯ"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			res, err := deleteRepsFromKeyword(testCase.input)
			require.NoError(t, err)
			require.Equal(t, testCase.expected, res)
		})
	}
}

func TestCreateCypherTable(t *testing.T) {
	testCases := []struct {
		shift    int
		keyword  string
		lang     string
		expected string
	}{
		{5, "diplomat", "en", "VWXYZDIPLOMATBCEFGHJKNQRSU"},
		{3, "key", "en", "WXZKEYABCDFGHIJLMNOPQRSTUV"},
		{3, "КАК ДЫМ ОТЕЧЕСТВА НАМ СЛАДОК И ПРИЯТЕН", "ru", "ЬЭЮКАДЫМОТЕЧСВНЛИПРЯБГЁЖЗЙУФХЦШЩЪ"},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Shift: %d Keyword: %s", testCase.shift, testCase.keyword), func(t *testing.T) {
			keywordNoReps, _ := deleteRepsFromKeyword(testCase.keyword)
			res := createCypherString(testCase.shift, keywordNoReps, testCase.lang)
			require.Equal(t, testCase.expected, res)
		})
	}
}

func TestCaesarCipherEncrypt(t *testing.T) {
	testCases := []struct {
		keyword       string
		shift         int
		textToEncrypt string
		lang          string
		expected      string
	}{
		{"diplomat", 5, "Send more money", "en", "Hzby tcgz tcbzs"},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Shift: %d Keyword: %s", testCase.shift, testCase.keyword), func(t *testing.T) {
			inputDataDto := dto.InputDataDto{Symbols: []rune(testCase.textToEncrypt), Shift: testCase.shift, Keyword: testCase.keyword, Lang: testCase.lang}
			encryptTable, _ := CreateEncryptTable(inputDataDto)
			encryptedData, _ := CaesarCipherEncrypt([]rune(testCase.textToEncrypt), encryptTable)
			require.Equal(t, testCase.expected, encryptedData)
		})
	}

}

func TestCaesarCipherDEncrypt(t *testing.T) {
	testCases := []struct {
		keyword       string
		shift         int
		textToDecrypt string
		lang          string
		expected      string
	}{
		{"diplomat", 5, "Hzby tcgz tcbzs", "en", "Send more money"},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Shift: %d Keyword: %s", testCase.shift, testCase.keyword), func(t *testing.T) {
			inputDataDto := dto.InputDataDto{Symbols: []rune(testCase.textToDecrypt), Shift: testCase.shift, Keyword: testCase.keyword, Lang: testCase.lang}
			encryptTable, _ := CreateEncryptTable(inputDataDto)
			decryptTable, _ := CreateDecryptTable(encryptTable)
			decryptedData, _ := CaesarCipherDecrypt([]rune(testCase.textToDecrypt), decryptTable)
			require.Equal(t, testCase.expected, decryptedData)
		})
	}

}
