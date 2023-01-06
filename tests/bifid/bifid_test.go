package bifid

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/galadd/ciphers/ciphers/bifid"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name string
		input string
		expectedOutput string
	}{
		{
			name: "Encrypt",
			input: "aBc",
			expectedOutput: "NAY",
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := bifid.Encrypt(test.input)
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
			name: "Decrypt",
			input: "NAY",
			expectedOutput: "ABC",
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := bifid.Decrypt(test.input)
			assert.Equal(t, test.expectedOutput, output)
		})
	}
}