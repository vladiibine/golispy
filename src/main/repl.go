package main
import (
	"fmt"
	"bufio"
	"os"
	"main/scopes"
	"main/skip_steps"
)

const GREETING = "----------<Welcome to the Golispy interpreter>----------"

const PS1 = "(GL)$"

func getUerInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

func Repl(){
	fmt.Println(GREETING, "\n")
	scope := new(scopes.Scope)
	for {
		fmt.Println(PS1)
		text := getUerInput()
		result := skip_steps.InterpretText(text, scope)
		fmt.Println(result.String())
	}
}
