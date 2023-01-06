package gronsfeld

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/galadd/ciphers/ciphers/gronsfeld"
)

func TestEncrypt(t *testing.T) {
	key := "123"
	tests := []struct {
		name string
		input string
		expectedOutput string
	}{
		{
			name: "Encrypt lowercase",
			input: "abc",
			expectedOutput: "bdf",
		},
		{
			name: "Encrypt uppercase",
			input: "ABC",
			expectedOutput: "BDF",
		},
		{
			name: "Encrypt mixed case",
			input: "aBc",
			expectedOutput: "bDf",
		},
		{
			name: "Encrypt special characters",
			input: "!@#$%^&*()",
			expectedOutput: "!@#$%^&*()",
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := gronsfeld.Encrypt(test.input, key)
			assert.Equal(t, test.expectedOutput, output)
		})
	}
}

func TestDecrypt(t *testing.T) {
	key := "123"
	tests := []struct {
		name string
		input string
		expectedOutput string
	}{
		{
			name: "Decrypt lowercase",
			input: "bdf",
			expectedOutput: "abc",
		},
		{
			name: "Decrypt uppercase",
			input: "BDF",
			expectedOutput: "ABC",
		},
		{
			name: "Decrypt mixed case",
			input: "bDf",
			expectedOutput: "aBc",
		},
		{
			name: "Decrypt special characters",
			input: "!@#$%^&*()",
			expectedOutput: "!@#$%^&*()",
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := gronsfeld.Decrypt(test.input, key)
			assert.Equal(t, test.expectedOutput, output)
		})
	}
}