package golas

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

// HandlerFunc func used as lexer state
type HandlerFunc func(*Lexer) HandlerFunc

// Lexer can perform lexical analysis on .las files
type Lexer struct {
	buffer   *bytes.Buffer
	char     rune
	dots     int
	handler  HandlerFunc
	line     int
	position int
	reader   *bufio.Reader
	tokens   chan Token
}

// NewLexer creates a new Lexer
func NewLexer(r io.Reader) *Lexer {
	return &Lexer{
		reader: bufio.NewReader(r),
		tokens: make(chan Token, 1),
		buffer: &bytes.Buffer{},
	}
}

// NextToken reads the next token from our tokens chan
func (l *Lexer) NextToken() Token {
	for {
		select {
		case token := <-l.tokens:
			return token
		default:
			if l.handler == nil {
				panic("private field `Lexer.handler` is nil : unable to acquire token without state")
			}
			l.handler = l.handler(l)
		}
	}
}

// Start sets our handler, in turn starting lexical analysis
func (l *Lexer) Start(hf HandlerFunc) {
	l.handler = hf
}

// emit places a token on our tokens chan
func (l *Lexer) emit(t TokenType) {
	l.tokens <- Token{t, strings.TrimSpace(l.buffer.String())}
	l.buffer.Reset()
}

// overwriteBuffer clears our buffer then writes a string to it
func (l *Lexer) overwriteBuffer(s string) {
	l.buffer.Reset()
	l.buffer.WriteString(s)
}

// step consumes the next rune from our reader
func (l *Lexer) step() {
	ch, _, err := l.reader.ReadRune()
	if err != nil {
		ch = CharEOF
	}
	// If no error, increment position before moving on
	l.position++

	switch ch {
	case CharNewLine:
		l.line++
		l.position = 0
		l.dots = 0
	case CharDot:
		l.dots = l.dots + 1
	}

	l.buffer.WriteRune(ch)
	l.char = ch
}

// stepUntil reads from current line position until we read a certain rune
func (l *Lexer) stepUntil(char rune) {
	for l.char != CharNewLine {
		l.step()
	}
}

// truncate our buffer by 1. If our buffer were a string, this removes the last character
func (l *Lexer) truncate() {
	l.buffer.Truncate(l.buffer.Len() - 1)
}
