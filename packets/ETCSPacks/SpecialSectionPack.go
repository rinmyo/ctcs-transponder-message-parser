package ETCSPacks

import "TransponderMsgParse/packets"

const Etcs68Nid = 0b01000100

type SpecialSectionPack struct {
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

func (s SpecialSectionPack) Encode() ([]byte, error) {
	panic("implement me")
}

func (s SpecialSectionPack) Decode([]byte) error {
	panic("implement me")
}
