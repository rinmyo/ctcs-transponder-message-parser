package ETCSPacks

import "TransponderMsgParse/packets"

type Etcs132 struct {
	packets.UserInfoPacket

	Head struct {
		NID_PACKET uint16 `json:"nid_packet"`
		Q_DIR      uint16 `json:"q_dir"`
		L_PACKET   uint16 `json:"l_packet"`
	}

	Q_ASPECT uint16
}

func (s Etcs132) Encode() ([]byte, error) {
	panic("implement me")
}

func (s *Etcs132) Decode(binSlice []byte) []byte {
	// 設置頭
	d := []uint16{8, 2, 13, 1}
	p := packets.GetPieces(binSlice, d)

	s.Head.NID_PACKET, s.Head.Q_DIR, s.Head.L_PACKET,
		s.Q_ASPECT =
		p[0], p[1], p[2],
		p[3]

	return binSlice[s.Length:]
}

func init() {
	packets.RegisterPacket("10000100", &Etcs132{})
}
