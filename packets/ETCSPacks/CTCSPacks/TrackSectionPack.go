package CTCSPacks

import "TransponderMsgParse/packets"

type Ctcs1 struct {
	packets.CTCS_Head

	D_SIGNAL uint16

	NID_SIGNAL    uint16
	NID_FREQUENCY uint16
	L_SECTION     uint16

	K []struct {
		NID_SIGNAL_K    uint16
		NID_FREQUENCY_K uint16
		L_SECTION_K     uint16
	}
}

func (t Ctcs1) Encode() ([]byte, error) {
	panic("implement me")
}

func (t Ctcs1) Decode([]byte) error {
	panic("implement me")
}
func init() {
	packets.RegisterPacket("000000001", &Ctcs1{})
}
