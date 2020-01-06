package main

import "testing"

func TestHammingDistance(t *testing.T) {
	input1 := []byte("this is a test")
	input2 := []byte("wokka wokka!!!")
	expectedResult := 37
	if result, _ := hammingDistance(input1, input2); result != expectedResult {
		t.Errorf("Expected '%v', got '%v' :(", expectedResult, result)
	}
}
