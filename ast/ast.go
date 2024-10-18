package ast

import (
	"bytes"
	"monkey-interpreter/token"
)

// Node è un'interfaccia che definisce i metodi di base per i nodi dell'AST
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement è un'interfaccia che estende Node e aggiunge un metodo specifico per le dichiarazioni
type Statement interface {
	Node
	statementNode()
}

// Expression è un'interfaccia che estende Node e aggiunge un metodo specifico per le espressioni
type Expression interface {
	Node
	expressionNode()
}

// Program è il nodo radice di ogni AST e contiene una lista di dichiarazioni
type Program struct {
	Statements []Statement
}

// TokenLiteral restituisce il valore letterale del primo token nel programma
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement rappresenta una dichiarazione 'let'
type LetStatement struct {
	Token token.Token // il token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier rappresenta un identificatore nel programma
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// ReturnStatement rappresenta una dichiarazione 'return'
type ReturnStatement struct {
	Token       token.Token // il token 'return'
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// ExpressionStatement rappresenta una dichiarazione che è semplicemente un'espressione
type ExpressionStatement struct {
	Token      token.Token // il primo token dell'espressione
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// Metodi String per stampa e debug

// String restituisce l'intero programma sotto forma di stringa
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// Metodo String per LetStatement
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

// Metodo String per ReturnStatement
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// Metodo String per ExpressionStatement
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
