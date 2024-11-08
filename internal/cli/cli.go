package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)
import (
	"github.com/pythonistD/inf-sec-lab1.1-encryption/internal/fileio"
	"github.com/pythonistD/inf-sec-lab1.1-encryption/pkg/decrypt"
	"github.com/pythonistD/inf-sec-lab1.1-encryption/pkg/dto"
	"github.com/pythonistD/inf-sec-lab1.1-encryption/pkg/encrypt"
)

func cryptOrDecrypt() string {
	var flag = true
	var mod string
	var err error
	for flag {
		fmt.Println("1: Зашифровать\n2: Расшифровать")
		_, err = fmt.Scan(&mod)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if mod != "1" && mod != "2" {
			fmt.Printf("Для выбора режима "+
				"необходимо выбрать 1 или 2, a не %v\n", mod)
			continue
		}
		flag = false
	}
	return mod
}

func fromCmdOrFile() string {
	var flag = true
	var mod string
	var err error
	for flag {
		fmt.Println("1: Ручной ввод(в консоль)\n2: Чтение из файла")
		_, err = fmt.Scan(&mod)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if mod != "1" && mod != "2" {
			fmt.Printf("Для выбора режима "+
				"необходимо выбрать 1 или 2, a не %v\n", mod)
			continue
		}
		flag = false
	}
	return mod
}

func getShift() string {
	var flag = true
	var mod string
	var err error
	var n int
	for flag {
		fmt.Println("Введите Ключ шифрования - значение сдвига N")
		_, err = fmt.Scan(&mod)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n, err = strconv.Atoi(mod); err != nil {
			fmt.Printf("Необходимо ввести одно число"+
				"без пробелов, а не %v\n", mod)
			continue
		} else if n <= 0 {
			fmt.Println("Значение сдвига должное быть > 0")
			continue
		}
		flag = false
	}
	return mod
}
func getLang() string {
	var flag = true
	var mod string
	var err error
	for flag {
		fmt.Println("Выберите язык текста: ru, en")
		_, err = fmt.Scan(&mod)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if mod != "ru" && mod != "en" {
			fmt.Printf("Необходимо выбрать один язык: ru или en"+
				", а не %v", mod)
			continue
		}
		flag = false
	}
	return mod
}

func getChars(desc io.Reader) []rune {
	scanner := bufio.NewScanner(desc)
	var lines []rune
	for scanner.Scan() {
		line := []rune(scanner.Text())
		line = append(line, '\n')
		lines = append(lines, line...)
	}
	fmt.Println("Содержимое файла:")
	fmt.Println("________________________")
	for _, v := range lines {
		fmt.Printf("%c", v)
	}
	fmt.Println("________________________")
	return lines
}

func getFileDescriptor() io.Reader {
	var flag = true
	var mod string
	var err error
	var desc io.Reader
	for flag {
		//fmt.Println("Введите путь до файла")
		//_, err = fmt.Scan(&mod)
		if err != nil {
			fmt.Println(err)
			continue
		}
		mod = "files/inData.txt"
		//mod = "files/outData.txt"
		if desc, err = os.Open(mod); err != nil {
			fmt.Printf("Ошибка чтения файла: %v\n", err)
			continue
		}
		flag = false
	}
	return desc
}

func Execute() {
	var err error
	var dataToWrite []rune

	mod := cryptOrDecrypt()
	fmt.Printf("Выбран режим: %s\n", mod)
	//inputOption := fromCmdOrFile()
	//lang := getLang()
	chars := getChars(getFileDescriptor())
	shift, _ := strconv.Atoi(getShift())
	inputDataDto := dto.InputDataDto{Symbols: chars, Shift: shift}
	if mod == "1" {
		dataToWrite = encrypt.CaesarCipherEncrypt(inputDataDto)
	} else if mod == "2" {
		dataToWrite = decrypt.CaesarCipherDecrypt(inputDataDto)
	}
	err = fileio.WriteText(dataToWrite, "./files/outData.txt")
	if err != nil {
		fmt.Printf("Ошибка во время выполнения программы: %v\n", err)
	}
}
