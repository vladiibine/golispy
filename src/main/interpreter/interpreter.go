package interpreter
import (
	"main/scopes"
	"main/ast"
	"main/types"
)

func InterpretAst(astree ast.IAstNode, scope scopes.IScope) types.IType {
	instance := astree.ToLisp()
	return instance
}