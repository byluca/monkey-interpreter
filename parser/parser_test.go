package parser

import (
	"fmt"
	"monkey-interpreter/ast"
	"monkey-interpreter/lexer"
	"testing"
)

// TestLetStatements verifica che il parser analizzi correttamente le dichiarazioni let.
// Controlla che il parser riesca a costruire un AST corretto per dichiarazioni di tipo let.
func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

	l := lexer.New(input)       // Crea un nuovo lexer con l'input fornito
	p := New(l)                 // Crea un nuovo parser
	program := p.ParseProgram() // Analizza il programma
	checkParserErrors(t, p)     // Verifica se ci sono errori nel parser

	// Se il programma è nil, il test fallisce
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	expectedStatements := 3 // Numero atteso di dichiarazioni let
	if len(program.Statements) != expectedStatements {
		t.Fatalf("program.Statements non contiene %d dichiarazioni. ottenuto=%d",
			expectedStatements, len(program.Statements))
	}

	// Array di identificatori attesi
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	// Verifica che ogni dichiarazione let contenga l'identificatore atteso
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

// testLetStatement verifica che una dichiarazione let abbia il nome e il formato corretti.
func testLetStatement(t *testing.T, s ast.Statement, expectedName string) bool {
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s non è *ast.LetStatement. ottenuto=%T", s)
		return false
	}

	// Controlla che il literal del token sia "let"
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral non è 'let'. ottenuto=%q", s.TokenLiteral())
		return false
	}

	// Controlla che il nome della variabile corrisponda a quello atteso
	if letStmt.Name.Value != expectedName {
		t.Errorf("letStmt.Name.Value non è '%s'. ottenuto=%s", expectedName, letStmt.Name.Value)
		return false
	}

	// Controlla che il literal del nome della variabile corrisponda a quello atteso
	if letStmt.Name.TokenLiteral() != expectedName {
		t.Errorf("letStmt.Name.TokenLiteral() non è '%s'. ottenuto=%s", expectedName, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

// checkParserErrors verifica se ci sono errori durante il parsing e li stampa.
func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("Il parser ha %d errori", len(errors))
	for _, msg := range errors {
		t.Errorf("errore del parser: %q", msg)
	}
	t.FailNow() // Termina il test immediatamente se ci sono errori
}

// TestReturnStatements verifica la corretta analisi delle dichiarazioni return.
func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return add(5,10);
`

	l := lexer.New(input)       // Crea un nuovo lexer con l'input fornito
	p := New(l)                 // Crea un nuovo parser
	program := p.ParseProgram() // Analizza il programma
	checkParserErrors(t, p)     // Verifica se ci sono errori nel parser

	expectedStatements := 3 // Numero atteso di dichiarazioni return
	if len(program.Statements) != expectedStatements {
		t.Fatalf("program.Statements non contiene %d dichiarazioni. ottenuto=%d",
			expectedStatements, len(program.Statements))
	}

	// Verifica che ogni dichiarazione return sia corretta
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt non è *ast.ReturnStatement. ottenuto=%T", stmt)
			continue
		}

		// Controlla che il literal del token sia "return"
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral non è 'return', ottenuto %q", returnStmt.TokenLiteral())
		}
	}
}

// TestIdentifierExpression verifica che il parser possa analizzare espressioni contenenti identificatori.
func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar", ident.TokenLiteral())
	}
}

// TestIntegerLiteralExpression verifica la corretta analisi delle espressioni letterali intere.
func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5", literal.TokenLiteral())
	}
}

// testIntegerLiteral verifica se un nodo dell'AST rappresenta correttamente un intero.
func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value, integ.TokenLiteral())
		return false
	}

	return true
}

// TestParsingPrefixExpressions verifica che il parser possa gestire espressioni prefisse come "!5" o "-15".
func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
}

// TestParsingInfixExpressions verifica che il parser possa analizzare correttamente espressioni infisse come "5 + 5".
func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("stmt.Expression is not ast.InfixExpression. got=%T", stmt.Expression)
		}

		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			return
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
			return
		}
	}
}

// TestOperatorPrecedenceParsing verifica che il parser rispetti le precedenze degli operatori nelle espressioni complesse.
func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"-a * b", "((-a) * b)"},
		{"!-a", "(!(-a))"},
		{"a + b + c", "((a + b) + c)"},
		{"a + b - c", "((a + b) - c)"},
		{"a * b * c", "((a * b) * c)"},
		{"a * b / c", "((a * b) / c)"},
		{"a + b / c", "(a + (b / c))"},
		{"a + b * c + d / e - f", "(((a + (b * c)) + (d / e)) - f)"},
		{"3 + 4; -5 * 5", "(3 + 4)((-5) * 5)"},
		{"5 > 4 == 3 < 4", "((5 > 4) == (3 < 4))"},
		{"5 < 4 != 3 > 4", "((5 < 4) != (3 > 4))"},
		{"3 + 4 * 5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		actual := program.String()

		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}
