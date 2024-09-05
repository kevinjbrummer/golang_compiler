package repl

import (
	"bufio"
	"fmt"
	"goblin/color"
	"goblin/evaluator"
	"goblin/lexer"
	"goblin/object"
	"goblin/parser"
	"io"
	"os/user"
)

const PROMPT = ">> "

const GOBLIN_LOGO = `             _     _ _       
            | |   | (_)      
  __ _  ___ | |__ | |_ _ __  
 / _  |/ _ \| '_ \| | | '_ \ 
| (_| | (_) | |_) | | | | | |
 \__, |\___/|_.__/|_|_|_| |_|
  __/ |                      
 |___/ 
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
  user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Fprint(out, color.ColorWrapper(color.GREEN, GOBLIN_LOGO + "\n"))
	fmt.Printf("Hello %s! This is the Goblin programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

	for {
		fmt.Fprint(out, color.ColorWrapper(color.GREEN, PROMPT))
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			return
		}

		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)

		evaluated := evaluator.Eval(expanded, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors[]string) {
	io.WriteString(out, color.ColorWrapper(color.RED, "Woops! We encountered some goblins!" + "\n"))
	io.WriteString(out, "parser errors:\n")
	
	for _, msg := range errors {
		io.WriteString(out, "\t" + msg + "\n")
	}
}