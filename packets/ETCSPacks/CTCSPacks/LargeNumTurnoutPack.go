package CTCSPacks

import "TransponderMsgParse/packets"

type Ctcs4 struct {
	packets.UserInfoPacket

	packets.CTCS_Head

	D_TURNOUT uint16
	V_TURNOUT uint16
}

func (s Ctcs4) Encode() ([]byte, error) {
	panic("implement me")
}

func (s Ctcs4) Decode(binSlice []byte) []byte {
	return binSlice[s.Length:]
}

func init() {
	packets.RegisterPacket("000000100", &Ctcs4{})
}
