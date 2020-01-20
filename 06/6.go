package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input := "6.txt"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	texts, _ := readFileToBytes(input)
	fmt.Println(texts)
}

// https://stackoverflow.com/a/40309527
func hammingDistance(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("a b are not the same length")
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
	return diff, nil
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
