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
				lowerSql := strings.ToLower(sql)

				if strings.Contains(lowerSql, "exit") {
					os.Stdout.WriteString("  byebye")
					os.Exit(1)
				}

				// todo:

			} else {
				_, err := buffer.Write(line)
				if err != nil {
					panic(err)
				}
			}

			_, err := os.Stdout.WriteString(cli.prompt + ">")
			if err != nil {
				panic(err)
			}
		}

	}
}

func (cli *Cli) printPrompt() {
	fmt.Printf("%s>", cli.prompt)
}
