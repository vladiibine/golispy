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
func (i Int)GetChildren()[]IAstNode {
	return nil
}
func (i Int)IsContainer() bool{
	return false
}

type Pair struct {
	Head IAstNode
	Tail IAstNode
}
func (p Pair)ToLisp()types.IType{
	var head, tail types.IType

	if p.Head != nil { head = p.Head.ToLisp() }
	if p.Tail != nil { tail = p.Tail.ToLisp() }

	return types.Pair{Head: head, Tail: tail}
}
func (p Pair)GetChildren()[]IAstNode{
	return []IAstNode{p.Head, p.Tail}
}
func (p Pair)IsContainer() bool{
	return true
}

type IAstNode interface {
	ToLisp() types.IType
	GetChildren() []IAstNode
	IsContainer() bool
}
