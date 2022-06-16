package main

import (
	"fmt"
	"os"
)

const (
	usage     = "Usage: [username] [password]"
	errUser   = "Access denide for %q.\n"
	errPasswd = "Invalid passwor for %q.\n"
	accessOk  = "Access granted for %q.\n"

	user1, user2 = "Roman", "Olha"
	pswd1, pswd2 = "0420", "0303"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}
	username, password := args[1], args[2]
	if username != user1 && username != user2 {
		fmt.Printf(errUser, username)
	} else if username == user1 && password == pswd1 ||
		username == user2 && password == pswd2 {
		fmt.Printf(accessOk, username)
	} else {
		fmt.Printf(errPasswd, username)
	}
}
