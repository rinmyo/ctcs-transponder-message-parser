package CTCSPacks

import "TransponderMsgParse/packets"

type Ctcs4 struct {
	packets.UserInfoPacket

	packets.CTCS_Head

	D_TURNOUT uint16
	V_TURNOUT uint16
}

func (s Ctcs4) GetLength() uint16 {
	return s.Length
}

func (s Ctcs4) GetNextPack() *packets.IEtcsPack {
	return &s.NextPack
}

func (s Ctcs4) Encode() ([]byte, error) {
	panic("implement me")
}

func (s Ctcs4) Decode([]byte) {
	panic("implement me")
}
func init() {
	packets.RegisterPacket("000000100", &Ctcs4{})
}
