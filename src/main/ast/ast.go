package ast
import (
	"main/types"
)

type AstType int
const(
	IntegerType AstType = iota
)

type Int struct {
	Value int
}
func (i Int)ToLisp()types.IType {
	return types.Int{Value: i.Value}
}
func (i Int) GetChildren()[]IAstNode {
	return nil
}

type IAstNode interface {
	ToLisp() types.IType
	GetChildren() []IAstNode
}
