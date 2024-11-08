package decrypt

import (
	"fmt"
	"github.com/pythonistD/inf-sec-lab1.1-encryption/pkg/common"
	"github.com/pythonistD/inf-sec-lab1.1-encryption/pkg/dto"
)

func CaesarCipherDecrypt(dto dto.InputDataDto) []rune {
	var decodedSymbols []rune
	/*
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
				decodedSymbols = append(decodedSymbols, pos)
			}
		}*/
	for _, v := range dto.Symbols {
		var shift rune = 0
		if _, exist := common.SpecialChars[v]; !exist {
			shift = rune(dto.Shift)
		}
		decodedSymbols = append(decodedSymbols, v-shift)
	}
	fmt.Println(string(decodedSymbols))
	return decodedSymbols
}
