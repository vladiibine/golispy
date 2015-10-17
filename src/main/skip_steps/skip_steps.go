package skip_steps
import (
	"main/scopes"
	"main/interpreter"
	"main/lexer"
	"main/parser"
)

func InterpretText(text string, scope *scopes.IScope) {
	tokens := lexer.GetFlatTokenList(text)
	cstree := lexer.GetCstFromTokens(&tokens)
	astree := parser.GetAstFromCst(cstree)
	interpreter.InterpretAst(astree, scope)
}