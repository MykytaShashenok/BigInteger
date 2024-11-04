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

	if carry > 0 {
		result.parts[len(result.parts)-1] += uint32(carry)
	}

	return result, carry
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
	hexString1 := "FFFFFFFF"
	hexString2 := "FFFFFFFF"

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

	fmt.Printf("BigInt1 parts: %v\n", b1.parts)
	fmt.Printf("BigInt2 parts: %v\n", b2.parts)

	result, carry := b1.Add(b2)

	fmt.Printf("Сумма: %v\n", result.parts)
	fmt.Printf("Перенос: %d\n", carry)

	fmt.Printf("Полный результат (включая перенос):\n")
	for i := 0; i < len(result.parts); i++ {
		fmt.Printf("Part %d: %d\n", i, result.parts[i])
	}
	if carry > 0 {
		fmt.Printf("Carry: %d\n", carry)
	}

	hexResult := result.ToHexString()
	fmt.Printf("Результат в шестнадцатеричном формате: %s\n", hexResult)
}
