package common

var SpecialChars = map[rune]struct{}{
	'\n': {},
	'\t': {},
	'\r': {},
}

const (
	RusLeft  rune = 1040
	RusRight rune = 1103
)
