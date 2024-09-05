package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TokenType string

const (
	word         TokenType = "Word"
	number       TokenType = "Number"
	punctuation  TokenType = "Punctuation"
	endOfLine    TokenType = "End Of Line"
	alphanumeric TokenType = "Alphanumeric"
)

type Token struct {
	Value string
	Type  TokenType
}

func Tokenize(input string) []Token {
	var tokens []Token
	var currentToken string

	for i, ch := range input {
		if ch == '-' {
			if len(currentToken) > 0 {
				tokens = append(tokens, classifyToken(currentToken))
				currentToken = ""
			}
			continue
		}

		if isSentenceEndPunctuation(ch) {
			if len(currentToken) > 0 {
				tokens = append(tokens, classifyToken(currentToken))
				currentToken = ""
			}
			tokens = append(tokens, Token{Value: string(ch), Type: punctuation})
			tokens = append(tokens, Token{Value: "\\n", Type: endOfLine})
			continue
		}

		if i == len(input)-1 {
			currentToken += string(ch)
			tokens = append(tokens, classifyToken(currentToken))
			continue

		}
		currentToken += string(ch)
	}

	return tokens

}

func classifyToken(token string) Token {
	if isNumber(token) {
		return Token{Value: token, Type: number}
	} else if isWord(token) {
		return Token{Value: token, Type: word}
	} else if isAlphanumeric(token) {
		return Token{Value: token, Type: alphanumeric}
	} else if isPunctuation(token) {
		return Token{Value: token, Type: punctuation}
	}
	return Token{Value: token, Type: endOfLine}
}

func isNumber(token string) bool {
	for _, ch := range token {
		if !isDigit(ch) {
			return false
		}
	}
	return true
}

func isWord(token string) bool {
	for _, ch := range token {
		if !isLetter(ch) {
			return false
		}
	}
	return true
}

func isPunctuation(token string) bool {
	for _, ch := range token {
		if isAsciiPunctuation(ch) {
			return true
		}
	}
	return false
}

func isAlphanumeric(token string) bool {
	for _, ch := range token {
		if isLetter(ch) {
			continue
		}
		if isDigit(ch) {
			continue
		}
		return false
	}
	return true
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isSentenceEndPunctuation(ch rune) bool {
	return ch == '.' || ch == '!' || ch == '?'
}

func isAsciiPunctuation(ch rune) bool {
	punctuationMarks := "!\"#$%&'()+,./:;<=>?@[\\]^_`{|}~"
	return containsRune(punctuationMarks, ch)
}

func containsRune(str string, ch rune) bool {
	for _, r := range str {
		if r == ch {
			return true
		}
	}
	return false
}

func printPhase1Output(tokens []Token) {
	fmt.Println("Phase 1 Output:")
	for _, token := range tokens {
		fmt.Printf(`Token: "%s" - Type: %s`+"\n", token.Value, token.Type)
	}
	fmt.Println("=============")
}

func printPhase2Output(tokens []Token) {
	fmt.Println("Phase 2 Output:")
	for _, token := range tokens {
		if token.Type == word || token.Type == number || token.Type == alphanumeric || token.Type == endOfLine || token.Type == punctuation {
			fmt.Printf(`Token: "%s" -> `, token.Value)
			for i, ch := range token.Value {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Printf("`%c`", ch)
			}
			fmt.Println()
		}
	}
}

func main() {
	fmt.Print("Please enter a sentence: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	for _, ch := range input {
		if ch == ' ' {
			fmt.Println("Error: Input can only contain hyphens as delimiters.")
			return
		}
	}

	tokens := Tokenize(input)
	printPhase1Output(tokens)
	printPhase2Output(tokens)

}
