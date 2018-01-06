package ylua

import "strconv"

// TokenType is the set of lexical tokens of the Go programming language.
type TokenType int

// The list of tokens.
const (
	TokError TokenType = iota
	TokEOF
	TokComment

	literalBegin
	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	TokIdent      // main
	TokInt        // 12345
	TokFloat      // 123.45
	TokBool       //true false
	TokSingleQuot // 'abc'
	TokDoubleQuot // "abc"
	literalEnd

	operatorBegin
	// Operators and delimiters
	TokAssign // =

	TokLBrack // [
	TokRBrack // ]

	TokLBrace // {
	TokRBrace // }

	TokComma // ,

	TokColon // :
	operatorEnd

	keywordBegin
	// Keywords
	keywordEnd
)

var tokenTypes = [...]string{
	TokError:   "ERROR",
	TokEOF:     "EOF",
	TokComment: "COMMENT",

	TokIdent:      "Ident",            // main
	TokInt:        "Int",              // 12345
	TokFloat:      "Float",            // 123.45
	TokBool:       "Bool",             // true, false
	TokSingleQuot: "SingleQuotString", // 'abc'
	TokDoubleQuot: "DoubleQuotString", // "abc"

	// Operators and delimiters
	TokAssign: "=", // =

	TokLBrack: "[", // [
	TokRBrack: "]", // ]

	TokLBrace: "{", // {
	TokRBrace: "}", // }

	TokComma: ",", // ,

	TokColon: ":", // :
}

// String returns the string corresponding to the token tok.
// For operators, delimiters, and keywords the string is the actual
// token character sequence (e.g., for the token ADD, the string is
// "+"). For all other tokens the string corresponds to the token
// constant name (e.g. for the token IDENT, the string is "IDENT").
//
func (tok TokenType) String() string {
	s := ""
	if 0 <= tok && tok < TokenType(len(tokenTypes)) {
		s = tokenTypes[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

// Predicates

// IsLiteral returns true for tokens corresponding to identifiers
// and basic type literals; it returns false otherwise.
//
func (tok TokenType) IsLiteral() bool {
	return literalBegin < tok && tok < literalEnd
}

// IsOperator returns true for tokens corresponding to operators and
// delimiters; it returns false otherwise.
//
func (tok TokenType) IsOperator() bool {
	return operatorBegin < tok && tok < operatorEnd
}

// IsKeyword returns true for tokens corresponding to keywords;
// it returns false otherwise.
//
func (tok TokenType) IsKeyword() bool {
	return keywordBegin < tok && tok < keywordEnd
}
