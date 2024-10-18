package parser

import (
	"fmt"
	"monkey-interpreter/ast"
	"monkey-interpreter/lexer"
	"monkey-interpreter/token"
	"strconv"
)

// Definisce le precedenze degli operatori nel linguaggio Monkey
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

// Mappa che associa i token degli operatori con la loro precedenza
var precedences = map[token.TokenType]int{
	token.EQ:           EQUALS,
	token.NOT_EQ:       EQUALS,
	token.LT:           LESSGREATER,
	token.GT:           LESSGREATER,
	token.PLUS:         SUM,
	token.DASH:         SUM,
	token.FORWARDSLASH: PRODUCT,
	token.STAR:         PRODUCT,
}

// Parser è la struttura che rappresenta il parser del linguaggio Monkey.
// È responsabile di analizzare il flusso di token generati dal lexer e costruire l'AST.
type Parser struct {
	l              *lexer.Lexer
	errors         []string
	curToken       token.Token
	peekToken      token.Token
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// New crea un nuovo parser e registra le funzioni di parsing per i token prefissi e infissi.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:              l,
		errors:         []string{},
		prefixParseFns: make(map[token.TokenType]prefixParseFn),
		infixParseFns:  make(map[token.TokenType]infixParseFn),
	}

	// Registrazione dei parser per i prefissi (es. identificatori, interi, operatori prefissi)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.DASH, p.parsePrefixExpression)

	// Registrazione dei parser per gli operatori infissi (es. +, -, *, /)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.DASH, p.parseInfixExpression)
	p.registerInfix(token.FORWARDSLASH, p.parseInfixExpression)
	p.registerInfix(token.STAR, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)

	// Avanza per impostare curToken e peekToken
	p.nextToken()
	p.nextToken()

	return p
}

// parseIdentifier crea un nodo AST per un identificatore.
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

// Errors restituisce gli errori incontrati durante il parsing.
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError aggiunge un errore alla lista se il token successivo non è quello atteso.
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// nextToken avanza il parser al prossimo token.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram crea un AST per il programma analizzando una lista di dichiarazioni.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// Continua a parsare finché non si raggiunge la fine dell'input
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// parseLetStatement analizza una dichiarazione let, ad esempio "let x = 5;".
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// Controlla se il prossimo token è un identificatore
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Controlla se il prossimo token è il segno di assegnazione "="
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// Salta fino al punto e virgola
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// curTokenIs verifica se il token corrente corrisponde a un determinato tipo.
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs verifica se il token successivo è di un certo tipo.
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek verifica se il prossimo token è quello atteso e avanza il parser.
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// parseStatement determina quale tipo di dichiarazione si sta analizzando.
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// parseExpressionStatement analizza una dichiarazione espressione.
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST) // Analizza l'espressione con la precedenza più bassa

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// parseReturnStatement analizza una dichiarazione return, ad esempio "return 5;".
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// Avanza fino al punto e virgola
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// registerPrefix registra una funzione di parsing per un operatore prefisso.
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// registerInfix registra una funzione di parsing per un operatore infisso.
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// noPrefixParseFnError registra un errore quando non esiste una funzione di parsing per un prefisso.
func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

// parseExpression gestisce il parsing delle espressioni, scegliendo tra operatori prefissi o infissi.
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()
		leftExp = infix(leftExp)
	}

	return leftExp
}

// parseIntegerLiteral gestisce il parsing di un valore intero.
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

// parsePrefixExpression gestisce il parsing di un operatore prefisso (es. "!5", "-10").
func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)

	return expression
}

// peekPrecedence restituisce la precedenza del token successivo.
func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

// curPrecedence restituisce la precedenza del token corrente.
func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

// parseInfixExpression gestisce il parsing di un'espressione con operatore infisso (es. "5 + 5").
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}
