package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if program == nil {
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")

		io.WriteString(out, "=>")
		result := evaluator.Eval(program)
		if result != nil {
			io.WriteString(out, result.Inspect())
			io.WriteString(out, "\n")
		} else {
			io.WriteString(out, "NIL\n")
		}
	}
}
