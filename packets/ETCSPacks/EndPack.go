package ETCSPacks

import (
	"TransponderMsgParse/packets"
	"fmt"
)

func init() {
	packets.RegisterPacket("00000101", &Etcs5{})
}

type End struct {
}

func (e End) Encode() ([]byte, error) {
	panic("implement me")
}

func (e End) Decode(binSlice []byte) {
	fmt.Println("解析結束")
}

func (e End) GetLength() uint16 {
	panic("implement me")
}

func (e End) GetNextPack() *packets.IEtcsPack {
	panic("implement me")
}
