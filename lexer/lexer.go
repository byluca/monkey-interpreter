package lexer

import "monkey-interpreter/token"

type Lexer struct {
	input        string // La stringa di input che il lexer analizzerà
	position     int    // La posizione attuale nell'input (punta al carattere corrente).
	readPosition int    // La posizione futura nell'input (punta al prossimo carattere da leggere).
	ch           byte   // Il carattere corrente che il lexer sta esaminando.
}

// New crea e restituisce un nuovo Lexer inizializzato con l'input fornito.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Inizializza il primo carattere in 'ch'
	return l
}

// ReadChar legge il prossimo carattere nell'input e aggiorna le posizioni del lexer.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII NUL (0x00) usato come EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// newToken crea un nuovo token dato un tipo di token e un carattere.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace() // Skip whitespace to determine the next token.

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar() // Move to next char to complete the '=='
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.DASH, l.ch)
	case '*':
		tok = newToken(token.STAR, l.ch)
	case '/':
		tok = newToken(token.FORWARDSLASH, l.ch)
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar() // Move to next char to complete the '<='
			tok = token.Token{Type: token.LT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar() // Move to next char to complete the '>='
			tok = token.Token{Type: token.GT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar() // Move to next char to complete the '!='
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal) // Verify if it's a keyword after reading an identifier.
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar() // Read next character before returning the token
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {

	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// skipWhitespace salta gli spazi bianchi nel testo di input finché non trova un carattere significativo.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// peekChar restituisce il prossimo carattere nell'input senza avanzare la posizione del lexer.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0 // ASCII NUL (0x00) usato come EOF
	} else {
		return l.input[l.readPosition]
	}
}
