package types
import "strconv"

type Int struct {
	Value int
}
func (i *Int)String() string{
	return strconv.Itoa(i.Value)
}
func (i *Int)IsContainer() bool{
	return false
}

// Not what you might think: It's a container of 1 element
type List struct {
	Value IType
}
func (l *List)String() string{
	return "(" + l.Value.String() + ")"
}
func (l *List)IsContainer() bool{
	return true
}

type Pair struct {
	Head IType
	Tail IType
}
func (p *Pair)String() string{
	return "( " + p.Head.String() + " " + p.Tail.String() + " )"
}
func (p *Pair)IsContainer() bool {
	return true
}

/**
	List: it's a container of 1 object

	Pair: Head|Tail
		repr: if the tail is a pair or a list (a container) " (<head> <tail>) "
				else: " (<head> . <tail>)"
	!!!TODO!!!!
		- (2 (3 (1))) != (2 3 1)
		- `cons` returns a pair, which IS a list, extended
		- `list` however, creates a new list with 2 elements, the the head, and the tail
 */

type IType interface {
	String() string
	IsContainer() bool
}

/**