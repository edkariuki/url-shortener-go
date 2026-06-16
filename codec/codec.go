package codec

import (
	"fmt"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Encode(num uint64) string {
	if num == 0 {
		return string(alphabet[0])
	}

	var code []byte
	base := uint64(len(alphabet))

	for num > 0 {
		code = append([]byte{alphabet[num%base]}, code...)
		num /= base
	}

	return string(code)
}

func Decode(code string) (uint64, error) {
	var num uint64
	base := uint64(len(alphabet))

	for _, char := range code {
		pos := strings.IndexRune(alphabet, char)
		if pos == -1 {
			return 0, fmt.Errorf("invalid character: %c", char)
		}
		num = (num * base) + uint64(pos)
	}

	return num, nil
}
