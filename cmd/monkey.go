package main

import (
	"bufio"
	"fmt"
	"os"

	"monkey/lexer"
	"monkey/parser"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("welcome to the monkey programming language\n")

	for {
		fmt.Print(":> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		l := lexer.New(input)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) == 0 {
			fmt.Println(program.ToString())
		} else {
			for _, e := range p.Errors() {
				err := fmt.Errorf(e.Info())
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
