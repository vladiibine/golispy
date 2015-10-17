package ast
import (
	"main/types"
)

type AstType int
const(
	Integer AstType = iota
)

type Int struct {
	Value int
}
func (i Int)ToLisp()types.IType {

}

type AstNode interface {
	ToLisp() types.IType
	GetChildren() []AstNode
}
