package lexer

import "monkey-interpreter/token"

// Lexer è responsabile della tokenizzazione dell'input. Analizza la stringa di input carattere per carattere
// e genera una sequenza di token che rappresentano i costrutti sintattici del linguaggio.
type Lexer struct {
	input        string // La stringa di input che il lexer analizzerà
	position     int    // La posizione attuale nell'input (punta al carattere corrente)
	readPosition int    // La posizione futura nell'input (punta al prossimo carattere da leggere)
	ch           byte   // Il carattere corrente che il lexer sta esaminando
}

// New crea e restituisce un nuovo Lexer inizializzato con l'input fornito.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Inizializza il primo carattere in 'ch'
	return l
}

// readChar legge il prossimo carattere nell'input e aggiorna le posizioni del lexer.
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

// NextToken restituisce il prossimo token presente nell'input.
// Esamina il carattere corrente e lo trasforma nel corrispondente token.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace() // Ignora spazi bianchi per identificare il prossimo token significativo.

	switch l.ch {
	case '=':
		// Verifica se l'operatore è "==" per uguaglianza
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch) // Altrimenti è un assegnamento "="
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
		// Controlla se è l'operatore "<=" (minore o uguale)
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.LT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.LT, l.ch) // Altrimenti è solo "<"
		}
	case '>':
		// Controlla se è l'operatore ">=" (maggiore o uguale)
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.GT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.GT, l.ch) // Altrimenti è solo ">"
		}
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '!':
		// Verifica se è "!=" (diverso)
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch) // Altrimenti è solo "!"
		}
	case 0:
		// Raggiunto la fine dell'input
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		// Verifica se è una lettera (parte di un identificatore o parola chiave)
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal) // Verifica se è una parola chiave
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber() // È un numero intero
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch) // Token non riconosciuto
		}
	}

	l.readChar() // Avanza al prossimo carattere
	return tok
}

// readIdentifier legge un identificatore (nome variabile o funzione) dall'input.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber legge un intero dall'input e lo restituisce come stringa.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// skipWhitespace salta gli spazi bianchi come spazi, tabulazioni e nuove righe.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// isLetter verifica se il carattere è una lettera (minuscola, maiuscola o underscore).
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

// isDigit verifica se il carattere è una cifra (0-9).
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
