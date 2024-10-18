package lexer

import (
	"monkey-interpreter/token"
	"testing"
)

// TestConditionals testa il lexer con varie espressioni condizionali e operatori.
func TestConditionals(t *testing.T) {
	input := `!-/*5     
    5<10<5;
    if(5<10){
    return true;
    } else {
    return false;
    }
    10==10;
    9<=9;
    `

	// Definisce i token attesi e i loro valori letterali.
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.BANG, "!"},
		{token.DASH, "-"},
		{token.FORWARDSLASH, "/"},
		{token.STAR, "*"},
		{token.INT, "5"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.LT, "<"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "9"},
		{token.LT_EQ, "<="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
	}

	l := New(input) // Inizializza un nuovo lexer con l'input.

	// Scorre i test e verifica che il tipo e il valore letterale di ogni token siano corretti.
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tipo token errato. Atteso=%q, ottenuto=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - valore letterale errato. Atteso=%q, ottenuto=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

// TestNextToken testa il lexer con un input di dichiarazioni let e funzioni.
func TestNextToken(t *testing.T) {
	// Stringa completa di input per il lexer da analizzare.
	input := `let five = 5;
			  let ten = 10;
			  let add = fn(x, y) { 
              x + y; 
              }
              let result = add(five, ten);
              `

	// Definisce i token attesi e i loro valori letterali.
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""}, // Il token di fine file (EOF) solitamente non ha un valore letterale visibile.
	}

	l := New(input) // Inizializza un nuovo lexer con l'input.

	// Scorre i test e verifica che il tipo e il valore letterale di ogni token siano corretti.
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tipo token errato. Atteso=%q, ottenuto=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - valore letterale errato. Atteso=%q, ottenuto=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
