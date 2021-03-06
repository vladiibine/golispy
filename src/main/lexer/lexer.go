package lexer
import (
	"regexp"
	"strconv"
	"strings"
	"main/cst"
//	"fmt"
)

type TokenType int
const (
	Increment TokenType = iota
	Decrement
	Integer
)


type Token struct {
	Type TokenType
	Value interface{} // dunno if I should defer the conversion, but I do
}

var SOURCE_PATTERNS = createSourcePatterns()

func createSourcePatterns()map[*regexp.Regexp]func(string)Token{
	result := make(map[*regexp.Regexp]func(string)Token)

	for key, val := range map[string]func(string)Token{
		// Parse ints
		"[0-9]+": func(s string)Token{
			result, _ := strconv.Atoi(s) // To hell with error handling!
			token := Token{Type: Integer, Value:result}
			return token
		},
		"\\(": func(s string)Token{
			return Token{Type: Increment, Value: nil}
		}} {
			rexp, _ := regexp.Compile(key)
			result[rexp] = val
	}
	return result
}


// Takes a string and tokenizes it
func GetFlatTokenList(text string) []Token{
	var tokens []Token
	remaining_text := strings.Trim(text, " \n\t\r\u000a")
	last_iteration_text := remaining_text

	for len(remaining_text) > 0 {
		for pattern, converter := range SOURCE_PATTERNS {
			match := pattern.FindIndex([]byte(remaining_text))
			if match == nil || match[0] != 0 {
				continue
			}
			token := converter(remaining_text[match[0]:match[1]])
//			fmt.Println("appended ", token)
			tokens = append(tokens, token)
			// cut out the regex match
			remaining_text = remaining_text[match[1]:]
			break
		}
//		println(remaining_text == "\u000a")
		if last_iteration_text == remaining_text{
			panic("loool... we couldn't find any match for the text")
		}
		last_iteration_text = strings.Trim(remaining_text, " \n\t\r\u000a")
	}
	return tokens
}

// Takes a slice of tokens, and returns a concrete syntax tree
func GetCstFromTokens(tokens []Token) cst.CSTNode {
	var current_children []cst.CSTNode

	cstree := cst.CSTNode{Children: current_children, Container:true}

	for _, token := range tokens{
		if token.Type == Integer {
			current_value := cst.CSTElement{Type:cst.Integer, Value:token.Value}
			if cstree.Container {
				new_child := cst.CSTNode{Value: &current_value}
				cstree.Children = append(cstree.Children, new_child)
			} else {
				cstree.Value = &current_value
			}
		} else if token.Type == Increment {
			// Add a new CSTNode element in the children list, and point to it
			current_child := cst.CSTNode{Parent: &cstree}
			cstree.Children = append(cstree.Children, current_child)
			cstree = current_child
		} else if token.Type == Decrement {
			// Point to the parent of the current node
			cstree = *cstree.Parent
		}
	}
	return cstree
}