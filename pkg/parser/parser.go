package parser

import (
	"fmt"
	"io"
	"strings"

	"github.com/oze4/golas/pkg/file"
	"github.com/oze4/golas/pkg/lexer"
	"github.com/oze4/golas/pkg/token"
)

// Parser parses a las file
type Parser struct {
	lex *lexer.Lexer
}

// NewParser creates a new parser
func NewParser(r io.Reader) Parser {
	return Parser{
		lex: lexer.NewLexer(r),
	}
}

// Parse odoealkj
func (p *Parser) Parse() file.LASFile {
	output := file.LASFile{}
	p.lex.Start()
	fmt.Println("[STARTED] lexer")
	fmt.Println("[STARTING] parser")
	p.parse(&output)
	fmt.Println("[STOPPED] parser & lexer")
	return output
}

func (p *Parser) parse(writeTo *file.LASFile) {
	var (
		currentToken   token.Token
		comments       string
		addline        bool
		lineToAdd      *file.Line
		currentSection *file.Section = &file.Section{}
	)

	panicIfLexerNotStarted(p.lex)
Loop:
	for {
		currentToken = p.lex.NextToken()

		if addline {
			currentSection.Data = append(currentSection.Data, *lineToAdd)
			currentSection.Comments = comments
			comments = ""
			addline = false
		}

		if currentToken.Type != token.TValue {
			currentToken.Value = strings.TrimSpace(currentToken.Value)
		}

		switch currentToken.Type {
		// only here til we write the ascii log lexer stateFn
		case token.TASCIILogData:
			break Loop

		case token.TEndOfFile:
			break Loop

		case token.TVersionInformation:
			writeTo.VersionInformation = file.Section{Name: "Version Information"}
			currentSection = &writeTo.VersionInformation

		case token.TWellInformation:
			writeTo.WellInformation = file.Section{Name: "Well Information"}
			currentSection = &writeTo.WellInformation

		case token.TCurveInformation:
			writeTo.CurveInformation = file.Section{Name: "Curve Information"}
			currentSection = &writeTo.CurveInformation

		case token.TParameterInformation:
			writeTo.ParameterInformation = file.Section{Name: "Parameter Information"}
			currentSection = &writeTo.ParameterInformation

		case token.TOther:
			writeTo.Other = file.Section{Name: "Other"}
			currentSection = &writeTo.Other

		//case TASCIILogData:
		//	section = Section{Name: "ASCII Log Data"}
		//	las.ASCIILogData = LogData{Section{}}

		case token.TSectionCustom:
			panic("should not have custom sects yet")
			// section = Section{Name: token.Value}
			// las.CustomSections = append(las.CustomSections, section)

		case token.TMnemonic:
			lineToAdd = &file.Line{}
			lineToAdd.Mnem = currentToken.Value

		case token.TUnits:
			lineToAdd.Units = currentToken.Value

		case token.TData:
			lineToAdd.Data = currentToken.Value

		case token.TDescription:
			lineToAdd.Description = currentToken.Value
			addline = true

		case token.TComment:
			comments += currentToken.Value

		}
	}
}

func panicIfLexerNotStarted(l *lexer.Lexer) {
	if l.State() == nil {
		panic("cannot start parser while lexer state is nil")
	}
}
