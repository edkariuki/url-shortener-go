package codec

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
