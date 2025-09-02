package main

import (
	"fmt"
	"os"

	"lem_in/helpers"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("ERROR: " + "Invalid number of arguments usage Example => go run . test.txt : ")
		return
	}
	file := args[0]

	if err := helpers.HandleError(file); err != nil {
		fmt.Println(err)
		return
	}
	
	AllPaths := helpers.FindPaths()
	if AllPaths == nil {
		fmt.Println("ERROR: " + "You don't have any links in your file : ")
		return
	}
	FileContent, err := os.ReadFile(helpers.PathFiles + file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(FileContent)+"\n")

	helpers.Solve(AllPaths)
}
