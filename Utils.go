package main

func BIN2DEC(bin string) (r int) {
	for i, x := range bin {
		r = r + int(x-'0')<<(len(bin)-i-1)
	}

	return
}

func DEC2BIN(num int, binSlice *[]byte) {
	for i, _ := range *binSlice {
		p := len(*binSlice) - i - 1
		(*binSlice)[i] = byte(num & (1 << p) >> p)
	}
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
