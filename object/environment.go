// File: object/environment.go

package object

// NewEnvironment crea un nuovo Environment vuoto.
func NewEnvironment() *Environment {
	store := make(map[string]Object)
	return &Environment{store: store}
}

// Environment tiene traccia delle variabili e dei loro valori.
type Environment struct {
	store map[string]Object
}

// Get recupera un valore dall'Environment dato un nome.
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set inserisce un nuovo valore nell'Environment, associandolo a un nome.
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
