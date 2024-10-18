package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey-interpreter/lexer"
	"monkey-interpreter/token"
)

const prompt = ">> "

// Start avvia il ciclo REPL che legge input dall'utente e restituisce i token generati
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		// Mostra il prompt e attende l'input
		fmt.Fprintf(out, prompt)
		scanned := scanner.Scan()

		// Se l'input termina, esci dal ciclo
		if !scanned {
			return
		}

		// Leggi la linea di input
		line := scanner.Text()
		lexer := lexer.New(line)

		// Scorri i token finché non raggiungi EOF
		for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
