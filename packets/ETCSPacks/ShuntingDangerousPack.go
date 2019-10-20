package ETCSPacks

import "TransponderMsgParse/packets"

const Etcs132Nid = 0b10000100

type ShuntingDangerousPack struct {
	packets.ETCS_Head

	Q_ASPECT uint16
}

func (s ShuntingDangerousPack) Encode() ([]byte, error) {
	panic("implement me")
}

func (s ShuntingDangerousPack) Decode([]byte) error {
	panic("implement me")
}
