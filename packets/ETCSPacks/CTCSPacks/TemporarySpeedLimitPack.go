package CTCSPacks

import "TransponderMsgParse/packets"

type Ctcs2 struct {
	packets.CTCS_Head

	L_TSRarea uint16

	D_TSR   uint16
	L_TSR   uint16
	Q_FRONT uint16
	V_TSR   uint16

	K []struct {
		D_TSR_N   uint16
		L_TSR_N   uint16
		Q_FRONT_N uint16
		V_TSR     uint16
	}
}

func (s Ctcs2) Encode() ([]byte, error) {
	panic("implement me")
}

func (s Ctcs2) Decode(binSlice []byte) {
}

func init() {
	packets.RegisterPacket("000000010", &Ctcs2{})
}
