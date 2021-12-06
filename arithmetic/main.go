package main

import "log"

type plugin struct {
}

func (p *plugin) Init() {
}

func (p *plugin) Help(name string) string {
	helpText := map[string]string{
		"AddInt": "Add two integer",
		"MulInt": "Multiply two integer",
	}

	help, ok := helpText[name]
	if !ok {
		return ""
	}

	return help
}

func (p *plugin) AddInt(a, b int) {
	log.Printf("Add: %d + %d = %d", a, b, a+b)
}

func (p *plugin) MulInt(a, b int) {
	log.Printf("Multiply: %d + %d = %d", a, b, a*b)
}
