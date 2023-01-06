package baconian

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/galadd/ciphers/ciphers/baconian"
)

func BaconianTest_Encrypt(t *testing.T) {
	tests := []struct {
		name string
		input string
		expectedOutput string
	}{
		{
			name: "Encrypt lowercase",
			input: "abc",
			expectedOutput: "000000000100010",
		},
		{
			name: "Encrypt uppercase",
			input: "ABC",
			expectedOutput: "000000000100010",
		},
		{
			name: "Encrypt mixed case",
			input: "aBc",
			expectedOutput: "000000000100010",
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := baconian.Encrypt(test.input)
			assert.Equal(t, test.expectedOutput, output)
		})
	}
}

func BaconianTest_Decrypt(t *testing.T) {
	tests := []struct {
		name string
		input string
		expectedOutput string
	}{
		{
			name: "Decrypt",
			input: "000000000100010",
			expectedOutput: "ABC",
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := baconian.Decrypt(test.input)
			assert.Equal(t, test.expectedOutput, output)
		})
	}
}