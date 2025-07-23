// File: object/environment.go
package object

// NewEnvironment crea un nuovo ambiente vuoto.
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment crea un nuovo ambiente "figlio" che è racchiuso
// da un ambiente "genitore" (outer).
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer // Collega questo nuovo ambiente a quello esterno.
	return env
}

// Environment tiene traccia delle variabili (identificatori e i loro valori).
type Environment struct {
	store map[string]Object
	outer *Environment // Puntatore all'ambiente esterno (per le chiusure).
}

// Get cerca una variabile. Se non la trova qui, la cerca nell'ambiente esterno.
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		// Se non trovata, delega la ricerca all'ambiente "genitore".
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set aggiunge o aggiorna una variabile nell'ambiente corrente.
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
