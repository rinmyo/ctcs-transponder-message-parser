package ETCSPacks

import (
	"TransponderMsgParse/packets"
)

const Etcs5Nid = 0b00000101

type TransponderLinkPack struct {
	packets.ETCS_Head

	D_LINK            uint16
	Q_NEWCOUNTRY      uint16
	NID_C             uint16
	NID_BG            uint16
	Q_LINKORIENTATION uint16
	Q_LINKREACTION    uint16
	Q_LOCACC          uint16

	K []struct {
		D_LINK            uint16
		Q_NEWCOUNTRY      uint16
		NID_C             uint16
		NID_BG            uint16
		Q_LINKORIENTATION uint16
		Q_LINKREACTION    uint16
		Q_LOCACC          uint16
	}
}

// constructor
func NewTransponderLinkPack() *TransponderLinkPack {
	return &TransponderLinkPack{
		ETCS_Head: packets.ETCS_Head{
			NID_PACKET: Etcs5Nid,
		},
	}
}

func (t TransponderLinkPack) Encode() ([]byte, error) {
	panic("implement me")
}

func (t *TransponderLinkPack) Decode(bytes []byte) error {
	t.ETCS_Head = *packets.NewETCS_Head(bytes[0:25])
	cl := packets.GetPieces(bytes[25:], []uint16{15, 1, 10, 12, 1, 2, 6, 5})

	t.D_LINK, t.Q_NEWCOUNTRY, t.NID_C, t.NID_BG, t.Q_LINKORIENTATION, t.Q_LINKREACTION, t.Q_LOCACC =
		cl[0], cl[1], cl[2], cl[3], cl[4], cl[5], cl[6]

	n := cl[7] //求取有幾個變長部分

	for i := 0; i < int(n); i++ {
		cl = packets.GetPieces(bytes[25+49+5+49*i:], []uint16{15, 1, 10, 12, 1, 2, 6})
		t.K[i].D_LINK, t.K[i].Q_NEWCOUNTRY, t.K[i].NID_C, t.K[i].NID_BG, t.K[i].Q_LINKORIENTATION, t.K[i].Q_LINKREACTION, t.K[i].Q_LOCACC =
			cl[0], cl[1], cl[2], cl[3], cl[4], cl[5], cl[6]
	}

	return nil

}
