package statements

import "lang/ast"

type Statement interface {
	ast.Node
	statementNode()
}
