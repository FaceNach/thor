package main

import (
	"fmt"
	"os"
	"os/user"
	"thor/relp"
)


func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Thor programming language!\n",user.Username)
	fmt.Printf("Feel free to type in commands\n")
	relp.Start(os.Stdin, os.Stdout)
}