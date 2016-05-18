package main

import (
	"fmt"
	"github.com/Marneus68/spac/packager"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: spac <directory> <filename>")
		fmt.Println("")
		fmt.Println("    <directory> : directory containing the resources to package")
		fmt.Println("    <filename> : name of the golong source file packaging all the resoucres")
		fmt.Println("")
		os.Exit(0)
	}

	packager.Create(os.Args[1], os.Args[2])
}
