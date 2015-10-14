package lexer
import (
	"regexp"
	"strconv"
	"strings"
	"main/cst"
)

type Token struct {

}

const SOURCE_PATTERNS = createSourcePatterns()

func createSourcePatterns()map[*regexp.Regexp]func(string)Token{
	result := make(map[*regexp.Regexp]func(string)Token)

	for key, val := range map[string]func(string)Token{
		"[0-9]+": func(s string){
			result, _ := strconv.Atoi(s)
			return result
		}}{
		rexp, _ := regexp.Compile(key)
		result[rexp] = val
	}
	return result
}


// Takes a string and tokenizes it
func GetFlatTokenList(text string) []Token{
	var tokens []Token
	remaining_text := strings.TrimLeft(text, " ")
	last_iteration_text := remaining_text

	for len(remaining_text) > 0 {
		for pattern, converter := range SOURCE_PATTERNS {
			match := pattern.FindIndex([]byte(remaining_text))
			if match == nil || match[0] != 0 {
				continue
			}
			token := converter(remaining_text[match[0]:match[1]])
			append(tokens, token)
			// cut out the regex match
			remaining_text = remaining_text[match[1]:]
			break
		}
		if last_iteration_text == remaining_text{
			panic("loool... we couldn't find any match for the text")
		}
		last_iteration_text = remaining_text
	}
	return tokens
}

// Takes a slice of tokens, and returns a concrete syntax tree
func GetCstFromTokens(tokens []Token)[]cst.Cst{

}