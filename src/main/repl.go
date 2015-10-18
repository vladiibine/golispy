package main
import (
	"fmt"
	"bufio"
	"os"
	"main/scopes"
	"main/skip_steps"
)

const GREETING = "----------<Welcome to the Golispy interpreter>----------"

const PS1 = "(GL)$ "
const PS2 = "..... "
const PS3 = "(GL)> "

func getUerInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

func Repl(){
	fmt.Println(GREETING, "\n")
	scope := new(scopes.Scope)
	for {
		fmt.Print(PS1)
		text := getUerInput()
		result := skip_steps.InterpretText(text, scope)
		fmt.Println(PS3, result.String(), "\n")
	}
}
