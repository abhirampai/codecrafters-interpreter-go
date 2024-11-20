package main

import (
	"fmt"
	"os"
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
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	
	if len(fileContents) > 0 {
		for i := 0; i < len(fileContents); i++ {
			text := string(fileContents[i])
			switch fileContents[i] {
				case '(':
					addToken("LEFT_PAREN", text);
				case ')':
					addToken("RIGHT_PAREN", text);
				case '{':
					addToken("LEFT_BRACE", text);
				case '}':
					addToken("RIGHT_BRACE", text);	
				case ',':
					addToken("COMMA", ",");
				case '.':
					addToken("DOT", ".");
				case '+':
					addToken("PLUS", "+");
				case '-':
					addToken("MINUS", "-");
				case '*':
					addToken("STAR", "*");
				case ';':
					addToken("SEMICOLON", ";");
				default:
					fmt.Printf("%c Unexpected character.c\n", fileContents[i])
			}
		}
		fmt.Println("EOF  null")
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
