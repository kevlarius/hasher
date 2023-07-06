package main

import (
	"bytes"
	"encoding/hex"
	"testing"
)

var TestData = []byte{
	0x00, 0x10, 0x20, 0x30, 0x40, 0x50, 0x60, 0x70,
	0x80, 0x80, 0xA0, 0xB0, 0xC0, 0xD0, 0xE0, 0xF0,
}

func TestCalcHashSumSha1(t *testing.T) {
	inputFile := bytes.NewReader(TestData)
	result := hex.EncodeToString(calcHashSum(inputFile, SHA1))
	expectedResult := "1e0a8306e1bb0535dc4fe2c939f50463322cff44"
	if result != expectedResult {
		t.Error("Expected: ", expectedResult, ", returned: ", result)
	}
}

func TestCalcHashSumSha256(t *testing.T) {
	inputFile := bytes.NewReader(TestData)
	result := hex.EncodeToString(calcHashSum(inputFile, SHA256))
	expectedResult := "2b9fda2c6249a2bf4e00616d7cf0de995f5863ad4962fe9b7c6af458e27af966"
	if result != expectedResult {
		t.Error("Expected: ", expectedResult, ", returned: ", result)
	}
}

func TestCalcHashSumSha512(t *testing.T) {
	inputFile := bytes.NewReader(TestData)
	result := hex.EncodeToString(calcHashSum(inputFile, SHA512))
	expectedResult := "01634a3ba27c04f751acc6427a9abac216e08fd2f1e3fe72e26f43aa7e24e065d80a1911881511d3e7539f9e44" +
		"70de2f15573e9cf5d6c1f3f04ee87c2902f2b2"
	if result != expectedResult {
		t.Error("Expected: ", expectedResult, ", returned: ", result)
	}
}

func TestCalcHashSumShaMd5(t *testing.T) {
	inputFile := bytes.NewReader(TestData)
	result := hex.EncodeToString(calcHashSum(inputFile, MD5))
	expectedResult := "9e0c9dad93a900ef7dfcb647cb0d5ccd"
	if result != expectedResult {
		t.Error("Expected: ", expectedResult, ", returned: ", result)
	}
}

func TestCalcHashSumShaCRC32(t *testing.T) {
	inputFile := bytes.NewReader(TestData)
	result := hex.EncodeToString(calcHashSum(inputFile, CRC32))
	expectedResult := "c1ac4088"
	if result != expectedResult {
		t.Error("Expected: ", expectedResult, ", returned: ", result)
	}
}
