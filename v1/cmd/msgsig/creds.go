package main

import (
	"fmt"
	"os"
)

const sep = "$"

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}
	switch cmd := os.Args[1]; cmd {
	case "message":
		message(os.Args[2:])
	case "help":
		help()
	default:
		fmt.Println("*** Invalid command; try $ msgsig help")
	}
}

func message(args []string) {
	switch cmd := args[0]; cmd {
	case "sign":
		signMessage(args[1:])
	case "verify":
		verifyMessage(args[1:])
	default:
		fmt.Println("*** Invalid command; try $ msgsig help")
	}
}

func help() {
	fmt.Println("Usage: msgsig <command> [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println()
	fmt.Println("    message")
	fmt.Println("        sign     Encrypt and sign a message")
	fmt.Println("        verify   Verify and decrypt a signed message")
	fmt.Println()
}
