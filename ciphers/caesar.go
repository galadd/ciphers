package caesar

func Encrypt(message string) string {
	var encrypted string
	
	// Create a lookup table for the encrypted characters
	lookup := make(map[rune]rune)
	for i := 'a'; i <= 'z'; i++ {
		lookup[i] = rune((int(i - 'a') + 3) % 26 + 'a')
	}
	for i := 'A'; i <= 'Z'; i++ {
		lookup[i] = rune((int(i - 'A') + 3) % 26 + 'A')
	}
	
	// Use the lookup table to encrypt the input string
	for _, char := range message {
		encrypted += string(lookup[char])
	}
	
	return encrypted
}

func Decrypt(message string) string {
	var decrypted string
	
	// Create a lookup table for the decrypted characters
	lookup := make(map[rune]rune)
	for i := 'a'; i <= 'z'; i++ {
		lookup[i] = rune((int(i - 'a') - 3) % 26 + 'a')
	}
	for i := 'A'; i <= 'Z'; i++ {
		lookup[i] = rune((int(i - 'A') - 3) % 26 + 'A')
	}
	
	// Use the lookup table to decrypt the input string
	for _, char := range message {
		decrypted += string(lookup[char])
	}
	
	return decrypted
}