package baconian

import "strings"

func Encrypt(message string) string {
	message = strings.ToUpper(message)
	var encrypted string

	lookup := map[string]string{
		"A": "00000","B": "00001","C": "00010","D": "00011","E": "00100",
		"F": "00101","G": "00110","H": "00111","I": "01000","J": "01001",
		"K": "01010","L": "01011","M": "01100","N": "01101","O": "01110",
		"P": "01111","Q": "10000","R": "10001","S": "10010","T": "10011",
		"U": "10100","V": "10101","W": "10110","X": "10111","Y": "11000",
		"Z": "11001",
	}
	
	for _, char := range message {
		if lookup[string(char)] != "" {
			encrypted += lookup[string(char)]
		} else {
			encrypted += string(char)
		}
	}

	return encrypted
}

func Decrypt(message string) string {
	var decrypted string 

	lookup := map[string]string{
		"00000":"A", "00001":"B", "00010":"C", "00011":"D", "00100":"E",
		"00101":"F", "00110":"G", "00111":"H", "01000":"I", "01001":"J",
		"01010":"K", "01011":"L", "01100":"M", "01101":"N", "01110":"O",
		"01111":"P", "10000":"Q", "10001":"R", "10010":"S", "10011":"T",
		"10100":"U", "10101":"V", "10110":"W", "10111":"X", "11000":"Y",
		"11001":"Z",
	}

	for i := 0; i < len(message); {
		if message[i] == 32 {
			decrypted += " "
			i++
		} else {
			a := string(message[i]) + string(message[i+1])
			b := string(message[i+2]) + string( message[i+3]) + string(message[i+4]) 
			char := a + b 
			decrypted += lookup[char]
			i += 5
		}		
	}
	
	return decrypted
}