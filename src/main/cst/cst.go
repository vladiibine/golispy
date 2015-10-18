package cst

type CstType int
const (
	Integer CstType = iota
)

type CSTElement struct {
	Type CstType
	Value interface{} // decide later the interface of the nodes
}

//
type CSTNode struct {
	Value *CSTElement
	Children []CSTNode
	Parent *CSTNode
}
