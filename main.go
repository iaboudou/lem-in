package main

import (
	"fmt"

	"lem_in/helpers"
)

func main() {
	if err := helpers.HandleError("test.txt") ; err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Succes file")
}
