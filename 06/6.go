package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type keysizeScore struct {
	keysize int
	score   float64
}

type xorText struct {
	bytes []byte
	key   byte
}

type letterFrequency struct {
	char  string
	score int
}

type passwordGuess struct {
	keysize int
	score   int
	text    string
}

func main() {
	input := "6.txt"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	text, _ := readFileToBytes(input)

	keysizes := bestKeysizes(text, 3)
	fmt.Printf("Best keysizes: %v\n", keysizes)

	guesses := []passwordGuess{}
	for _, keysize := range keysizes {
		guess := string(makeBestGuess(text, keysize))
		guesses = append(guesses, passwordGuess{keysize, englishScore(guess), guess})
	}

	for _, guess := range guesses {
		fmt.Printf("%v(%v): %v\n", guess.score, guess.keysize, guess.text)
	}

	key := generateKey(guesses[0].text, len(text))
	fmt.Println("Full Text Output:")
	fmt.Println(string(xorBytes(key, text)))
}

func generateKey(input string, length int) (key []byte) {
	keyBytes := []byte(input)

	for i := 0; len(key) < length; i++ {
		key = append(key, keyBytes[i%len(input)])
	}

	return
}

func xorBytes(a []byte, b []byte) (o []byte) {
	if len(a) != len(b) {
		panic("Byte arrays were of different lengths!")
	}

	for i := 0; i < len(a); i++ {
		o = append(o, a[i]^b[i])
	}
	return
}

func (xt xorText) englishScore() int {
	return englishScore(string(xt.bytes))
}

func englishScore(text string) (score int) {
	heuristic := []letterFrequency{
		letterFrequency{"a", 651738},
		letterFrequency{"b", 124248},
		letterFrequency{"c", 217339},
		letterFrequency{"d", 349835},
		letterFrequency{"e", 1041442},
		letterFrequency{"f", 197881},
		letterFrequency{"g", 158610},
		letterFrequency{"h", 492888},
		letterFrequency{"i", 558094},
		letterFrequency{"j", 9033},
		letterFrequency{"k", 50529},
		letterFrequency{"l", 331490},
		letterFrequency{"m", 0202124},
		letterFrequency{"n", 0564513},
		letterFrequency{"o", 596302},
		letterFrequency{"p", 0137645},
		letterFrequency{"q", 8606},
		letterFrequency{"r", 497563},
		letterFrequency{"s", 0515760},
		letterFrequency{"t", 729357},
		letterFrequency{"u", 0225134},
		letterFrequency{"v", 82903},
		letterFrequency{"w", 0171272},
		letterFrequency{"x", 13692},
		letterFrequency{"y", 145984},
		letterFrequency{"z", 7836},
		letterFrequency{" ", 1918182},
	}

	for _, c := range heuristic {
		score += c.score * strings.Count(text, c.char)
	}

	return
}

func makeBestGuess(cipherText []byte, keysize int) (guess []byte) {
	blocks := makeBlocks(cipherText, keysize)
	tBlocks := transposeBlocks(blocks)

	for _, tBlock := range tBlocks {
		guess = append(guess, findKey(tBlock))
	}

	return
}

func transposeBlocks(blocks [][]byte) (transposedBlocks [][]byte) {
	blockLength := len(blocks[0])

	for i := 0; i < blockLength; i++ {
		tBlock := []byte{}
		for j := 0; j < len(blocks); j++ {
			tBlock = append(tBlock, blocks[j][i])
		}
		transposedBlocks = append(transposedBlocks, tBlock)
	}
	return
}

func makeBlocks(text []byte, keysize int) (blocks [][]byte) {
	for i := 0; i < len(text)-keysize; i += keysize {
		blocks = append(blocks, text[i:i+keysize])
	}
	return
}

func bestKeysizes(text []byte, topNum int) (keysizes []int) {
	keysizeScores := []keysizeScore{}
	for keysize := 2; keysize <= 40; keysize++ {
		keysizeScores = append(keysizeScores, testKeysize(text, keysize))
	}

	// sort keysizeScores
	sort.Slice(keysizeScores, func(i, j int) bool {
		return keysizeScores[i].score < keysizeScores[j].score
	})

	// return top 3 as array of keysizes
	for i := 0; i < topNum; i++ {
		keysizes = append(keysizes, keysizeScores[i].keysize)
	}
	return
}

func testKeysize(text []byte, keysize int) keysizeScore {
	score, counter := float64(0), 0
	blocks := makeBlocks(text, keysize)

	for i, block := range blocks {
		if i != len(blocks)-1 {
			nextBlock := blocks[i+1]
			result := hammingDistance(block, nextBlock)

			score += (float64(result))
			counter++
		}
	}

	score /= float64(counter) * float64(keysize)

	return keysizeScore{keysize, score}
}

// https://stackoverflow.com/a/40309527
func hammingDistance(a, b []byte) int {
	if len(a) != len(b) {
		return -1
	}

	diff := 0
	for i := 0; i < len(a); i++ {
		b1 := a[i]
		b2 := b[i]
		for j := 0; j < 8; j++ {
			mask := byte(1 << uint(j))
			if (b1 & mask) != (b2 & mask) {
				diff++
			}
		}
	}
	return diff
}

func findKey(text []byte) byte {
	// fmt.Printf("Finding key for: %v\n", string(text))
	xordTexts := generateTexts(text)
	sort.Slice(xordTexts, func(i, j int) bool {
		return xordTexts[i].englishScore() > xordTexts[j].englishScore()
	})
	return xordTexts[0].key
}

func generateTexts(text []byte) (texts []xorText) {
	for i := 0; i <= 255; i++ {
		iByte := byte(i)
		xord := xorSingleByte(text, iByte)
		texts = append(texts, xorText{xord, iByte})
	}
	return
}

func xorSingleByte(a []byte, key byte) (o []byte) {
	for i := 0; i < len(a); i++ {
		o = append(o, a[i]^key)
	}
	return
}

func readFileToBytes(fname string) ([]byte, error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	decoded, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
