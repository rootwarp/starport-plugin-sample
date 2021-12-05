package main

import (
	"fmt"
	"log"

	"github.com/motki/cli/text/banner"
)

var (
	helpStr = map[string]string{
		"SayHello": "Just say hello",
	}
)

type hello struct {
}

func (p *hello) Init() {
	log.Println("Initialize plugin")
}

func (p *hello) SayHello(name string) {
	sep := `
================================================================================
`
	fmt.Println("")
	fmt.Println(sep)
	fmt.Printf("Hello %s\n", name)
	fmt.Println(sep)
	fmt.Println("")
}

func (p *hello) Banner(name string) {
	banner.Printf(name + "\n")
}

func (p *hello) Help(name string) string {
	t, ok := helpStr[name]
	if !ok {
		return ""
	}

	return t
}

var Plugin hello
