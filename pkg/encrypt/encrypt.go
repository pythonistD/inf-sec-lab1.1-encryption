package encrypt

import (
	"fmt"
)

import (
	"github.com/pythonistD/inf-sec-lab1.1-encryption/pkg/common"
	"github.com/pythonistD/inf-sec-lab1.1-encryption/pkg/dto"
)

const (
	RusLeft  rune = 1040
	RusRight rune = 1103
)

func CaesarCipherEncrypt(dto dto.InputDataDto) []rune {
	var encodedSymbols []rune
	/*
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
		}*/
	for _, v := range dto.Symbols {
		var shift rune = 0
		if _, exist := common.SpecialChars[v]; !exist {
			shift = rune(dto.Shift)
		}
		encodedSymbols = append(encodedSymbols, v+shift)
	}
	fmt.Println(string(encodedSymbols))
	return encodedSymbols
}
