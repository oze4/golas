package token

// CWLS LAS 2.0 Standard
// http://www.cwls.org/wp-content/uploads/2014/09/LAS_20_Update_Jan2014.pdf

// TokenType represents a lexical token type
type TokenType uint

// Token represents a lexical token
type Token struct {
	Type  TokenType
	Value string
}

// comment
const (
	TEndOfFile TokenType = iota
	TSyntaxError

	TIdentifier
	TValue

	// Flags
	TComment // #
	TTilde   // ~

	// Sections : Required
	TVersionInformation
	TWellInformation
	TCurveInformation
	TASCIILogData

	// Sections : Optional
	TParameterInformation
	TOther

	// Sections : Custom
	TSection
	TSectionCustom

	// Line Delimeters
	TMnemonic
	TSpace
	TUnits
	TData
	TDescription

	// Version Information
	TVers
	TWrap
)
