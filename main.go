// console application to choose cipher and encrypt/decrypt a message

package main

import (
	"fmt"

	"github.com/galadd/ciphers/ciphers"
)

func main() {
	encryption := caesar.Encrypt("hello")
	fmt.Println(encryption)
}