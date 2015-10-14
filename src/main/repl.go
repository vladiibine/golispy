package main
import (
	"fmt"
	"bufio"
	"os"
	"main/scopes"
)

const GREETING = "----------<Welcome to the Golispy interpreter>----------"

const PS1 = "(GL)$"

func getUerInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	_ = text
	return "nothing implemented yet, lol"
}

func Repl(){
	fmt.Println(GREETING, "\n")
	scope := new(scopes.Scope)
	for {
		fmt.Println(PS1)
		text := getUerInput()

		fmt.Println(text)
	}
}
