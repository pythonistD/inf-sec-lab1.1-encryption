package encrypt

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

import (
	"github.com/pythonistD/inf-sec-lab1.1-encryption/pkg/common"
	"github.com/pythonistD/inf-sec-lab1.1-encryption/pkg/dto"
)

const (
	LettersInRus = 33
	LettersInLat = 26
)

func deleteRepsFromKeyword(keyword string) (string, error) {
	keywordRunes := []rune(keyword)
	var builder strings.Builder
	seen := map[rune]struct{}{}
	for _, char := range keywordRunes {
		if _, exists := seen[char]; unicode.IsLetter(char) && !exists {
			seen[char] = struct{}{}
			builder.WriteRune(char)
		}
	}
	if builder.Len() <= 0 {
		return "", errors.New("неподходящее ключевое слово")
	}
	return builder.String(), nil
}

func deleteRepsFromTable(keyword string, table string) string {
	lettersToDelete := make(map[rune]struct{})
	var builder strings.Builder
	for _, char := range keyword {
		if _, exists := lettersToDelete[char]; !exists {
			lettersToDelete[char] = struct{}{}
		}
	}
	for _, char := range table {
		if _, exists := lettersToDelete[char]; !exists {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}

// createCypherString создаёт две таблицы для шифрования заглавных и строчных букв для данного
// сдвига, ключевого слова и языка
func createCypherString(shift int, keyword string, lang string) string {
	var tableSize int
	var table string
	keyword = strings.ToUpper(keyword)
	if lang == "ru" {
		tableSize = LettersInRus
		table = common.RussianCapital
	} else {
		tableSize = LettersInLat
		table = common.EnglishCapital
	}
	// Буквы, которые остались после того, как из исходного алфавита
	// вычли буквы из ключевого слова
	letters := deleteRepsFromTable(keyword, table)
	cypherTable := make([]rune, tableSize)
	var keywordPointer int
	for ind, ch := range []rune(keyword) {
		keywordPointer = (ind + shift) % tableSize
		cypherTable[keywordPointer] = ch
	}
	keywordPointer++
	for ind, ch := range []rune(letters) {
		pos := (keywordPointer + ind) % tableSize
		cypherTable[pos] = ch
	}
	return string(cypherTable)
}

func CreateEncryptTable(dto dto.InputDataDto) (map[rune]rune, error) {
	keyWord, err := deleteRepsFromKeyword(dto.Keyword)
	if err != nil {
		fmt.Println(keyWord)
		return nil, fmt.Errorf("ошибка шифрации: %w", err)
	}
	var shift int
	var alphabet []rune
	var numberOfLetters int
	if dto.Lang == "ru" {
		shift = dto.Shift % LettersInRus
		alphabet = []rune(common.RussianAll)
		numberOfLetters = LettersInRus
	} else if dto.Lang == "en" {
		shift = dto.Shift % LettersInLat
		alphabet = []rune(common.EnglishAll)
		numberOfLetters = LettersInLat
	}
	capitalCypherTable := createCypherString(shift, keyWord, dto.Lang)
	if len([]rune(capitalCypherTable)) > numberOfLetters {
		return nil, errors.New("некорректный ключ: ключ должен соответствовать выбранному языку")
	}
	lowerCypherTable := strings.ToLower(capitalCypherTable)
	finalCypherTable := capitalCypherTable + lowerCypherTable

	table := make(map[rune]rune)
	for ind, ch := range []rune(finalCypherTable) {
		table[alphabet[ind]] = ch
	}
	return table, nil
}

func CreateDecryptTable(encryptTable map[rune]rune) (map[rune]rune, error) {
	decryptTable := make(map[rune]rune)
	for key, value := range encryptTable {
		decryptTable[value] = key
	}
	return decryptTable, nil
}

func CaesarCipherEncrypt(symbols []rune, table map[rune]rune) (string, error) {
	var builder strings.Builder
	for _, s := range symbols {
		if _, exists := table[s]; unicode.IsLetter(s) && exists {
			builder.WriteRune(table[s])
		} else {
			builder.WriteRune(s)
		}
	}
	return builder.String(), nil
}

func CaesarCipherDecrypt(symbols []rune, table map[rune]rune) (string, error) {
	var builder strings.Builder
	for _, s := range symbols {
		if _, exists := table[s]; unicode.IsLetter(s) && exists {
			builder.WriteRune(table[s])
		} else {
			builder.WriteRune(s)
		}
	}
	return builder.String(), nil
}
