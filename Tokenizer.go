package main

import (
	"fmt"
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
	return Token{Value: token, Type: word}
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

func output(tokens []Token) string {
	var result strings.Builder
	result.WriteString("Phase 1 Output:\n")
	for _, token := range tokens {
		result.WriteString(fmt.Sprintf(`Token: "%s" - Type: %s`+"\n", token.Value, token.Type))
	}
	result.WriteString("\n===========================================\n\n")

	result.WriteString("Phase 2 Output:\n")
	for _, token := range tokens {
		if token.Type == word || token.Type == number || token.Type == alphanumeric || token.Type == punctuation || token.Type == endOfLine {
			result.WriteString(fmt.Sprintf(`Token "%s" -> `, token.Value))
			for i, ch := range token.Value {
				if i > 0 {
					result.WriteString(", ")
				}
				result.WriteString(fmt.Sprintf("`%c`", ch))
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
