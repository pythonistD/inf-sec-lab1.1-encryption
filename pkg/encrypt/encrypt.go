package encrypt

import (
	"fmt"
	"github.com/pythonistD/inf-sec-lab1.1-encryption/pkg/dto"
	"os"
)

const (
	RusLeft  rune = 1040
	RusRight rune = 1103
)

func writeEncryptedText(symbols []rune) {
	var file *os.File
	file, err := os.Open("files/outData.txt")
	text := string(symbols)
	if err != nil {
		file, _ = os.Create("files/outData.txt")
	}
	_, err = file.WriteString(text)

}

func CaesarCipherEncrypt(dto dto.InputDataDto) {
	var encodedSymbols []rune
	if dto.Lang == "ru" {
		for _, v := range dto.Symbols {
			shift := rune(dto.Shift % 64)
			var pos rune
			if index := v + shift; index > RusRight {
				index = index - RusRight - 1
				pos = RusLeft + index
			} else {
				pos = index
			}
			encodedSymbols = append(encodedSymbols, pos)
		}
	}
	writeEncryptedText(encodedSymbols)
	fmt.Println(string(encodedSymbols))
}
