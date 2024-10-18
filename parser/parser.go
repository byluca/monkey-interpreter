package parser

import (
	"fmt"
	"monkey-interpreter/ast"
	"monkey-interpreter/lexer"
	"monkey-interpreter/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

// New crea un nuovo parser con il lexer passato come parametro
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Legge due token per impostare curToken e peekToken
	p.nextToken()
	p.nextToken()
	return p
}

// Errors restituisce gli errori raccolti durante il parsing
func (p *Parser) Errors() []string {
	return p.errors
}

// aggiunge un messaggio di errore quando il prossimo token non è quello atteso
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// Avanza al prossimo token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram crea un ast.Program e popola la lista Statements con le istruzioni parse
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// parseLetStatement analizza una dichiarazione let
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// Controlla se il prossimo token è un identificatore
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Controlla se il prossimo token è il segno di assegnazione
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// Salta le espressioni fino al punto e virgola
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// curTokenIs controlla se il token corrente è di un certo tipo
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs controlla se il prossimo token è di un certo tipo
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek controlla se il prossimo token è quello atteso, e avanza
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// parseStatement analizza la dichiarazione corrente
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

// parseReturnStatement analizza una dichiarazione return
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	// Avanza al prossimo token
	p.nextToken()

	// Salta le espressioni fino al punto e virgola
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
