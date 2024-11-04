package main

import (
	"encoding/hex"
	"fmt"
)

type BigInt struct {
	parts [4]uint32
}

func HexToBigInt(hexStr string) (BigInt, error) {
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

func (b1 BigInt) Add(b2 BigInt) (BigInt, uint64) {
	var carry uint64 = 0
	var result BigInt

	for i := 0; i < len(b1.parts); i++ {
		temp := uint64(b1.parts[i]) + uint64(b2.parts[i]) + carry
		result.parts[i] = uint32(temp & 0xFFFFFFFF)
		carry = temp >> 32
	}

	return result, carry
}

func (b1 BigInt) Subb(b2 BigInt) (BigInt, uint32) {
	var borrow uint32 = 0
	var result BigInt

	for i := 0; i < len(b1.parts); i++ {
		temp := int64(b1.parts[i]) - int64(b2.parts[i]) - int64(borrow)
		borrow = uint32((temp >> 32) & 1)
		result.parts[i] = uint32(temp & 0xFFFFFFFF)
	}

	return result, borrow
}

func (b BigInt) ToHexString() string {
	hexParts := make([]byte, 0, 32)

	for i := 0; i < len(b.parts); i++ {
		partHex := fmt.Sprintf("%08x", b.parts[i])
		hexParts = append(hexParts, partHex...)
	}

	return string(hexParts)
}

func main() {
	hexString1 := "FFFFFFfffff"
	hexString2 := "123456ff123"

	b1, err := HexToBigInt(hexString1)
	if err != nil {
		fmt.Println("Error converting hex to BigInt:", err)
		return
	}

	b2, err := HexToBigInt(hexString2)
	if err != nil {
		fmt.Println("Error converting hex to BigInt:", err)
		return
	}

	result, borrow := b1.Subb(b2)
	fmt.Printf("Результат вычитания: %v\n", result.parts)
	fmt.Printf("Заимствование: %d\n", borrow)

	hexResult := result.ToHexString()
	fmt.Printf("Результат в шестнадцатеричном формате: %s\n", hexResult)
}
