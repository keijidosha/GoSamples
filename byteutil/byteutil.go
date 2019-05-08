package byteutil

type ByteArray []byte

func (buf ByteArray) ToUIntLE(start int, len int) uint {
	var sum uint = 0
	max := start + len
	var pos uint = 0
	for idx := start; idx < max; idx++ {
		sum += uint(buf[idx]) << (pos * 8)
		pos++
	}
	return sum
}

func (buf ByteArray) ToUIntBE(start int, len int) uint {
	var sum uint = 0
	max := start + len
	var pos uint = 0
	for idx := max - 1; idx >= 0; idx-- {
		sum += uint(buf[idx]) << (pos * 8)
		pos++
	}
	return sum
}

func UIntToBytesLE(value uint, len int) []byte {
	buf := make([]byte, len, len)
	remain := value

	for ii := 0; ii < len; ii++ {
		v := remain & 0xff
		buf[ii] = byte(v)
		remain = remain >> 8
	}

	return buf
}

func UIntToBytesBE(value uint, len int) []byte {
	buf := make([]byte, len, len)
	remain := value
	maxIdx := len - 1

	for ii := maxIdx; ii >= 0; ii-- {
		v := remain & 0xff
		buf[ii] = byte(v)
		remain = remain >> 8
	}

	return buf
}
