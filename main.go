package main

import "github.com/jacksonyoudi/litedb/repl"

func main() {
	cli := repl.NewCliWithInit("db")
	cli.Run()
}
