package ylua

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type Token struct {
	typ  TokenType
	pos  Pos
	val  string
	line int
}

func (t Token) String() string {
	switch t.typ {
	case TokEOF:
		return "EOF"
	case TokError:
		return "Error"
	}
	return t.val

	// return fmt.Sprintf("%q", t.val)
}

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*lexer) stateFn

// lexer holds the state of the scanner.
type lexer struct {
	name   string  // the name of the input; used only for error reports
	input  string  // the string being scanned
	length int     //length of the input string
	state  stateFn // the next lexing function to enter
	start  Pos     // start position of this item
	pos    Pos     // current position in the input
	width  Pos     // width of last rune read from input
	line   int
	tokens chan Token // channel of scanned items
}

// Pos represents a byte position in the original input text from which
// this template was parsed.
type Pos int

const eof = -1

// var spaceCharacters = []rune{' ', '\t', '\n', '\r', ','}

// lex creates a new scanner for the input string.
//initializes itself to lex a string and launches the state machine as a goroutine, returning the lexer itself and a channel of items.
func lex(name, input string) *lexer {
	l := &lexer{
		name:   name,
		input:  input,
		length: len(input),
		tokens: make(chan Token, 10),
	}
	go l.run()
	return l
}

// next returns the next rune in the input.
func (l *lexer) next() rune {
	if int(l.pos) >= l.length {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = Pos(w)
	l.pos += l.width
	if r == '\n' {
		l.line++
	}
	return r
}

// run lexes the input by executing state functions until
// the state is nil.
func (l *lexer) run() {
	for l.state = lexAny(l); l.state != nil; {
		l.state = l.state(l)
	}
	close(l.tokens)
}

// emit passes an item back to the client.
func (l *lexer) emit(t TokenType) {
	l.tokens <- Token{t, l.start, l.input[l.start:l.pos], l.line}
	l.start = l.pos //move to current pos
}

// nextItem returns the next item from the input.
// Called by the parser, not in the lexing goroutine.
func (l *lexer) nextToken() Token {
	tok := <-l.tokens
	return tok
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.tokens <- Token{TokError, l.start, fmt.Sprintf(format, args...), l.line}
	return nil
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width

	if l.width == 1 && l.input[l.pos] == '\n' {
		l.line--
	}
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.start = l.pos
}

// lexInit start from beginning
func lexInit(l *lexer) stateFn {
	return lexAny(l)
}

//lexAny deal with patten
func lexAny(l *lexer) stateFn {
	// }
	switch r := l.next(); {
	case isSpace(r):
		return lexSpace
	case r == '{':
		l.emit(TokLBrace)
	case r == '}':
		l.emit(TokRBrace)
	case isAlphabet(r):
		return lexIdent
	case isDigit(r):
		return lexNumber
	case r == '=':
		l.emit(TokAssign)
	case r == '"':
		return lexDoubleQuote
	case r == '\'':
		return lexSingleQuote
	case r == '[':
		l.emit(TokLBrack)
	case r == ']':
		l.emit(TokRBrack)
	case r == eof:
		l.emit(TokEOF)
		return nil
	}
	return lexAny
}

// lexNumber scan a number. number can be a int or float
func lexNumber(l *lexer) stateFn {
Loop:
	for {
		switch r := l.next(); {
		case isDigit(r):
		case r == '.':
			return lexFloat
		default:
			l.backup()
			l.emit(TokInt)
			break Loop
		}
	}
	return lexAny
}

// lexFloat scan a float number, start from '.'
func lexFloat(l *lexer) stateFn {
Loop:
	for {
		switch r := l.next(); {
		case isDigit(r):
		default:
			l.backup()
			l.emit(TokFloat)
			break Loop
		}
	}
	return lexAny
}

// lexIndet scan a identity
func lexIdent(l *lexer) stateFn {
Loop:
	for {
		switch r := l.next(); {
		case isDigit(r):
		case isAlphabet(r):
		case r == '_':
		default:
			l.backup()
			word := l.input[l.start:l.pos]
			switch {
			case word == "true", word == "false":
				l.emit(TokBool)
			default:
				l.emit(TokIdent)
			}
			break Loop
		}
	}
	return lexAny
}

// lexSingleQuote scans a single quoted string.
func lexSingleQuote(l *lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case '\\':
			if r := l.next(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case eof, '\n':
			return l.errorf("未闭合的字符串。")
		case '\'':
			break Loop
		}
	}
	l.emit(TokSingleQuot)

	return lexAny
}

// lexDoubleQuote scans a double quoted string.
func lexDoubleQuote(l *lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case '\\':
			if r := l.next(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case eof, '\n':
			return l.errorf("未闭合的字符串。")
		case '"':
			break Loop
		}
	}
	l.emit(TokDoubleQuot)

	return lexAny
}

// lexSpace scans a run of space characters.
// One space has already been seen.
func lexSpace(l *lexer) stateFn {
	for {
		if !isSpace(l.next()) {
			l.backup()
			break
		}
	}
	//l.emit(TokS)
	l.ignore()
	return lexAny
}

// isSpace reports whether r is a space character.
// space include \t \n \r ,
func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == ','
}

// isAlphabet reports whether r is a alphabet character
func isAlphabet(r rune) bool {
	if r > unicode.MaxASCII {
		return false
	}
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

// isDigit reports whether r is digit
func isDigit(r rune) bool {
	if r > unicode.MaxASCII {
		return false
	}
	return r >= '0' && r <= '9'
}
