// File: repl/repl.go

package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey-interpreter/evaluator"
	"monkey-interpreter/lexer"
	"monkey-interpreter/object" // Assicurati che questo import sia presente
	"monkey-interpreter/parser"
)

const PROMPT = ">> "

const MONKEY_FACE = `
            .--.  .-"     "-.  .--.
           / .. \/  .-. .-.  \/ .. \
          | |  '|  /   Y   \  |'  | |
          | \   \  \ 0 | 0 /   /   / |
           \ '- ,\.-"""""""-./, -' /
            ''-' /_   ^ ^   _\ '-''
               |  \._   _./  |
               \   \ '~' /   /
                '._ '-=-' _.'
                   '-----'
`

// Start avvia il ciclo Read-Eval-Print-Loop.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	// Crea un singolo Environment che verrà riutilizzato per tutta la sessione del REPL.
	// Questo permette di mantenere lo stato (le variabili) tra un input e l'altro.
	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		// Passa sia l'AST (program) che l'ambiente (env) all'evaluator.
		// L'evaluator userà 'env' per leggere e scrivere le variabili.
		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

// printParserErrors stampa gli errori del parser in un formato carino.
func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Ops! Abbiamo incontrato un problema con la scimmia!\n")
	io.WriteString(out, "Errori del parser:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
