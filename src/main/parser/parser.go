package parser
import (
	"main/cst"
	"main/ast"
	"fmt"
)

func GetAstFromCst(cstree *cst.CSTNode)ast.IAstNode {
//	fmt.Println(">>>VWH parser.GAFC cstree.children[0].value.value",
//		cstree.Children[0].Value.Value)
	result := transformOneToOne(cstree)
	return result
}


func transformOneToOne(cstree *cst.CSTNode)ast.IAstNode {
	//hardcode the crap out of this... we must test that ints work!!!
	fmt.Println(">>>VWH parser.transform121", cstree.Value)
	//!! value is nil, but it has children...
	inter := *cstree.Value
	val, _ := inter.Value.(int)
	node := ast.Int{val} // We know this will be an Int...but!!!
	return node
}