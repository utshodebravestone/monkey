package main

import (
	"bufio"
	"fmt"
	"os"

	"monkey/lexer"
	"monkey/token"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("welcome to the monkey programming language\n")
	for {
		fmt.Print("::> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		l := lexer.New(input)
		for t := l.NextToken(); t.Kind != token.EOF; t = l.NextToken() {
			fmt.Println(t.ToString())
		}
	}
}
