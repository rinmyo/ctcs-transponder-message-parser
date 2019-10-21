package ETCSPacks

import "TransponderMsgParse/packets"

type Etcs132 struct {
	packets.ETCS_Head

	Q_ASPECT uint16
}

func (s Etcs132) Encode() ([]byte, error) {
	panic("implement me")
}

func (s Etcs132) Decode([]byte) error {
	panic("implement me")
}

func init() {
	packets.RegisterPacket("10000100", &Etcs132{})
}
