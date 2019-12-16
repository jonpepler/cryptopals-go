package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type letterFrequency struct {
	char  string
	score int
}

func main() {
	input := "input.txt"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	texts, _ := readFileToText(input)

	xorTexts := []xorText{}
	for _, text := range texts {
		ba := hexToByteArray(text)
		xorTexts = append(xorTexts, generateTexts(ba)...)
	}
	sort.Slice(xorTexts, func(i, j int) bool {
		return xorTexts[i].englishScore() > xorTexts[j].englishScore()
	})

	fmt.Println("\n\nMost likely:")
	fmt.Printf("%v: %v", string(xorTexts[0].key), string(xorTexts[0].bytes))
}

type xorText struct {
	bytes []byte
	key   byte
}

func (xt xorText) englishScore() (score int) {
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

	text := string(xt.bytes)

	for _, c := range heuristic {
		score += c.score * strings.Count(text, c.char)
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

func readFileToText(fname string) (texts []string, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	return lines, nil
}
