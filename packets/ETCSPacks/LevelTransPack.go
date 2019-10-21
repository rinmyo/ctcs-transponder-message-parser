package ETCSPacks

import "TransponderMsgParse/packets"

type Etcs41 struct {
	packets.UserInfoPacket

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

func (s Etcs41) Encode() ([]byte, error) {
	panic("implement me")
}

func (s Etcs41) Decode([]byte) {
	panic("implement me")
}
func init() {
	packets.RegisterPacket("00101001", &Etcs41{})
}
