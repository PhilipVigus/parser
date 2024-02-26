package expressions

import "lang/ast"

type Expression interface {
	ast.Node
	expressionNode()
}
