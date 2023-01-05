// console application to choose cipher and encrypt/decrypt a message

package main

import (
	"fmt"

	"github.com/galadd/ciphers/ciphers"
)

func main() {
	var ed int
	fmt.Println("Enter 1 to Encrypt or 2 to Decrypt")
	fmt.Scanln(&ed)

	var choice int
	fmt.Println("Choose a cipher:")
	fmt.Println("1. Caesar")
	fmt.Scanln(&choice)

	var message string
	fmt.Println("Write your message:")
	fmt.Scanln(&message)

	var codec string

	switch choice {
	case 1: 
	if ed == 1 {
		codec = caesar.Encrypt(message)
	}
	if ed == 2 {
		codec = caesar.Decrypt(message)
	}
	}
	fmt.Println()
	fmt.Println("⬇️  ⬇️  ⬇️  ⬇️  ⬇️")

	if ed == 1 {
		fmt.Printf("Message Encrypted: %s\n", codec)
	}
	if ed == 2 {
		fmt.Println("Message Decrypted: %s\n", codec)
	}
}