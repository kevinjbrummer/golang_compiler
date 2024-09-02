package repl

import (
	"bufio"
	"fmt"
	"goblin/lexer"
	"goblin/parser"
	"io"
	"os/user"
)

var COLORS = map[string]string {
	"RESET": "\033[0m",
	"RED": "\033[31m",
	"GREEN": "\033[32;1m",
	"YELLOW": "\033[33m",
	"BLUE": "\033[34m",
	"MAGENTA": "\033[35m",
	"CYAN": "\033[36m",
	"GRAY": "\033[37m",
	"WHITE": "\033[97m",
}

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

	fmt.Fprint(out, COLORS["GREEN"] + GOBLIN_LOGO + COLORS["RESET"] + "\n")
	fmt.Printf("Hello %s! This is the Goblin programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	for {
		fmt.Fprint(out, COLORS["GREEN"] + PROMPT + COLORS["RESET"])
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors[]string) {
	io.WriteString(out, COLORS["RED"] + "Woops! We encountered some goblins!" + COLORS["RESET"] + "\n")
	io.WriteString(out, "parser errors:\n")
	
	for _, msg := range errors {
		io.WriteString(out, "\t" + msg + "\n")
	}
}