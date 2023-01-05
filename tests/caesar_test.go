package caesar

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/galadd/ciphers/ciphers"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name string
		input string
		expectedOutput string
	}{
		{
			name: "Encrypt lowercase",
			input: "abc",
			expectedOutput: "def",
		},
		{
			name: "Encrypt uppercase",
			input: "ABC",
			expectedOutput: "DEF",
		},
		{
			name: "Encrypt mixed case",
			input: "aBc",
			expectedOutput: "dEf",
		},
		{
			name: "Encrypt special characters",
			input: "!@#$%^&*()",
			expectedOutput: "!@#$%^&*()",
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := caesar.Encrypt(test.input)
			assert.Equal(t, test.expectedOutput, output)
		})
	}
}
