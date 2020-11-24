package golas

import (
	"fmt"
)

// HandlerFunc handles some handler
type HandlerFunc func(*Lexer) HandlerFunc

// HandleBegin is a state function
func HandleBegin(lexer *Lexer) HandlerFunc {
	var stateFn HandlerFunc
	for {
		switch lexer.char {
		case Flags.Section:
			stateFn = HandleSection
			goto Done
		case Flags.Mnemonic:
			stateFn = HandleMnemonic
			goto Done
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
	if !lexer.isLinePositionAt(1) {
		panic(fmt.Errorf("invalid section : line %d : position %d", lexer.line+1, lexer.linePosition))
	}
	lexer.step()
	switch lexer.char {
	case 'V':
		lexer.emit(TVersionInformation)
	case 'W':
		lexer.emit(TWellInformation)
	case 'C':
		lexer.emit(TCurveInformation)
	case 'A':
		lexer.emit(TASCIILogData)
	case 'P':
		lexer.emit(TParameterInformation)
	case 'O':
		lexer.emit(TOther)
	default:
		lexer.emit(TSectionCustom)
	}
	lexer.dumpLine()
	return HandleBegin
}

// HandleComment lexes a comment within a line
func HandleComment(lexer *Lexer) HandlerFunc {
	for lexer.char != Flags.NewLine {
		lexer.step()
	}
	lexer.emit(TComment)
	return HandleBegin
}

// HandleMnemonic lexes a mnemonic within a non-ascii log data line
func HandleMnemonic(lexer *Lexer) HandlerFunc {
	if lexer.isFirstDotOnLine {
		for lexer.char != Flags.Mnemonic {
			lexer.step()
		}
		lexer.emit(TMnemonic)
		return HandleUnits
	}
	return HandleBegin
}

// HandleUnits lexes units within a non-ascii log data line
func HandleUnits(lexer *Lexer) HandlerFunc {
	lexer.step()
	for !isCharASpace(lexer) {
		lexer.step()
	}
	lexer.emit(TUnits)
	return HandleLineData
}

// HandleLineData lexes data within a non-ascii log data line
func HandleLineData(lexer *Lexer) HandlerFunc {
	for lexer.char != Flags.Data {
		lexer.step()
	}
	lexer.emit(TData)
	return HandleDescription
}

// HandleDescription lexes a description within a non-ascii log data line
func HandleDescription(lexer *Lexer) HandlerFunc {
	for lexer.char != '\n' {
		lexer.step()
	}
	lexer.emit(TDescription)
	return HandleBegin
}

func isCharASpace(lexer *Lexer) bool {
	if string(lexer.char) != " " {
		return false
	}
	return true
}
