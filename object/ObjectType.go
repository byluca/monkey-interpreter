// File: object/object.go

package object

import (
	"bytes"
	"fmt"
	"monkey-interpreter/ast"
	"strings"
)

// ObjectType è una stringa che usiamo per identificare i tipi di oggetti (es. "INTEGER").
type ObjectType string

// Object è l'interfaccia che ogni tipo di dato nel nostro linguaggio deve implementare.
// Questo ci permette di trattare numeri, booleani, funzioni, ecc., allo stesso modo.
type Object interface {
	Type() ObjectType // Restituisce il tipo dell'oggetto (es. INTEGER_OBJ).
	Inspect() string  // Restituisce una rappresentazione in stringa dell'oggetto, utile per il debug.
}

// Definiamo delle costanti per i tipi di oggetto, così evitiamo errori di battitura.
const (
	INTEGER_OBJ      = "INTEGER"      // Per i numeri interi
	BOOLEAN_OBJ      = "BOOLEAN"      // Per i valori vero/falso
	NULL_OBJ         = "NULL"         // Per il valore nullo `null`
	RETURN_VALUE_OBJ = "RETURN_VALUE" // Un tipo speciale per gestire le istruzioni `return`
	ERROR_OBJ        = "ERROR"        // Per gestire gli errori di runtime
	FUNCTION_OBJ     = "FUNCTION"     // Il nuovo tipo per rappresentare le funzioni
)

// --- DEFINIZIONI DELLE STRUCT PER OGNI TIPO DI OGGETTO ---

// Integer rappresenta un numero intero.
type Integer struct {
	Value int64 // Il valore numerico effettivo.
}

// Implementazione dell'interfaccia Object per Integer.
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

// Boolean rappresenta un valore booleano (true o false).
type Boolean struct {
	Value bool // Il valore booleano effettivo.
}

// Implementazione dell'interfaccia Object per Boolean.
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

// Null rappresenta l'assenza di un valore.
type Null struct{}

// Implementazione dell'interfaccia Object per Null.
func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

// ReturnValue è un "involucro" speciale per i valori restituiti da un'istruzione `return`.
// Serve a segnalare all'interprete di interrompere l'esecuzione di un blocco di codice.
type ReturnValue struct {
	Value Object // Il valore effettivo che viene restituito (es. un Integer o un Boolean).
}

// Implementazione dell'interfaccia Object per ReturnValue.
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() } // Mostra il valore interno.

// Error rappresenta un errore che si verifica durante l'esecuzione del codice.
type Error struct {
	Message string // Il messaggio di errore da mostrare.
}

// Implementazione dell'interfaccia Object per Error.
func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

// Function rappresenta una funzione creata dall'utente nel linguaggio Monkey.
type Function struct {
	// I parametri che la funzione accetta.
	Parameters []*ast.Identifier
	// Il blocco di codice che viene eseguito quando la funzione è chiamata.
	Body *ast.BlockStatement
	// L'ambiente (scope) in cui la funzione è stata definita.
	// Questo è il segreto delle chiusure (closures): la funzione "ricorda" le variabili
	// che erano disponibili al momento della sua creazione.
	Env *Environment
}

// Implementazione dell'interfaccia Object per Function.
func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	// Crea una rappresentazione testuale della funzione, es. "fn(x, y) { ... }".
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}
