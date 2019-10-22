package ETCSPacks

import (
	"TransponderMsgParse/packets"
)

type Etcs5 struct {
	packets.UserInfoPacket

	packets.ETCS_Head

	D_LINK            uint16
	Q_NEWCOUNTRY      uint16
	NID_C             uint16
	NID_BG            uint16
	Q_LINKORIENTATION uint16
	Q_LINKREACTION    uint16
	Q_LOCACC          uint16

	N_ITER uint16
	K      []struct {
		D_LINK            uint16
		Q_NEWCOUNTRY      uint16
		NID_C             uint16
		NID_BG            uint16
		Q_LINKORIENTATION uint16
		Q_LINKREACTION    uint16
		Q_LOCACC          uint16
	}
}

func init() {
	packets.RegisterPacket("00000101", &Etcs5{})
}

func (s Etcs5) Encode() ([]byte, error) {
	panic("implement me")
}

func (s *Etcs5) Decode(binSlice []byte) []byte {
	d := []uint16{8, 2, 13, 2, 15, 1, 10, 14, 1, 2, 6, 5} //劃分
	p := packets.GetPieces(binSlice[:], d)                //切片

	//定長部分
	s.NID_PACKET, s.Q_DIR, s.L_PACKET, s.Q_SCALE, //頭
		s.D_LINK, s.Q_NEWCOUNTRY, s.NID_C, s.NID_BG, s.Q_LINKORIENTATION, s.Q_LINKREACTION, s.Q_LOCACC,
		s.N_ITER =
		p[0], p[1], p[2], p[3],
		p[4], p[5], p[6], p[7], p[8], p[9], p[10],
		p[11]
	s.Length += packets.Sum(d)

	s.K = make([]struct {
		D_LINK            uint16
		Q_NEWCOUNTRY      uint16
		NID_C             uint16
		NID_BG            uint16
		Q_LINKORIENTATION uint16
		Q_LINKREACTION    uint16
		Q_LOCACC          uint16
	}, s.N_ITER)

	//變長部分
	for i := uint16(0); i < s.N_ITER; i++ {
		d1 := []uint16{15, 1, 10, 14, 1, 2, 6}
		p1 := packets.GetPieces(binSlice[s.Length:], d1)
		s.K[i].D_LINK, s.K[i].Q_NEWCOUNTRY, s.K[i].NID_C, s.K[i].NID_BG, s.K[i].Q_LINKORIENTATION, s.K[i].Q_LINKREACTION, s.K[i].Q_LOCACC =
			p1[0], p1[1], p1[2], p1[3], p1[4], p1[5], p1[6]
		s.Length += packets.Sum(d1)
	}

	return binSlice[s.Length:]
}
