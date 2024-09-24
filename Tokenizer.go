package main

import (
	"fmt"
	"strings"
)

type TokenType string // declares a custom type

// defines different token types
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

// Tokenize tokenizes an input string and breaks it down into tokens
func Tokenize(input string) []Token {
	var tokens []Token      // calls the Token struct and initializes it as a slice
	var currentToken string // to accumulate characters for the current token

	// iterates through each character in the input string
	for i, ch := range input {
		// skips hyphens to avoid splitting words like "word-word"
		if ch == '-' {
			// if there's an ongoing token, classify it before moving on to the next loop
			if len(currentToken) > 0 {
				tokens = append(tokens, classifyToken(currentToken))
				currentToken = "" // resets the accumulated token
			}
			continue
		}

		if ch == '.' && i > 0 && isDigit(rune(input[i-1])) && isDigit(rune(input[i+1])) {
			currentToken += string(ch)
			continue
		}

		// handles sentence ending punctuation
		if isSentenceEndPunctuation(ch) {
			if len(currentToken) > 0 {
				tokens = append(tokens, classifyToken(currentToken))
				currentToken = ""
			}
			// adds the punctuation as a separate token
			tokens = append(tokens, Token{Value: string(ch), Type: punctuation})
			// adds a new token to indicate end of line
			if i == len(input)-1 || input[i+1] == ' ' || input[i+1] == '\n' {
				tokens = append(tokens, Token{Value: "\\n", Type: endOfLine})
			}
			continue
		}

		// if it's the last character, classify the final token
		if i == len(input)-1 {
			currentToken += string(ch)
			tokens = append(tokens, classifyToken(currentToken))
			continue

		}

		// builds the token to be classified before reaching a delimiter
		currentToken += string(ch)
	}

	return tokens

}

// classifyToken identifies the type of token based on its contents
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
	// default to word if none apply
	return Token{Value: token, Type: word}
}

func isNumber(token string) bool {
	dotSeen := false
	for _, ch := range token {
		if ch == '.' {
			if dotSeen {
				return false
			}
			dotSeen = true
			continue
		}
		if !isDigit(ch) {
			return false
		}
	}
	return true
}

func isWord(token string) bool {
	allowedPunctuations := "'@"
	for _, ch := range token {
		if isLetter(ch) || containsRune(allowedPunctuations, ch) {
			continue
		}
		return false
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
	punctuationMarks := "!\"#$%&'()+.,/:;<=>?@[\\]^_`{|}~"
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
