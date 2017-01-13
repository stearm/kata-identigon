package main

import (
	"fmt"
	"os"
	"github.com/stearm/kata-identigon/utils"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("id and location required")
		return
	}
	utils.BuildImage(string(os.Args[1]), os.Args[2])
}
