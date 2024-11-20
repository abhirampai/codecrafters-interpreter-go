package main

import (
	"fmt"
	"os"
	"strings"
)

func addToken(tokenType string, text string) {
	fmt.Printf("%s %s null\n", tokenType, text)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	has_errors := false
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
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
				addToken("COMMA", ",")
			case '.':
				addToken("DOT", ".")
			case '+':
				addToken("PLUS", "+")
			case '-':
				addToken("MINUS", "-")
			case '*':
				addToken("STAR", "*")
			case ';':
				addToken("SEMICOLON", ";")
			case '=':
				if i+1 < len(fileContents) && fileContents[i+1] == '=' {
					addToken("EQUAL_EQUAL", "==")
					i += 1
				} else {
					addToken("EQUAL", "=")
				}
			case '!':
				if i+1 < len(fileContents) && fileContents[i+1] == '=' {
					addToken("BANG_EQUAL", "!=")
					i += 1
				} else {
					addToken("BANG", "!")
				}
			case '<':
				if i+1 < len(fileContents) && fileContents[i+1] == '=' {
					addToken("LESS_EQUAL", "<=")
					i += 1
				} else {
					addToken("LESS", "<")
				}
			case '>':
				if i+1 < len(fileContents) && fileContents[i+1] == '=' {
					addToken("GREATER_EQUAL", ">=")
					i += 1
				} else {
					addToken("GREATER", ">")
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
					addToken("SLASH", "/")
				}
			case ' ', '\t', '\r':
				//Ignore whitespace
			case '\n':
				lineNumber += 1
			default:
				fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", lineNumber, fileContents[i])
				has_errors = true
			}
		}
		fmt.Println("EOF  null")
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}

	if has_errors {
		os.Exit(65)
	}
}
