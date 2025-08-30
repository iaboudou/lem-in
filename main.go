package main

import (
	"fmt"

	"lem_in/helpers"
)

func main() {
	if err := helpers.HandleError("test.txt"); err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range helpers.FindPaths() {
		fmt.Println(v)
	}
	fmt.Println("*************************\n***************************")
	fp := helpers.FindPaths()
	helpers.Solve(fp)
}
