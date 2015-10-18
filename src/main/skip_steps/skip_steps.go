package skip_steps
import (
	"main/scopes"
	"main/interpreter"
	"main/lexer"
	"main/parser"
	"main/types"
//	"fmt"
)

func InterpretText(text string, scope scopes.IScope) types.IType {
	tokens := lexer.GetFlatTokenList(text)
	cstree := lexer.GetCstFromTokens(tokens)
//	fmt.Println("vwh skpi_steps: IText", cstree)
	astree := parser.GetAstFromCst(&cstree)
	result := interpreter.InterpretAst(astree, scope)
	return result
}