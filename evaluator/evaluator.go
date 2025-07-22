// File: evaluator/evaluator.go

package evaluator

import (
	"monkey-interpreter/ast"
	"monkey-interpreter/object"
)

// Definizioni globali per gli oggetti singleton, per ottimizzazione.
var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

// Eval è la funzione principale che attraversa l'AST.
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// --- STATEMENTS ---
	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	// --- EXPRESSIONS ---
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)

	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
	}

	return nil
}

// evalStatements valuta una slice di istruzioni.
func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range stmts {
		result = Eval(statement)
	}
	return result
}

// nativeBoolToBooleanObject è una funzione helper per restituire i singleton TRUE/FALSE.
func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

// evalPrefixExpression gestisce gli operatori prefissi.
func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return NULL // Più avanti diventerà un errore
	}
}

// evalBangOperatorExpression implementa la logica per l'operatore '!'.
func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE // !null è considerato true
	default:
		return FALSE // Tutti gli altri valori (es. numeri) sono "truthy"
	}
}

// evalMinusPrefixOperatorExpression implementa la logica per l'operatore '-'.
func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	// L'operatore di negazione funziona solo sugli interi
	if right.Type() != object.INTEGER_OBJ {
		return NULL // Più avanti diventerà un errore
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

// evalInfixExpression è il "centralino" per le operazioni infisse.
// Controlla i tipi degli operandi e delega alla funzione specifica.
func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	// Caso per operazioni tra INTERI
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)

	// Caso per operazioni di uguaglianza tra BOOLEANI (confronto di puntatori)
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)

	default:
		return NULL // Più avanti diventerà un errore
	}
}

// evalIntegerInfixExpression gestisce tutte le operazioni infisse tra due interi.
func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	// Operatori aritmetici
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}

	// Operatori di confronto
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)

	default:
		return NULL // Più avanti diventerà un errore
	}
}