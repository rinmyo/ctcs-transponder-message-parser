package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	str, err := ioutil.ReadFile("in.txt")
	if err != nil {
		fmt.Print(err)
	}

	a := ""
	for i := 0; i < 1000; i++ {
		a += "0"
	}

	b := NewBinMessage((string)(str))
	fmt.Println(b.Decode2FrameMark())

	c := Decode2EtcsPacket(b.body)
	for _, x := range c {
		fmt.Println(x)
	}
}
