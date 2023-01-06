package gronsfeld

import "unicode"

func Encrypt(message, key string) string {
	var encrypted string
	var keyPosition int

	for _, char := range message {
		if keyPosition == len(key) {
			keyPosition = 0
		}
		shift := int(key[keyPosition] - '0')
		keyPosition++

		if unicode.IsLetter(char) {
			if unicode.IsUpper(char) {
				char = (char - 'A' + rune(shift)) % 26 + 'A'
			} else {
				char = (char - 'a' + rune(shift)) % 26 + 'a'
			}
		}

		encrypted += string(char)
	}

	return encrypted
}

func Decrypt(message, key string) string {
	var decrypted string
	var keyPosition int

	for _, char := range message {
		if keyPosition == len(key) {
			keyPosition = 0
		}
		shift := int(key[keyPosition] - '0')
		keyPosition++

		if unicode.IsLetter(char) {
			if unicode.IsUpper(char) {
				char = (char - 'A' - rune(shift) + 26) % 26 + 'A'
			} else {
				char = (char - 'a' - rune(shift) + 26) % 26 + 'a'
			}
		}

		decrypted += string(char)
	}

	return decrypted
}