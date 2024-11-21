package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func addToken(tokenType string, text string, value ...string) {
	lexVal := "null"
	if len(value) > 0 {
		lexVal = value[0]
	}

	fmt.Printf("%s %s %s\n", tokenType, text, lexVal)
}

func indexAt(s, sep string, n int) int {
    idx := strings.Index(s[n:], sep)
    if idx > -1 {
        idx += n
    }
    return idx
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c == '_')
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

func identifier(s string) int {
	current_index := 0
	reservedWords := map[string]string{
		"true": "TRUE",
		"false": "FALSE",
		"nil": "NIL",
		"and": "AND",
		"or": "OR",
		"not": "NOT",
		"for": "FOR",
		"if": "IF",
		"else": "ELSE",
		"while": "WHILE",
		"var": "VAR",
		"this": "THIS",
		"super": "SUPER",
		"return": "RETURN",
		"print": "PRINT",
		"fun": "FUN",
		"class": "CLASS",
	}

	for current_index < len(s) && isAlphaNumeric(s[current_index]) {
		current_index++
	}
	
	text := s[:current_index]
	token, ok := reservedWords[text]
	
	if ok {
		addToken(token, text)
	} else {
		addToken("IDENTIFIER", s[:current_index])
	}

	return current_index
}

func scanner(fileContents []byte) bool {
	hasErrors := false
	lineNumber := 1
	for i := 0; i < len(fileContents); i++ {
		text := string(fileContents[i])
		switch fileContents[i] {
		case '(':
			addToken("LEFT_PAREN", text)
		case ')':
			addToken("RIGHT_PAREN", text)
		case '{':
			addToken("LEFT_BRACE", text)
		case '}':
			addToken("RIGHT_BRACE", text)
		case ',':
			addToken("COMMA", text)
		case '.':
			addToken("DOT", text)
		case '+':
			addToken("PLUS", text)
		case '-':
			addToken("MINUS", text)
		case '*':
			addToken("STAR", text)
		case ';':
			addToken("SEMICOLON", text)
		case '=':
			if i+1 < len(fileContents) && fileContents[i+1] == '=' {
				addToken("EQUAL_EQUAL", "==")
				i += 1
			} else {
				addToken("EQUAL", text)
			}
		case '!':
			if i+1 < len(fileContents) && fileContents[i+1] == '=' {
				addToken("BANG_EQUAL", "!=")
				i += 1
			} else {
				addToken("BANG", text)
			}
		case '<':
			if i+1 < len(fileContents) && fileContents[i+1] == '=' {
				addToken("LESS_EQUAL", "<=")
				i += 1
			} else {
				addToken("LESS", text)
			}
		case '>':
			if i+1 < len(fileContents) && fileContents[i+1] == '=' {
				addToken("GREATER_EQUAL", ">=")
				i += 1
			} else {
				addToken("GREATER", text)
			}
		case '/':
			if i+1 < len(fileContents) && fileContents[i+1] == '/' {
				newLineIndex := strings.Index(string(fileContents[i+1:]), "\n")
				if (newLineIndex >= 0) {
					i += strings.Index(string(fileContents[i:]), "\n")
					lineNumber += 1
				} else {
					i += len(fileContents)
				}
			} else {
				addToken("SLASH", text)
			}
		case ' ', '\t', '\r':
			//Ignore whitespace
		case '\n':
			lineNumber += 1
		case '"':
			endingQuoteIndex := indexAt(string(fileContents), "\"", i + 1)
			if endingQuoteIndex >= 0 {
				addToken("STRING", string(fileContents[i:endingQuoteIndex+1]), string(fileContents[i+1:endingQuoteIndex]))
				i = endingQuoteIndex
			} else {
				fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", lineNumber)
				hasErrors = true
				i += len(fileContents)
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			isFloat := false
			number := string(fileContents[i])
			for i+1 < len(fileContents) && ((fileContents[i+1] >= '0' && fileContents[i+1] <= '9') || fileContents[i+1] == '.') {
				if fileContents[i+1] == '.' {
					isFloat = true
				}
				number += string(fileContents[i+1])
				i += 1
			}
			if isFloat {
				result, _ := strconv.ParseFloat(number, 64)
				resultString := strconv.FormatFloat(result, 'f', -1, 64)
				if strings.Index(resultString, ".") == -1 {
					resultString += ".0"
				}

				addToken("NUMBER", number, resultString)
			} else {
				addToken("NUMBER", number, number + ".0")
			}
		default:
			if isAlpha(fileContents[i]) {
				i += identifier(string(fileContents[i:])) - 1
				break
			}
			fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", lineNumber, fileContents[i])
			hasErrors = true
		}
	}
	fmt.Println("EOF  null");
	
	return hasErrors
}