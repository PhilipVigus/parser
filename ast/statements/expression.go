package statements

import "lang/lexer/token"
import "lang/ast/expressions"

type ExpressionStatement struct {
	Token      token.Token
	Expression expressions.Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenValue() string {
	return es.Token.Value
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
