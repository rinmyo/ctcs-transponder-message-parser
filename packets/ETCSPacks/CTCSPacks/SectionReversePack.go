package CTCSPacks

import "TransponderMsgParse/packets"

type Ctcs3 struct {
	packets.UserInfoPacket

	packets.CTCS_Head

	D_STARTREVERSE uint16
	L_REVERSEAREA  uint16
}

func (s Ctcs3) Encode() ([]byte, error) {
	panic("implement me")
}

func (s Ctcs3) Decode(binSlice []byte) []byte {
	return binSlice[s.Length:]
}

func init() {
	packets.RegisterPacket("000000011", &Ctcs3{})
}
