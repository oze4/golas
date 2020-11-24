package golas

import (
	"bufio"
	"bytes"
	"io"
)

// Lexer is a lexer
type Lexer struct {
	tokens           chan Token
	isFirstDotOnLine bool
	value            *bytes.Buffer
	char             rune
	line             int
	linePosition     int
	reader           *bufio.Reader
	handler          HandlerFunc
}

// NewLexer creates a new Lexer
func NewLexer(r io.Reader) *Lexer {
	lexer := Lexer{
		reader: bufio.NewReader(r),
		tokens: make(chan Token, 1),
		value:  &bytes.Buffer{},
	}

	return &lexer
}

// NextToken reads the next token from our tokens chan
func (lexer *Lexer) NextToken() Token {
	for {
		select {
		case token := <-lexer.tokens:
			lexer.value.Reset()
			return token
		default:
			lexer.handler = lexer.handler(lexer)
		}
	}
}

// Run runs some handler
func (lexer *Lexer) Run(handler HandlerFunc) {
	lexer.handler = handler
}

// Start is shorthand for lexer.Run(HandleBegin)
func (lexer *Lexer) Start() {
	lexer.Run(HandleBegin)
}

// State returns the current handler of our lexer
func (lexer *Lexer) State() HandlerFunc {
	return lexer.handler
}

// * Private methods *

// emit places a token of type t on our tokens chan
func (lexer *Lexer) emit(t TokenType) {
	lexer.tokens <- Token{t, lexer.value.String()}
}

// dumpLine reads from current line position until
// end of line, without writing to our value buffer
func (lexer *Lexer) dumpLine() {
	for lexer.char != Flags.NewLine {
		lexer.step()
	}
	lexer.value.Reset()
}

func (lexer *Lexer) isLinePositionAt(i int) bool {
	if lexer.linePosition != i {
		return false
	}
	return true
}

// step consumes the next rune from the current line
func (lexer *Lexer) step() {
	ch, _, err := lexer.reader.ReadRune()
	if err != nil {
		ch = Flags.EOF
	}

	if ch == Flags.NewLine {
		lexer.line++
		lexer.linePosition = 0
		lexer.isFirstDotOnLine = false
	} else {
		lexer.linePosition++
	}

	if ch == '.' {
		if !lexer.isFirstDotOnLine {
			lexer.isFirstDotOnLine = true
		} else {
			lexer.isFirstDotOnLine = false
		}
	}

	lexer.value.WriteRune(ch)
	lexer.char = ch
}

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
