package main

import (
	"fmt"

	"github.com/darksasori/finance/cli/cmd"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	cmd.Execute()
}
