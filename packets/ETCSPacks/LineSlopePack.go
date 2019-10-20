package ETCSPacks

import "TransponderMsgParse/packets"

const Etcs21Nid = 0b00010101

type LineSlopePack struct {
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

func (l LineSlopePack) Encode() ([]byte, error) {
	panic("implement me")
}

func (l LineSlopePack) Decode([]byte) error {
	panic("implement me")
}
