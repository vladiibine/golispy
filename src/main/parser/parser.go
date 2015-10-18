package parser
import (
	"main/cst"
	"main/ast"
//	"fmt"
	"fmt"
)

func GetAstFromCst(cstree *cst.CSTNode)ast.IAstNode {
//	fmt.Println(">>>VWH parser.GAFC cstree.children[0].value.value",
//		cstree.Children[0].Value.Value)
	result := transformOneToOne(cstree)
	return result
}


func transformOneToOne(cstree *cst.CSTNode)ast.IAstNode {
//	inter := *cstree.Value
//	val, _ := inter.Value.(int)
//	node := ast.Int{val} // We know this will be an Int...but!!!
//	return node
	var node_intf ast.IAstNode
	// keep the reference the current top node (the one corresponding to the
	// current cstree)
	var node_ref ast.IAstNode

	if cstree.Container{
		node_impl := ast.Pair{Tail:nil}
		node_intf = &node_impl

		if node_ref == nil {node_ref = node_intf} // save the reference
		child_idx := 0
		for child_idx < len(cstree.Children) {
			fmt.Println(">>>vwh here first!!!!")
			node_impl.Head = transformOneToOne(&cstree.Children[child_idx])
			fmt.Println(">>VWH parser.t121 ", node_impl.Head.ToLisp().String())
			fmt.Println(">>>vwh parser.t121 ", node_ref.ToLisp().String())
			new_node := ast.Pair{Tail: nil}
			node_intf = &node_impl
			fmt.Println(">>>VWH3: ",node_intf.ToLisp().String())
			fmt.Println(">>VWH &new_node", &new_node)
			node_impl.Tail = new_node
			fmt.Println(">>>VWH6 ", node_intf.ToLisp().String())
			fmt.Println(">>>VWH7 ", node_ref.ToLisp().String())
			// TODO: AHAAA! got it!
			// it seems the & operator gets the address of the VARIABLE,
			// NOT of the data. That means that if I change the value of
			// a variable (put something else in its place), the
			// variable that held a reference to it ALSO gets updated.
			// Totally unlike what Python does.
			// Well it just means that to save a reference to the first
			// node, we leave alone the variable of type ast.Pair or ast.Int
			//... and then have fun with the references.... weird, huh?
			node_impl = new_node
			fmt.Println(">>>VWH8 ", node_ref.ToLisp().String())
			child_idx++
		}
	} else{
		// hardcode... only ints for now.. ignore errors
		intval, _ := cstree.Value.Value.(int)
		node_intf = ast.Int{intval}
		fmt.Println(">>>VWH4 &node_intf, value", &node_intf, node_intf.ToLisp().String())
		if node_ref == nil {node_ref = node_intf} // save the ref
		fmt.Println(">>>VWH4 &node_ref, value", &node_ref, node_ref.ToLisp().String())
	}
//	return ast.Pair{
//		Head: ast.Pair{
//			Head: ast.Int{3},
//			Tail:ast.Int{4}},
//		Tail: ast.Int{7}}
	fmt.Println(">>VWH & node_ref", &node_ref)
	fmt.Println(">>VWH5 ", node_ref.ToLisp().String())
	return node_ref
}