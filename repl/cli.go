package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cli struct {
	prompt string
}

func NewCliWithInit(prompt string) *Cli {
	return &Cli{
		prompt: prompt,
	}
}

func NewCli() *Cli {
	return NewCliWithInit("db")
}

func (cli *Cli) Run() {
	for {
		cli.printPrompt()
		buffer := newInputBuffer()
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, ";") {
				split := strings.Split(line, ";")
				if len(split[0]) > 0 {
					_, err := buffer.Write(split[0])
					if err != nil {
						panic(err)
					}
				}
				sql := buffer.ReadString()
				_, err := os.Stdout.WriteString(sql)
				if err != nil {
					panic(err)
				}
				fmt.Println("ok")
				os.Exit(1)
			} else {
				_, err := buffer.Write(line)
				if err != nil {
					panic(err)
				}
			}
		}

	}
}

func (cli *Cli) printPrompt() {
	fmt.Printf("%s>", cli.prompt)
}
