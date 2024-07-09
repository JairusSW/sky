package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var args = os.Args[1:]

func main() {
	if len(args) == 0 {
		color.New(color.Bold, color.FgBlue).Print("Sky")
		color.New(color.Bold).Print(" is a self hosted server toolkit written in Go. ")
		color.New(color.Bold, color.Faint).Println("(v0.0.0)")

		color.New(color.Bold).Print("\nUsage: sky ")
		color.New(color.Bold, color.Faint).Print(" <command>")
		color.New(color.Bold, color.FgBlue).Print(" [...flags]")
		color.New(color.Bold).Print(" [...args]\n")

		color.New(color.Bold).Println("Commands:")
		color.New(color.Bold, color.FgBlue).Print("  tunnel    "); color.New(color.Bold, color.Faint).Print("-p <port>    "); fmt.Println("Tunnel port through server")
		color.New(color.Bold, color.FgBlue).Print("  file      "); color.New(color.Bold, color.Faint).Print("-p <port>    "); fmt.Println("Setup file server on server")
		color.New(color.Bold, color.FgBlue).Print("  dash      "); color.New(color.Bold, color.Faint).Print("             "); fmt.Println("Access the Admin dashboard")
	}
}
