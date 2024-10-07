package problemset1

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"
	"math"
	"slices"
	"strings"
)

// Set 1 challenge 1
func HexToBase64(hexStr string) (string, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}

	return base64.RawStdEncoding.EncodeToString(bytes), nil
}

// Set 1 challenge 2
func FixedXOR(buf1 string, buf2 string) (string, error) {
	if len(buf1) != len(buf2) {
		return "", errors.New("Buffers are not of equal length")
	}

	b1, err := hex.DecodeString(buf1)
	if err != nil {
		return "", err
	}

	b2, err := hex.DecodeString(buf2)
	if err != nil {
		return "", err
	}

	out := make([]byte, len(b1))
	for i := 0; i < len(b1); i++ {
		out[i] = b1[i] ^ b2[i]
	}

	stringified := hex.EncodeToString(out)
	return stringified, nil
}

// horrible and incorrect, fix it
// Set 1 challenge 3
func SingleByteXORCipher(hexStr string) (rune, error) {
	results := make(map[rune]float64)
	outputStrs := make(map[rune]string)

	var start rune = 0
	for ; start <= 256; start++ {
		repeated := slices.Repeat[[]byte, byte]([]byte{byte(start)}, len(hexStr)/2)
		output, err := FixedXOR(hexStr, hex.EncodeToString(repeated))
		if err != nil {
			log.Panicf("Unexpected error %v", err)
		}

		// should never be errored
		bytes, err := hex.DecodeString(output)
		if err != nil {
			log.Panicf("Unexpected error %v", err)
		}

		outputStrs[start] = string(bytes)
		results[start] = scoreEnglish(string(bytes))
	}

	currMin := math.MaxFloat64
	currBest := '-'
	for k, v := range results {
		if v < currMin {
			currMin = v
			currBest = k
		}
	}

	log.Printf("The best string is '%s'", outputStrs[currBest])
	return currBest, nil
}

func scoreEnglish(proposed string) float64 {
	// TODO: maybe just download a book from gutenberg and manually generate this table
	// source: https://en.wikipedia.org/wiki/Letter_frequency
	frequencyTable := map[rune]float64{
		'a': 0.082,
		'b': 0.015,
		'c': 0.028,
		'd': 0.043,
		'e': 0.127,
		'f': 0.022,
		'g': 0.020,
		'h': 0.061,
		'i': 0.070,
		'j': 0.002,
		'k': 0.008,
		'l': 0.040,
		'm': 0.024,
		'n': 0.067,
		'o': 0.075,
		'p': 0.019,
		'q': 0.001,
		'r': 0.060,
		's': 0.063,
		't': 0.091,
		'u': 0.028,
		'v': 0.010,
		'w': 0.024,
		'x': 0.002,
		'y': 0.020,
		'z': 0.001,
	}

	// get count of characters in string
	count := make(map[rune]int)
	normalized := strings.ToLower(proposed)
	for _, c := range normalized {
		count[c] += 1
	}

	// get density of characters in string
	tot := len(normalized)
	density := make(map[rune]float64)
	for k, v := range count {
		density[k] = float64(v) / float64(tot)
	}

	diffScore := 0.0
	// compare density to english lexicon density
	for k, v := range frequencyTable {
		diffScore += math.Abs(v - density[k])
	}
	return diffScore
}
