package CTCSPacks

import "TransponderMsgParse/packets"

type Ctcs4 struct {
	packets.CTCS_Head

	D_TURNOUT uint16
	V_TURNOUT uint16
}

func (l Ctcs4) Encode() ([]byte, error) {
	panic("implement me")
}

func (l Ctcs4) Decode([]byte) error {
	panic("implement me")
}
func init() {
	packets.RegisterPacket("000000100", &Ctcs4{})
}
