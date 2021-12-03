package main

import "fmt"

var (
	helpStr = map[string]string{
		"SayHello": "Just say hello",
	}
)

type hello struct {
}

func (p *hello) Initialize() {
	fmt.Println("Initialize plugin")
}

func (p *hello) SayHello(name string) {
	fmt.Printf("Hello %s\n", name)
}

func (p *hello) Help(name string) string {
	t, ok := helpStr[name]
	if !ok {
		return ""
	}

	return t
}

var Plugin hello
