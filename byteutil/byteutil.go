package byteutil

type ByteArray []byte

func (buf ByteArray) ToUIntLE( start int, len int ) uint {
	var sum uint = 0
	max := start + len
	var pos uint = 0
	for idx := start; idx < max; idx++ {
		sum += uint( buf[idx] ) << (pos * 8)
		pos++
	}
	return sum
}
