package golas

import (
	"io"
	"strings"
)

// Parser parses a las file
type Parser struct {
	lex *Lexer
}

// NewParser creates a new parser
func NewParser(r io.Reader) Parser {
	return Parser{
		lex: NewLexer(r),
	}
}

// Parse odoealkj
func (p *Parser) Parse() LAS {
	output := LAS{}
	p.lex.Start()
	p.parse(&output)
	return output
}

func (p *Parser) parse(lasFile *LAS) {
	var (
		token         Token
		comments      string
		line          *Line
		sectn         *Section
		isCustomSectn bool = false
	)

    // If our lexer has not been started do not continue
    panicIfLexerNotStarted(p.lex)
    
Loop:
	for {
		token = p.lex.NextToken()

		switch token.Type {
		case TEndOfFile:
			break Loop

		case TVersionInformation:
			lasFile.VersionInformation, sectn = Section{Name: "Version Information"}, &lasFile.VersionInformation

		case TWellInformation:
			lasFile.WellInformation, sectn = Section{Name: "Well Information"}, &lasFile.WellInformation

		case TCurveInformation:
			lasFile.CurveInformation, sectn = Section{Name: "Curve Information"}, &lasFile.CurveInformation

		case TParameterInformation:
			lasFile.ParameterInformation, sectn = Section{Name: "Parameter Information"}, &lasFile.ParameterInformation

		case TOther:
			lasFile.Other, sectn = Section{Name: "Other"}, &lasFile.Other

		case TASCIILogData:
			break Loop
			//section = Section{Name: "ASCII Log Data"}
			//las.ASCIILogData = LogData{Section{}}

		case TSectionCustom:
			sectn = &Section{Name: strings.TrimSpace(token.Value)}
			isCustomSectn = true

		case TMnemonic:
			if isCustomSectn {
				lasFile.CustomSections = append(lasFile.CustomSections, Section{Name: strings.TrimSpace(token.Value)})
				isCustomSectn = false
				continue
			} else {
                line = &Line{}
                line.Mnem = strings.TrimSpace(token.Value)
            }

		case TUnits:
			line.Units = strings.TrimSpace(token.Value)

		case TData:
			line.Data = strings.TrimSpace(token.Value)

		case TDescription:
			line.Description = strings.TrimSpace(token.Value)
			sectn.Data = append(sectn.Data, *line)
			sectn.Comments = comments
			comments = ""

		case TComment:
			comments += token.Value
		}
	}
}

func panicIfLexerNotStarted(l *Lexer) {
	if l.State() == nil {
		panic("cannot start parser while lexer state is nil")
	}
}
