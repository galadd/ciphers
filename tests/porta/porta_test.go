package porta

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/galadd/ciphers/ciphers/porta"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name string
		input string
		expectedOutput string
	}{
		{
			name: "Encrypt message",
			input: "abc DEF",
			expectedOutput: "ace HJL",
		},
		{
			name: "Encrypt special characters",
			input: "!@#$%^&*()",
			expectedOutput: "!@#$%^&*()",
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := porta.Encrypt(test.input)
			assert.Equal(t, test.expectedOutput, output)
		})
	}
}

func TestDecrypt(t *testing.T) {
	tests := []struct {
		name string
		input string
		expectedOutput string
	}{
		{
			name: "Decrypt message",
			input: "ace HJL",
			expectedOutput: "abc DEF",
		},
		{
			name: "Decrypt special characters",
			input: "!@#$%^&*()",
			expectedOutput: "!@#$%^&*()",
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := porta.Decrypt(test.input)
			assert.Equal(t, test.expectedOutput, output)
		})
	}
}