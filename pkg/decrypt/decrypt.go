package decrypt

import (
	"fmt"
	"github.com/pythonistD/inf-sec-lab1.1-encryption/pkg/dto"
)

const (
	RusLeft  rune = 1040
	RusRight rune = 1103
)

func CaesarCipherDecrypt(dto dto.InputDataDto) {
	var encodedSymbols []rune
	if dto.Lang == "ru" {
		for _, v := range dto.Symbols {
			shift := rune(dto.Shift % 64)
			var pos rune
			if index := v - shift; index < RusLeft {
				index = RusLeft - index - 1
				pos = RusRight - index
			} else {
				pos = index
			}
			encodedSymbols = append(encodedSymbols, pos)
		}
	}
	fmt.Println(string(encodedSymbols))
}
