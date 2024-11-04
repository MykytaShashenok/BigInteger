package utils

import (
	"biginteger"
	"encoding/hex"
	"fmt"
)

// Функція для перетворення шестнадцятеричної строки в BigInt з урахуванням little endian
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

// Функція для перетворення BigInt в шестнадцятеричну строку з урахуванням little endian
func (b BigInt) ToHexString() string {
	// Створюємо зріз для зберігання шестнадцятеричного представлення
	hexParts := make([]byte, 0, 32) // 4 частини по 8 символів

	for i := 0; i < len(b.parts); i++ {
		// Перетворюємо кожну частину в шестнадцятеричне представлення
		partHex := fmt.Sprintf("%08x", b.parts[i])
		hexParts = append(hexParts, partHex...)
	}

	// Реверсуємо зріз для little endian
	for i, j := 0, len(hexParts)-1; i < j; i, j = i+1, j-1 {
		hexParts[i], hexParts[j] = hexParts[j], hexParts[i]
	}

	return string(hexParts)
}
