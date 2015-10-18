package types
import "strconv"

type Int struct {
	Value int
}
func (i Int)String() string{
	return strconv.Itoa(i.Value)
}

type NILType struct {
}
func (n *NILType)String() string {
	return "NILType"
}


type Pair struct {
	Head IType
	Tail IType // restriction: this must return only false
}
func (p *Pair)String() string{
	// If Tail is neither NIL nor a pair, we put the '.'
	return "( " + p.Head.String() + " . " + p.Tail.String() + " )"
}
func (p *Pair)Container(){ } // marks this is an IPair


/**
Ok... so there's no qualitative difference between a CONS pair and a LIST
	...Cons pairs hold just data in their CDR
	 	lists hold either references to other cons, OR nil

	So (list 1) is the same type and structure as (cons 1 Nil)
	And (list 1 2) is (cons 1 (cons 2 NIL))

	So therefore the `list` function must do funny stuff, while `cons` simply
	acts natural
 */

type IType interface {
	String() string
}

type IPair interface {
	IType
	Container() // marker method... just to have some static checking
}