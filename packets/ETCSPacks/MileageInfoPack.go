package ETCSPacks

import (
	"TransponderMsgParse/packets"
)

type Etcs79 struct {
	packets.ETCS_Head

	Q_NEWCOUNTRY uint16
	NID_C        uint16
	NID_BG       uint16
	D_POSOFF     uint16
	Q_MPOSITION  uint16
	M_POSITION   uint16

	N_ITER uint16

	K []struct {
		Q_NEWCOUNTRY uint16
		NID_C        uint16
		NID_BG       uint16
		D_POSOFF     uint16
		Q_MPOSITION  uint16
		M_POSITION   uint16
	}
}

func init() {
	packets.RegisterPacket("01001111", &Etcs79{})
}

func (s Etcs79) Encode() ([]byte, error) {
	panic("implement me")
}

func (s *Etcs79) Decode(binSlice []byte) {
	d := []uint16{8, 2, 13, 2, 1, 10, 14, 15, 1, 20, 5} //劃分
	p := packets.GetPieces(binSlice[:], d)              //切片

	//定長部分
	s.NID_PACKET, s.Q_DIR, s.L_PACKET, s.Q_SCALE, //頭
		s.Q_NEWCOUNTRY, s.NID_C, s.NID_BG, s.D_POSOFF, s.Q_MPOSITION, s.M_POSITION,
		s.N_ITER =
		p[0], p[1], p[2], p[3],
		p[4], p[5], p[6], p[7], p[8], p[9],
		p[10]

	s.K = make([]struct {
		Q_NEWCOUNTRY uint16
		NID_C        uint16
		NID_BG       uint16
		D_POSOFF     uint16
		Q_MPOSITION  uint16
		M_POSITION   uint16
	}, s.N_ITER)

	//變長部分
	for i := uint16(0); i < s.N_ITER; i++ {
		d1 := []uint16{1, 10, 14, 15, 1, 20}
		p1 := packets.GetPieces(binSlice[packets.Sum(d):], d1)

		s.K[i].Q_NEWCOUNTRY, s.K[i].NID_C, s.K[i].NID_BG, s.K[i].D_POSOFF, s.K[i].Q_MPOSITION, s.K[i].M_POSITION =
			p1[0], p1[1], p1[2], p1[3], p1[4], p1[5]
	}
}
