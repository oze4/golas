package lexer

import (
	"fmt"

	"github.com/oze4/golas/pkg/token"
)

// HandlerFunc handles some handler
type HandlerFunc func(*Lexer) HandlerFunc

// HandleBegin is a state function
func HandleBegin(lexer *Lexer) HandlerFunc {
	fmt.Println("\nlexing begin")
	var stateFn HandlerFunc
	for {
		switch lexer.char {
		case Flags.EOF:
			stateFn = nil
			goto Done
		case Flags.Section:
			stateFn = HandleSection
			goto Done
		case Flags.Mnemonic:
			stateFn = HandleMnemonic
			goto Done
		case Flags.Data:
			stateFn = HandleLineData
		case Flags.Comment:
			stateFn = HandleComment
			goto Done
		default:
			lexer.step()
		}
	}
Done:
	return stateFn
}

// HandleSection alkdj
func HandleSection(lexer *Lexer) HandlerFunc {
	fmt.Println("lexing section")

	if !lexer.isLinePositionAt(1) {
		panic(fmt.Errorf("invalid section : line %d : position %d", lexer.line+1, lexer.linePosition))
	}

	lexer.step()

	switch lexer.char {
	case 'V':
		lexer.emit(token.TVersionInformation)
	case 'W':
		lexer.emit(token.TWellInformation)
	case 'C':
		lexer.emit(token.TCurveInformation)
	case 'A':
		lexer.emit(token.TASCIILogData)
	case 'P':
		lexer.emit(token.TParameterInformation)
	case 'O':
		lexer.emit(token.TOther)
	default:
		lexer.emit(token.TSectionCustom)
	}

	lexer.ignoreRestOfLine()
	return HandleBegin
}

// HandleComment lexes a comment within a line
func HandleComment(lexer *Lexer) HandlerFunc {
	fmt.Println("lexing comment")
	for lexer.char != Flags.NewLine {
		lexer.step()
	}
	lexer.emit(token.TComment)
	return HandleBegin
}

// HandleMnemonic lexes a mnemonic within a non-ascii log data line
func HandleMnemonic(lexer *Lexer) HandlerFunc {
	fmt.Println("lexing mnem")
	if lexer.isFirstDotOnLine {
		for lexer.char != Flags.Mnemonic {
			lexer.step()
		}
		lexer.emit(token.TMnemonic)
		return HandleUnits
	}
	return HandleBegin
}

// HandleUnits lexes units within a non-ascii log data line
func HandleUnits(lexer *Lexer) HandlerFunc {
	fmt.Println("lexing units")
	lexer.step()
	for !isCharASpace(lexer) {
		lexer.step()
	}
	lexer.emit(token.TUnits)
	return HandleLineData
}

// HandleLineData lexes data within a non-ascii log data line
func HandleLineData(lexer *Lexer) HandlerFunc {
	fmt.Println("lexing line data")
	times := 0
	for lexer.char != Flags.Data {
		if times > 2000 {
			panic("d")
		}
		lexer.step()
		times++
	}
	lexer.emit(token.TData)
	return HandleDescription
}

// HandleDescription lexes a description within a non-ascii log data line
func HandleDescription(lexer *Lexer) HandlerFunc {
	fmt.Println("lexing description")
	for lexer.char != '\n' {
		lexer.step()
	}
	lexer.emit(token.TDescription)
	return HandleBegin
}

func isCharASpace(lexer *Lexer) bool {
	if string(lexer.char) != " " {
		return false
	}
	return true
}
