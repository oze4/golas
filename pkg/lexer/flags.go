package lexer

// Flags are a rune representation of a TokenType
var Flags = flags{
	Comment:  '#',
	Data:     ':',
	Section:  '~',
	EOF:      rune(-1),
	NewLine:  '\n',
	Mnemonic: '.',
}

type flags struct {
	Comment  rune
	Data     rune
	Section  rune
	EOF      rune
	NewLine  rune
	Mnemonic rune
}
