package token

// TokenType rappresenta il tipo di token come stringa
type TokenType string

// Token rappresenta un singolo token con il suo tipo e valore letterale
type Token struct {
	Type    TokenType
	Literal string
}

// Definizione dei token come costanti
const (
	// Tokeni speciali
	ILLEGAL = "ILLEGAL" // Token non riconosciuto
	EOF     = "EOF"     // Fine del file/input

	// Identificatori e letterali
	IDENT = "IDENT" // Identificatore, es: variabile
	INT   = "INT"   // Intero

	// Operatori
	ASSIGN       = "="
	PLUS         = "+"
	DASH         = "-"
	STAR         = "*"
	BANG         = "!"
	FORWARDSLASH = "/"
	LT           = "<"
	GT           = ">"

	// Operatori di confronto
	EQ     = "==" // uguale a
	NOT_EQ = "!=" // diverso da
	GT_EQ  = ">=" // maggiore o uguale a
	LT_EQ  = "<=" // minore o uguale a

	// Delimitatori
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Parole chiave
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// keywords è una mappa che associa identificatori testuali alle parole chiave corrispondenti
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
}

// LookupIdent verifica se un identificatore è una parola chiave o un identificatore generico
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
