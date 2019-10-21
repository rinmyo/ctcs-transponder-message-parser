package ETCSPacks

import "TransponderMsgParse/packets"

type Etcs68 struct {
	packets.ETCS_Head

	Q_TRACKINIT uint16
	D_TRACKINIT uint16
	D_TRACKCOND uint16
	L_TRACKCOND uint16
	M_TRACKCOND uint16

	K []struct {
		D_TRACKCOND_K uint16
		L_TRACKCOND_K uint16
		M_TRACKCOND_K uint16
	}
}

func (s Etcs68) Encode() ([]byte, error) {
	panic("implement me")
}

func (s Etcs68) Decode([]byte) error {
	panic("implement me")
}

func init() {
	packets.RegisterPacket("01000100", &Etcs68{})
}
