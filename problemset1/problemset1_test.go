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
	context := problemset1.Context{
		BookPath: "../throughthelookingglass.txt",
	}

	testCases := []struct {
		input     string
		outputRun rune
		outputStr string
		err       error
		name      string
	}{
		{
			input:     "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",
			outputRun: 'X',
			outputStr: "Cooking MC's like a pound of bacon",
			err:       nil,
			name:      "Provided input",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			results, err := context.SingleByteXORCipher(tt.input)
			if results.Byte != tt.outputRun {
				t.Errorf("Incorrect single byte cipher:\nexpected:\n%s\ngot:\n%s\n", string(tt.outputRun), string(results.Byte))
			} else if results.Decrypted != tt.outputStr {
				t.Errorf("Incorrect output from single byte cipher:\nexpected:\n%s\ngot:\n%s\n", tt.outputStr, results.Decrypted)
			} else if tt.err != nil && err != nil && tt.err.Error() != err.Error() {
				t.Errorf("Incorrect error while calculating cipher:\nexpected:\n%s\ngot:\n%s\n", tt.err.Error(), err.Error())
			}
		})
	}
}

func TestFindSingleByteEncryption(t *testing.T) {
	context := problemset1.Context{
		BookPath: "../throughthelookingglass.txt",
	}

	testCases := []struct {
		input     string
		outputRun rune
		outputStr string
		err       error
		name      string
	}{
		{
			input:     "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",
			outputRun: 'X',
			outputStr: "Cooking MC's like a pound of bacon",
			err:       nil,
			name:      "Provided input",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			results, err := context.SingleByteXORCipher(tt.input)
			if results.Byte != tt.outputRun {
				t.Errorf("Incorrect single byte cipher:\nexpected:\n%s\ngot:\n%s\n", string(tt.outputRun), string(results.Byte))
			} else if results.Decrypted != tt.outputStr {
				t.Errorf("Incorrect output from single byte cipher:\nexpected:\n%s\ngot:\n%s\n", tt.outputStr, results.Decrypted)
			} else if tt.err != nil && err != nil && tt.err.Error() != err.Error() {
				t.Errorf("Incorrect error while calculating cipher:\nexpected:\n%s\ngot:\n%s\n", tt.err.Error(), err.Error())
			}
		})
	}
}

func TestICEEncryption(t *testing.T) {
	testCases := []struct {
		input  string
		output string
		name   string
	}{
		{
			input:  "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal",
			output: "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f",
			name:   "Given input",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			output := problemset1.ICEEncryption(tt.input)
			if output != tt.output {
				t.Errorf("Did not receieve correct answer\nExpected:\n\t%s\nReceieved:\n\t%s", tt.output, output)
			}
		})
	}
}
