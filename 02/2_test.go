package main

import "testing"

func TestHexToBase64(t *testing.T) {
	input1 := hexToByteArray("1c0111001f010100061a024b53535009181c")
	input2 := hexToByteArray("686974207468652062756c6c277320657965")
	expectedResult := "746865206b696420646f6e277420706c6179"
	if result := byteArrayToHex(xorBuffer(input1, input2)); result != expectedResult {
		t.Errorf("Expected '%v', got '%v' :(", result, expectedResult)
	}
}
