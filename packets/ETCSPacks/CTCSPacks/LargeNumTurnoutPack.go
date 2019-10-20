package CTCSPacks

import "TransponderMsgParse/packets"

const Ctcs4Nid = 0b000000100

type LargeNumTurnoutPack struct {
	packets.CTCS_Head

	D_TURNOUT uint16
	V_TURNOUT uint16
}

func (l LargeNumTurnoutPack) Encode() ([]byte, error) {
	panic("implement me")
}

func (l LargeNumTurnoutPack) Decode([]byte) error {
	panic("implement me")
}
