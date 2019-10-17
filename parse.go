package main

func divide(msg []byte) (frameMark []byte, userInfoPack []byte, infoEnd []byte) {
	return msg[0:50], msg[50 : 50+772], msg[50+772 : 50+772+8]
}



func invert(str string) (bytes []byte) {
	bytes = []byte(str)
	for i, j := range str {
		bytes[i] = byte(j - '0')
	}
	return
}

func Bytes2Uint(bytes []byte) (r uint16) {
	for i, x := range bytes {
		r = r + uint16( x<<( len( bytes) - i - 1) )
	}
	return
}
