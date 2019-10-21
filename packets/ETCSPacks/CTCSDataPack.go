package ETCSPacks

import (
	"TransponderMsgParse/packets"
	"fmt"
)

type Etcs44 struct {
	packets.UserInfoPacket

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

func (s *Etcs44) Decode(binSlice []byte) {
	d := []uint16{8, 2, 13, 9} //擷取定長部分
	p := packets.GetPieces(binSlice[:], d)

	s.Head.NID_PACKET, s.Head.Q_DIR, s.Head.L_PACKET, s.NID_XUSER =
		p[0], p[1], p[2], p[3]
	s.Length += packets.Sum(d)
	s.XXXXXX = packets.GetPacket(packets.GetStr(binSlice[23 : 23+9]))
	s.XXXXXX.Decode(binSlice[23:])

	s.Length += s.XXXXXX.GetLength() - 9 //重複了9位

	s.NextPack = packets.GetPacket(packets.GetStr(binSlice[s.Length : s.Length+8]))
	fmt.Println(packets.GetStr(binSlice[s.Length : s.Length+8]))
	s.NextPack.Decode(binSlice[s.Length:])
}

func init() {
	packets.RegisterPacket("00101100", &Etcs44{})
}
