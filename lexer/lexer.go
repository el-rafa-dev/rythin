package rythin

import (
	"unicode"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	line         int
	column       int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input:  input,
		line:   1,
		column: 0,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	if l.ch == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	switch l.ch {
	case '.':
		startCol := l.column
		return l.newTokenWithPos(TOKEN_DOT, ".", startCol)
	case ';':
		if l.peekChar() == ';' {
			pos := l.position
			for l.ch != '\n' && l.ch != 0 {
				l.readChar()
			}
			return Token{Type: TOKEN_COMMENT, Literal: l.input[pos:l.position], Line: l.line, Column: l.column}
		}
		startCol := l.column
		return l.newTokenWithPos(TOKEN_COMMENT, ";", startCol)
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			startCol := l.column
			return l.newTokenWithPos(TOKEN_EQUAL, "==", startCol)
		}
		startCol := l.column
		return l.newTokenWithPos(TOKEN_ASSIGN, "=", startCol)
	case '!':
		if l.peekChar() == '=' {
			l.readChar() // move to '='
			startCol := l.column
			return l.newTokenWithPos(TOKEN_NOT_EQUAL, "!=", startCol)
		}

	case '+':
		if l.peekChar() == '+' {
			l.readChar()
			startCol := l.column
			return l.newTokenWithPos(TOKEN_PLUSPLUS, "++", startCol)
		}
		startCol := l.column
		return l.newTokenWithPos(TOKEN_PLUS, "+", startCol)
	case '-':
		if l.peekChar() == '>' {
			l.readChar()
			startCol := l.column
			return l.newTokenWithPos(TOKEN_ARROW_SET, "->", startCol)
		}
		startCol := l.column
		return l.newTokenWithPos(TOKEN_MINUS, "-", startCol)
	case '*':
		startCol := l.column
		return l.newTokenWithPos(TOKEN_MULTIPLY, "*", startCol)
	case '/':
		startCol := l.column
		return l.newTokenWithPos(TOKEN_DIVIDE, "/", startCol)
	case '&':
		startCol := l.column
		return l.newTokenWithPos(TOKEN_REF, "&", startCol)
	case '$':
		startCol := l.column
		return l.newTokenWithPos(TOKEN_DEREF, "$", startCol)
	case ':':
		startCol := l.column
		return l.newTokenWithPos(TOKEN_COLON, ":", startCol)
	case ',':
		startCol := l.column
		return l.newTokenWithPos(TOKEN_COMMA, ",", startCol)
	case '(':
		startCol := l.column
		return l.newTokenWithPos(TOKEN_LPAREN, "(", startCol)
	case ')':
		startCol := l.column
		return l.newTokenWithPos(TOKEN_RPAREN, ")", startCol)
	case '[':
		startCol := l.column
		return l.newTokenWithPos(TOKEN_LBRACKET, "[", startCol)
	case ']':
		startCol := l.column
		return l.newTokenWithPos(TOKEN_RBRACKET, "]", startCol)
	case '\\':
		if l.peekChar() == 'n' {
			l.readChar()
			startCol := l.column
			return l.newTokenWithPos(TOKEN_NEW_LINE, "\\n", startCol)
		}
	case '"':
		startCol := l.column
		literal := l.readString()
		return Token{Type: TOKEN_STRING, Literal: literal, Line: l.line, Column: startCol}

	case 0:
		return Token{Type: TOKEN_EOF, Literal: "", Line: l.line, Column: l.column}
	}
	if isLetter(l.ch) {
		startCol := l.column
		lit := l.readIdentifier()
		typ := LookupIdent(lit)
		return Token{Type: typ, Literal: lit, Line: l.line, Column: startCol}
	}

	if isDigit(l.ch) {
		startCol := l.column
		lit := l.readNumber()
		return Token{Type: TOKEN_NUMBER, Literal: lit, Line: l.line, Column: startCol}
	}

	l.readChar()
	return Token{Type: "ILLEGAL", Literal: string(l.ch), Line: l.line, Column: l.column}
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readString() string {
	l.readChar()
	start := l.position
	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}
	str := l.input[start:l.position]
	l.readChar()
	return str
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

/*func (l *Lexer) newToken(tokenType TokenType, ch string) Token {
	tok := Token{
		Type:    tokenType,
		Literal: ch,
		Line:    l.line,
		Column:  l.column - len(ch),
	}
	l.readChar()
	return tok
}*/

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch))
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

var keywords = map[string]TokenType{
	"print":    TOKEN_PRINT,
	"printnl":  TOKEN_PRINT_NEW_LINE,
	"printe":   TOKEN_PRINT_ERROR,
	"printel":  TOKEN_PRINT_ERROR_LOG,
	"def":      TOKEN_DEF,
	"using":    TOKEN_USING,
	"from":     TOKEN_FROM,
	"get":      TOKEN_GET,
	"loop":     TOKEN_LOOP,
	"stop":     TOKEN_STOP,
	"continue": TOKEN_CONTINUE,
	"if":       TOKEN_IF,
	"but":      TOKEN_BUT,
	"try":      TOKEN_TRY,
	"catch":    TOKEN_CATCH,
	"return":   TOKEN_RETURN,
	"alloc":    TOKEN_ALLOC,
	"flush":    TOKEN_FLUSH,
	"input":    TOKEN_INPUT,
	"fwrite":   TOKEN_FWRITE,
	"fread":    TOKEN_FREAD,
	"mkdir":    TOKEN_MKDIR,
	"len":      TOKEN_LEN,
	"has":      TOKEN_HAS,
	"int":      TOKEN_INT,
	"bool":     TOKEN_BOOL,
	"str":      TOKEN_STRING,
	"lint":     TOKEN_LONG_INT,
	"void":     TOKEN_VOID,
	"char":     TOKEN_CHAR,
	"charseq":  TOKEN_CHARSEQ,
	"true":     TOKEN_BOOLEAN,
	"false":    TOKEN_BOOLEAN,
	"nil":      TOKEN_NULL,
	"main":     TOKEN_MAIN_FUNC,
	"null":     TOKEN_NULL,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return TOKEN_IDENTIFIER
}

func (l *Lexer) newTokenWithPos(tokenType TokenType, literal string, col int) Token {
	tok := Token{
		Type:    tokenType,
		Literal: literal,
		Line:    l.line,
		Column:  col,
	}
	l.readChar()
	return tok
}
