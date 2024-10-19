package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey-interpreter/lexer"
	"monkey-interpreter/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)        // Creiamo un nuovo lexer
		p := parser.New(l)          // Creiamo un nuovo parser basato sul lexer
		program := p.ParseProgram() // Analizziamo l'input
		if len(p.Errors()) != 0 {   // Controlliamo se ci sono errori
			printParserErrors(out, p.Errors())
			continue
		}
		io.WriteString(out, program.String()) // Stampiamo l'AST come stringa
		io.WriteString(out, "\n")
	}
}

const MONKEY_FACE = `
        .--.  .-"     "-.  .--.
       / .. \/  .-. .-.  \/ .. \
      | |  '|  /   Y   \  |'  | |
      | \   \  \0|0/   /   / |
       \ '- ,\.-"""""""-./, -' /
        ''-' /_   ^ ^   _\ '-'' 
           |  \._   _./  |
           \   \ '~' /   /
            '._ '-=-' _.'
               '-----'
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE) // Stampa la faccia di scimmia
	io.WriteString(out, "Ops! Abbiamo incontrato un problema con la scimmia!\n")
	io.WriteString(out, "Errori del parser:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
