package ast

import (
	"bytes"
	"monkey-interpreter/token"
)

// Node è l'interfaccia di base per tutti i nodi nell'AST (Albero Sintattico Astratto).
// Ogni nodo deve implementare i metodi TokenLiteral e String.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement rappresenta un nodo nell'AST che esprime una dichiarazione.
// Un Statement non restituisce un valore e deve implementare il metodo statementNode.
type Statement interface {
	Node
	statementNode()
}

// Expression rappresenta un nodo nell'AST che esprime un'espressione.
// Un Expression restituisce un valore e deve implementare il metodo expressionNode.
type Expression interface {
	Node
	expressionNode()
}

// Program è il nodo radice dell'AST. Contiene una lista di dichiarazioni.
// È il punto di partenza per l'intero programma analizzato.
type Program struct {
	Statements []Statement
}

// TokenLiteral restituisce il valore letterale del token associato al primo statement nel programma.
// Viene usato principalmente per debugging e test.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// String restituisce l'intero programma sotto forma di stringa leggibile.
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// LetStatement rappresenta una dichiarazione 'let', usata per assegnare valori a variabili.
type LetStatement struct {
	Token token.Token // il token 'let'
	Name  *Identifier // il nome della variabile
	Value Expression  // il valore assegnato alla variabile
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String restituisce la dichiarazione let in formato stringa.
func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// Identifier rappresenta un identificatore, cioè il nome di una variabile o funzione nel programma.
type Identifier struct {
	Token token.Token // il token dell'identificatore
	Value string      // il nome della variabile o funzione
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}

// ReturnStatement rappresenta una dichiarazione 'return', che restituisce un valore da una funzione.
type ReturnStatement struct {
	Token       token.Token // il token 'return'
	ReturnValue Expression  // il valore restituito
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// String restituisce la dichiarazione return in formato stringa.
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement rappresenta una dichiarazione che consiste in un'espressione.
// Ad esempio, "5 + 5;" è un'espressione che può essere usata come dichiarazione.
type ExpressionStatement struct {
	Token      token.Token // il primo token dell'espressione
	Expression Expression  // l'espressione contenuta nella dichiarazione
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// String restituisce l'espressione in formato stringa.
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// IntegerLiteral rappresenta un valore numerico intero nell'AST.
type IntegerLiteral struct {
	Token token.Token // il token che rappresenta l'intero
	Value int64       // il valore numerico
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// PrefixExpression rappresenta un'espressione con un operatore prefisso.
// Esempi: "-5", "!true". Un operatore prefisso viene applicato a un singolo operando.
type PrefixExpression struct {
	Token    token.Token // Il token prefisso, es. "!" o "-"
	Operator string      // L'operatore prefisso, es. "!" o "-"
	Right    Expression  // L'espressione a destra dell'operatore
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

// InfixExpression rappresenta un'espressione con un operatore infisso.
// Esempi: "5 + 5", "a == b". L'operatore è situato tra due espressioni.
type InfixExpression struct {
	Token    token.Token // Il token dell'operatore
	Left     Expression  // L'espressione a sinistra
	Operator string      // L'operatore infisso, es. "+", "-", "==", etc.
	Right    Expression  // L'espressione a destra
}

func (oe *InfixExpression) expressionNode()      {}
func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }
func (oe *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")
	return out.String()
}
