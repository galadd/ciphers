// console application to choose cipher and encrypt/decrypt a message

package main

import (
	"fmt"
	"bufio"
	"os"

	"github.com/galadd/ciphers/ciphers/caesar"
	"github.com/galadd/ciphers/ciphers/baconian"
	"github.com/galadd/ciphers/ciphers/gronsfeld"
	"github.com/galadd/ciphers/ciphers/porta"
)

func main() {
	var ed int
	fmt.Println("Enter 1 to Encrypt or 2 to Decrypt")
	fmt.Scanln(&ed)

	var choice int
	fmt.Println("Choose a cipher:")
	fmt.Println("1. Caesar")
	fmt.Println("2. Baconian")
	fmt.Println("3. Gronsfeld")
	fmt.Println("4. Porta")
	fmt.Println()
	fmt.Print("⏩ ⏩ ⏩  ")
	fmt.Scanln(&choice)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	fmt.Println("Write your message:")
	scanner.Scan()
	input := scanner.Text()
	var codec string

	switch choice {
	case 1: 
		if ed == 1 {
			codec = caesar.Encrypt(input)
		}
		if ed == 2 {
			codec = caesar.Decrypt(input)
		}
	case 2:
		if ed == 1 {
			codec = baconian.Encrypt(input)
		}
		if ed == 2 {
			codec = baconian.Decrypt(input)
		}
	case 3:
		var key string
		fmt.Print("Input key: ")
		fmt.Scan(&key)

		if ed == 1 {
			codec = gronsfeld.Encrypt(input, key)
		}
		if ed == 2 {
			codec = gronsfeld.Decrypt(input, key)
		}
	case 4:
		if ed == 1 {
			codec = porta.Encrypt(input)
		}
		if ed == 2 {
			codec = porta.Decrypt(input)
		}
	}
	fmt.Println()
	fmt.Println("⬇ 💻🤓 ⬇ 💻🤓 ⬇ 💻🤓 ⬇ 💻🤓 ⬇")

	if ed == 1 {
		fmt.Println("Message Encrypted: ", codec)
	}
	if ed == 2 {
		fmt.Println("Message Decrypted: ", codec)
	}
}