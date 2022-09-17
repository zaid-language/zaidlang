package evaluator

import (
	"zaidlang.tech/x/zaid/ast"
	"zaidlang.tech/x/zaid/object"
	"zaidlang.tech/x/zaid/value"
)

func evaluateIf(node *ast.If, scope *object.Scope) object.Object {
	condition := Evaluate(node.Condition, scope)

	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Evaluate(node.Consequence, scope)
	} else if node.Alternative != nil {
		return Evaluate(node.Alternative, scope)
	}

	return value.NULL
}
