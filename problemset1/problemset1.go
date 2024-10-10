package problemset1

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"
	"math"
	"os"
	"slices"
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

type SingleByteXORCipherResults struct {
	Byte      rune
	Decrypted string
	Score     float64
}

type Context struct {
	BookPath string
	S1c4Path string
}

// Set 1 challenge 3
func (ctx *Context) SingleByteXORCipher(hexStr string) (SingleByteXORCipherResults, error) {
	results := make(map[rune]float64)
	outputStrs := make(map[rune]string)
	densityTable := ctx.getEnglishLetterFrequency()

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
		results[start] = scoreEnglish(string(bytes), densityTable)
	}

	currMin := math.MaxFloat64
	currBest := '-'
	for k, v := range results {
		if v < currMin {
			currMin = v
			currBest = k
		}
	}

	return SingleByteXORCipherResults{currBest, outputStrs[currBest], currMin}, nil
}

func (ctx *Context) getEnglishLetterFrequency() map[rune]float64 {
	frequencyTable := map[rune]int{}

	inputName := ctx.BookPath
	data, err := os.ReadFile(inputName)
	if err != nil {
		log.Panicln("Coudln't read book")
	}
	book := string(data)

	for _, ch := range book {
		if ch >= 'A' && ch <= 'z' {
			frequencyTable[ch]++
		}
	}

	densityTable := map[rune]float64{}
	for k, v := range frequencyTable {
		densityTable[k] = float64(v) / float64(len(book))
	}

	return densityTable
}

func scoreEnglish(proposed string, densityTable map[rune]float64) float64 {
	// get count of characters in string
	count := make(map[rune]int)
	normalized := proposed
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
	for k, v := range density {
		score := math.Abs(v - densityTable[k])
		// If it's not a normal english letter or space, score it a little more skeptically
		if (k < 'A' || k > 'z') && k != ' ' {
			score *= 3
		}
		diffScore += score
	}
	return diffScore
}

// Set 1 challenge 4
func (ctx *Context) FindSingleByteEncryption(lines []string) (SingleByteXORCipherResults, error) {
	results := []SingleByteXORCipherResults{}
	for _, line := range lines {
		result, err := ctx.SingleByteXORCipher(line)
		if err != nil {
			return SingleByteXORCipherResults{}, err
		}
		results = append(results, result)
	}

	best := SingleByteXORCipherResults{Score: math.MaxFloat64}
	for _, result := range results {
		if result.Score < best.Score {
			best = result
		}
	}

	return best, nil
}

// Set 1 challenge 4
func ICEEncryption(input string) string {
	key := []byte{'I', 'C', 'E'}
	inputBytes := []byte(input)
	outputBytes := []byte{}
	for i, b := range inputBytes {
		outputBytes = append(outputBytes, b^key[i%3])
	}
	return hex.EncodeToString(outputBytes)
}
