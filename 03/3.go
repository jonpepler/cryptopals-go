package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"sort"
	"strings"
)

const heuristic = "etaoin shrdlu"

func main() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	hexInput := hexToByteArray(input)
	result := string(xorSingleByte(hexInput, findKey(hexInput)))

	fmt.Println("\n\nMost likely:")
	fmt.Println(result)
}

type xorText struct {
	bytes []byte
	key   byte
}

func (xt xorText) englishScore() (score int) {
	fmt.Printf("Scoring: %v\n", string(xt.bytes))
	byteHeuristic := []byte(heuristic)
	text := string(xt.bytes)
	frequencies := make(map[byte]int)
	for _, hbyte := range byteHeuristic {
		frequencies[hbyte] = strings.Count(text, string(hbyte))
	}

	previousCount := 0
	for _, count := range frequencies {
		if count < previousCount {
			score++
		}
		previousCount = count
	}

	return
}

func hexToByteArray(hexString string) []byte {
	bytes, _ := hex.DecodeString(hexString)
	return bytes
}

func byteArrayToHex(bytes []byte) string {
	hex := hex.EncodeToString(bytes)
	return hex
}

func findKey(text []byte) byte {
	fmt.Printf("Finding key for: %v\n", string(text))
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
		fmt.Printf("Generating text with key %v: %v\n", iByte, string(xord))
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
