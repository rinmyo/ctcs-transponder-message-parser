package ETCSPacks

import (
	"TransponderMsgParse/packets"
)

type Etcs44 struct {
	Head struct {
		NID_PACKET uint16 `json:"nid_packet"`
		Q_DIR      uint16 `json:"q_dir"`
		L_PACKET   uint16 `json:"l_packet"`
	}

	NID_XUSER uint16
	XXXXXX    packets.ICtcsPack
}

func (s Etcs44) Encode() ([]byte, error) {
	panic("implement me")
}

func (s *Etcs44) Decode(binSlice []byte) error {
	d := []uint16{8, 2, 13, 9} //擷取定長部分
	p := packets.GetPieces(binSlice[:], d)

	s.Head.NID_PACKET, s.Head.Q_DIR, s.Head.L_PACKET, s.NID_XUSER =
		p[0], p[1], p[2], p[3]

	//Route the CTCS packet inside the ETCS-44
	s.XXXXXX = packets.GetPacket(string(binSlice[23 : 23+9]))
	err := s.XXXXXX.Decode(binSlice[23:])
	return err
}

func init() {
	packets.RegisterPacket("00101100", &Etcs44{})
}
