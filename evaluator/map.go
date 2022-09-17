package evaluator

import (
	"zaidlang.tech/x/zaid/ast"
	"zaidlang.tech/x/zaid/object"
)

func evaluateMap(node *ast.Map, scope *object.Scope) object.Object {
	pairs := make(map[object.MapKey]object.MapPair)

	for keyNode, valueNode := range node.Pairs {
		key := Evaluate(keyNode, scope)

		if isError(key) {
			return key
		}

		mapKey, ok := key.(object.Mappable)

		if !ok {
			return newError("%d:%d:%s: runtime error: unusable as map key: %s", node.Token.Line, node.Token.Column, node.Token.File, key.Type())
		}

		value := Evaluate(valueNode, scope)

		if isError(value) {
			return value
		}

		hashed := mapKey.MapKey()

		pairs[hashed] = object.MapPair{Key: key, Value: value}
	}

	return &object.Map{Pairs: pairs}
}
