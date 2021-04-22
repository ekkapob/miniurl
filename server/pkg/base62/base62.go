package base62

const (
	chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	base  = uint64(len(chars))
)

func Encode(num uint64) string {
	encoded := ""
	for num > 0 {
		r := num % base
		num /= base
		encoded = string(chars[r]) + encoded
	}
	return encoded
}
