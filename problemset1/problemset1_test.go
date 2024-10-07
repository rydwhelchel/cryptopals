package problemset1_test

import (
	"encoding/hex"
	"errors"
	"testing"

	"github.com/rydwhelchel/cryptopals/problemset1"
)

func TestHexTo64(t *testing.T) {
	testCases := []struct {
		input  string
		output string
		err    error
		name   string
	}{
		{
			input:  "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			output: "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
			err:    nil,
			name:   "Given input",
		},
		{
			input:  "49276d20gb696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			output: "",
			err:    hex.InvalidByteError('g'),
			name:   "Invalid input (g)",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			out, err := problemset1.HexToBase64(tt.input)
			if tt.output != out {
				t.Errorf("Failed to correctly convert hex to base64:\nexpected:\n%s\ngot:\n%s\n", tt.output, out)
			} else if tt.err != nil && err != nil && tt.err.Error() != err.Error() {
				t.Errorf("Incorrect error while converting hex to base64:\nexpected:\n%s\ngot:\n%s\n", tt.err.Error(), err.Error())
			}
		})
	}
}

func TestFixedXOR(t *testing.T) {
	testCases := []struct {
		input1 string
		input2 string
		output string
		err    error
		name   string
	}{
		{
			input1: "1c0111001f010100061a024b53535009181c",
			input2: "686974207468652062756c6c277320657965",
			output: "746865206b696420646f6e277420706c6179",
			err:    nil,
			name:   "Given input",
		},
		{
			input1: "1c0111001f010100061a024b53535009181c",
			input2: "686974207468652062756c6c277320657965123",
			output: "",
			err:    errors.New("Buffers are not of equal length"),
			name:   "Inequal length buffers",
		},
		{
			input1: "1c0111001f010100061g024b53535009181c",
			input2: "686974207468652062756c6c277320657965",
			output: "",
			err:    hex.InvalidByteError('g'),
			name:   "Invalid character in buffer1",
		},
		{
			input1: "1c0111001f010100061a024b53535009181c",
			input2: "6869742074686y2062756c6c277320657965",
			output: "",
			err:    hex.InvalidByteError('y'),
			name:   "Invalid character in buffer2",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			out, err := problemset1.FixedXOR(tt.input1, tt.input2)
			if out != tt.output {
				t.Errorf("Unable to correctly XOR:\nexpected:\n%s\ngot:\n%s\n", tt.output, out)
			} else if tt.err != nil && err != nil && tt.err.Error() != err.Error() {
				t.Errorf("Incorrect error while converting hex to base64:\nexpected:\n%s\ngot:\n%s\n", tt.err.Error(), err.Error())
			}
		})
	}
}

func TestSingleByteXORCipher(t *testing.T) {
	testCases := []struct {
		input  string
		output rune
		err    error
		name   string
	}{
		{
			input: "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",
			// todo: i don't know the answer yet, just using this as a runner
			output: 'i',
			err:    nil,
			name:   "Provided input",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			output, err := problemset1.SingleByteXORCipher(tt.input)
			if output != tt.output {
				t.Errorf("Incorrect single byte cipher:\nexpected:\n%s\ngot:\n%s\n", string(tt.output), string(output))
			} else if tt.err != nil && err != nil && tt.err.Error() != err.Error() {
				t.Errorf("Incorrect error while calculating cipher:\nexpected:\n%s\ngot:\n%s\n", tt.err.Error(), err.Error())
			}
		})
	}
}
