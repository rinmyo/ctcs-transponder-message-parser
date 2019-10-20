package ETCSPacks

import "TransponderMsgParse/packets"

const Etcs44Nid = 0b00101100

type CTCSDataPack struct {
	packets.ETCS_Head

	NID_XUSER uint16
	XXXXXX    interface{}
}

func NewCTCSDataPack() *CTCSDataPack {
	return &CTCSDataPack{
		ETCS_Head: packets.ETCS_Head{
			NID_PACKET: Etcs44Nid,
		},
	}
}
