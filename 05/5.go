package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	input := stringToByteArray("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	keyInput := "ICE"
	expectedOutput := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	fmt.Printf("Encrypting '%v' with key '%v'...\n", string(input), keyInput)

	key := generateKey(keyInput, len(input))
	fmt.Printf("Generated key: %v\n", key)
	fmt.Printf("Length check: %v\n", len(key) == len(input))

	xord := xorBytes(input, key)
	out := byteArrayToHex(xord)
	fmt.Println(out)
	fmt.Printf("Matches expectation: %v", out == expectedOutput)
}

func stringToByteArray(str string) []byte {
	return []byte(str)
}

func generateKey(input string, length int) (key []byte) {
	keyBytes := []byte(input)

	for i := 0; len(key) < length; i++ {
		key = append(key, keyBytes[i%3])
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

func byteArrayToHex(bytes []byte) string {
	hex := hex.EncodeToString(bytes)
	return hex
}
