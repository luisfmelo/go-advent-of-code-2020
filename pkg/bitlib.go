package pkg

// SetBit sets the bit in number `n`, at position `pos` to 1. It will return the number changed.
func SetBit(n int, pos uint) int {
	n |= 1 << pos
	return n
}

// ClearBit clears the bit in number `n`, at position `pos` to 0. It will return the number changed.
func ClearBit(n int, pos uint) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}

// ForceBit will change the bit in number `n`, at position `pos` to value `bitValue`. It will return the number changed.
func ForceBit(n int, pos uint, bitValue int) int {
	if bitValue == 1 {
		return SetBit(n, pos)
	}
	return ClearBit(n, pos)
}
