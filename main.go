package main

import (
	"encoding/hex"
	"errors"
	"fmt"
)

type BigInt struct {
	parts [4]uint32
}

func HexToBigInt(hexStr string) (BigInt, error) {
	hexStr = removeSpaces(hexStr)
	if len(hexStr)%2 != 0 {
		hexStr = "0" + hexStr
	}

	if len(hexStr) > 32 {
		return BigInt{}, errors.New("hex string is too long")
	}

	data, err := hex.DecodeString(hexStr)
	if err != nil {
		return BigInt{}, err
	}

	var bigInt BigInt
	for i := 0; i < len(data); i++ {
		bigInt.parts[i/4] |= uint32(data[i]) << (8 * (i % 4))
	}

	return bigInt, nil
}

func removeSpaces(s string) string {
	result := ""
	for _, c := range s {
		if c != ' ' {
			result += string(c)
		}
	}
	return result
}

func main() {
	// Пример использования
	hexString := "1" // Пример hex-строки с нечётной длиной
	bigInt, err := HexToBigInt(hexString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Печать результата
	fmt.Printf("BigInt parts: %v\n", bigInt.parts)
}
