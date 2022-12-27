package main

import (
	"github.com/fatih/color"
	"hello/greeting"
)

func main() {
	c := color.New(color.FgCyan).Add(color.Underline)
	c.Printf(greeting.Hello())

}
