package main

import (
	"fmt"
	"onepiece/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Time to Hacking!\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
