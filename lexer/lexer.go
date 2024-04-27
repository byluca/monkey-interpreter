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

// NextToken genera il prossimo token basato sul carattere corrente nel lexer.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
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
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0: // Considera '0' come segnale di fine file
		tok = token.Token{Type: token.EOF, Literal: ""}
	default: // Gestisci qualsiasi carattere non riconosciuto come illegale
		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar() // Leggi il prossimo carattere prima di ritornare il token
	return tok
}
