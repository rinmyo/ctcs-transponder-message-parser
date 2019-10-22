package packets

func Invert(str string) (bytes []byte) {
	bytes = []byte(str)
	for i, j := range str {
		bytes[i] = byte(j - '0')
	}
	return
}

func BINSlice2Uint(binSlice []byte) (r uint16) {
	for i, x := range binSlice {
		r = r + uint16(int(x)<<(len(binSlice)-i-1))
	}

	return
}

func Uint2BINSlice(num uint16, binSlice *[]byte) {
	for i, _ := range *binSlice {
		p := len(*binSlice) - i - 1
		(*binSlice)[i] = byte(num & uint16(1<<p) >> p)
	}
}

func GetPieces(binSlice []byte, num []uint16) (result []uint16) {
	result = make([]uint16, len(num))
	p := 0
	for i, x := range num {
		result[i] = BINSlice2Uint(binSlice[p : p+int(x)])
		p += int(x)
	}
	return
}

func Sum(d []uint16) (sum uint16) {
	for _, x := range d {
		sum += x
	}
	return
}

func GetStr(binSlice []byte) (result string) {
	for _, x := range binSlice {
		result += string(x + 48)
	}
	return
}
