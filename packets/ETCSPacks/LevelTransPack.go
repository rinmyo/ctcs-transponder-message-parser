package ETCSPacks

import "TransponderMsgParse/packets"

const Etcs41Nid = 0b00101001

type LevelTransPack struct {
	packets.ETCS_Head

	D_LEVELTR    uint16
	M_LEVELTR    uint16
	NID_STM      uint16
	L_ACKLEVELTR uint16

	K []struct {
		M_LEVELTR_K    uint16
		NID_STM_K      uint16
		L_ACKLEVELTR_K uint16
	}
}

func (l LevelTransPack) Encode() ([]byte, error) {
	panic("implement me")
}

func (l LevelTransPack) Decode([]byte) error {
	panic("implement me")
}
