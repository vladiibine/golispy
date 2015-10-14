package skip_steps
import (
	"main/scopes"
	"main/interpreter"
	"main/lexer"
)

func InterpretText(text string, scope *scopes.IScope){
	interpreter.InterpretAst(
		parser.GetAstFromCst(
			lexer.GetCstFromTokens(
				lexer.GetFlatTokenList(

				)
			)
		)
	)
}