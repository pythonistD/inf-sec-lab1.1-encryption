package common

var SpecialChars = map[rune]struct{}{
	'\n': {},
	'\t': {},
	'\r': {},
}

const (
	RusLeft        rune = 1040
	RusRight       rune = 1103
	RussianCapital      = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
	RussianLower        = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	RussianAll          = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	EnglishCapital      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	EnglishAll          = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)
