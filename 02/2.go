package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	if len(os.Args) > 2 {
		input1 = os.Args[1]
		input2 = os.Args[2]
	}

	input1Bytes := hexToByteArray(input1)
	input2Bytes := hexToByteArray(input2)

	xord := xorBuffer(input1Bytes, input2Bytes)

	fmt.Println(byteArrayToHex(xord))
}

func hexToByteArray(hexString string) []byte {
	bytes, _ := hex.DecodeString(hexString)
	return bytes
}

func byteArrayToHex(bytes []byte) string {
	hex := hex.EncodeToString(bytes)
	return hex
}

func xorBuffer(a []byte, b []byte) (o []byte) {
	if len(a) != len(b) {
		panic("Byte arrays were of different lengths!")
	}

	for i := 0; i < len(a); i++ {
		o = append(o, a[i]^b[i])
	}
	return
}
