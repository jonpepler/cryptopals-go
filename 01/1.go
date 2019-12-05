package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println(hexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
}

func hexToBase64(hexString string) string {
	bytes, _ := hex.DecodeString(hexString)
	return base64.StdEncoding.EncodeToString(bytes)
}
