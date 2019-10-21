package ETCSPacks

import "TransponderMsgParse/packets"

type Etcs21 struct {
	packets.ETCS_Head

	D_GRADIENT uint16
	Q_GDIR     uint16
	G_A        uint16

	K []struct {
		D_GRADIENT_K uint16
		Q_GDIR_K     uint16
		G_A_K        uint16
	}
}

func (l Etcs21) Encode() ([]byte, error) {
	panic("implement me")
}

func (l Etcs21) Decode([]byte) error {
	panic("implement me")
}

func init() {
	packets.RegisterPacket("00010101", &Etcs21{})
}
