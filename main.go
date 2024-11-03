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

func (b1 BigInt) Add(b2 BigInt) (BigInt, uint32) {
	var carry uint32
	var result BigInt

	for i := 0; i < len(b1.parts); i++ {
		// Сложение с учетом переноса
		temp := b1.parts[i] + b2.parts[i] + carry
		// result.parts[i] = temp & 0xFFFFFFFF
		result.parts[i] = temp
		carry = uint32(temp >> 32) // перенос, если превышает 32 бита
	}

	return result, carry
}

func main() {
	// Пример использования
	hexString1 := "11100000000ffffff"                // Первая hex-строка
	hexString2 := "ffffffffffffffffffffffffffffffff" // Вторая hex-строка

	bigInt1, err := HexToBigInt(hexString1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	bigInt2, err := HexToBigInt(hexString2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Сложение двух BigInt
	sum, carry := bigInt1.Add(bigInt2)

	// Печать результата
	fmt.Printf("BigInt1 parts: %v\n", bigInt1.parts)
	fmt.Printf("BigInt2 parts: %v\n", bigInt2.parts)
	fmt.Printf("Сумма: %v\n", sum.parts)
	fmt.Printf("Перенос: %d\n", carry)
}
