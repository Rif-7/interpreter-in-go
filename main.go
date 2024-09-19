package main

import (
	"fmt"
	"interpreter-in-go/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the Rif programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)

}

