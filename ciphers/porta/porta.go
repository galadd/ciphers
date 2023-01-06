package porta

import "unicode"

func Encrypt(message string) string {
	var encrypted string
	var keyPosition int
	key := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for _, char := range message {
		if keyPosition == len(key) {
			keyPosition = 0
		}
		shift := int(key[keyPosition] - 'A')
		keyPosition++

		if unicode.IsLetter(char) {
			if shift > 13 {
				shift = 26 - shift
			}
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

func Decrypt(message string) string {
	var decrypted string
	var keyPosition int
	key := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for _, char := range message {
		if keyPosition == len(key) {
			keyPosition = 0
		}
		shift := int(key[keyPosition] - 'A')
		keyPosition++

		if unicode.IsLetter(char) {
			if shift > 13 {
				shift = 26 - shift
			}
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