package packets

func Invert(str string) (bytes []byte) {
	bytes = []byte(str)
	for i, j := range str {
		bytes[i] = byte(j - '0')
	}
	return
}

func Bytes2Uint(bytes []byte) (r uint16) {
	for i, x := range bytes {
		r = r + uint16(x<<(len(bytes)-i-1))
	}
	return
}

func Uint2Bytes(num uint16, bytes *[]byte) {
	for i, _ := range *bytes {
		p := len(*bytes) - i - 1
		(*bytes)[i] = byte(num & uint16(1<<p) >> p)
	}
}

func GetPieces(bytes []byte, num []uint16) (result []uint16) {
	result = make([]uint16, len(num))
	p := 0
	for i, x := range num {
		result[i] = Bytes2Uint(bytes[p : p+int(x)])
	}
	return
}
