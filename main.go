package main

import (
	"fmt"
	"github.com/Marneus68/spac/packager"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Not enough arguments provided.\n")
		PrintUsage()
		os.Exit(1)
	}

	d, err := os.Stat(os.Args[1])
	if os.IsNotExist(err) {
		fmt.Println("Directory " + os.Args[1] + " doesn't exist.\n")
		PrintUsage()
		os.Exit(1)
	}

	if !d.IsDir() {
		fmt.Println(os.Args[1] + "isn't a directory.\n")
		PrintUsage()
		os.Exit(1)
	}

	packager.Create(os.Args[1], os.Args[2])
}

func PrintUsage() {
	fmt.Println("Usage: spac <directory> <filename>")
	fmt.Println("")
	fmt.Println("    <directory> : directory containing the resources to package")
	fmt.Println("    <filename> : name of the golang source file packaging all the resources")
	fmt.Println("")
}
