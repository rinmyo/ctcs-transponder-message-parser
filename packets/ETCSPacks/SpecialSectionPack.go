package ETCSPacks

import "TransponderMsgParse/packets"

type Etcs68 struct {
	packets.ETCS_Head

	Q_TRACKINIT uint16
	D_TRACKINIT uint16
	D_TRACKCOND uint16
	L_TRACKCOND uint16
	M_TRACKCOND uint16

	N_ITER uint16

	K []struct {
		D_TRACKCOND uint16
		L_TRACKCOND uint16
		M_TRACKCOND uint16
	}
}

func init() {
	packets.RegisterPacket("01000100", &Etcs68{})
}

func (s Etcs68) Encode() ([]byte, error) {
	panic("implement me")
}

func (s *Etcs68) Decode(binSlice []byte) {
	d := []uint16{8, 2, 13, 2, 1, 15, 15, 15, 4, 5} //擷取定長部分
	p := packets.GetPieces(binSlice[:], d)

	s.NID_PACKET, s.Q_DIR, s.L_PACKET, s.Q_SCALE,
		s.Q_TRACKINIT, s.D_TRACKINIT, s.D_TRACKCOND, s.L_TRACKCOND, s.M_TRACKCOND,
		s.N_ITER =
		p[0], p[1], p[2], p[3],
		p[4], p[5], p[6], p[7], p[8],
		p[9]

	s.K = make([]struct {
		D_TRACKCOND uint16
		L_TRACKCOND uint16
		M_TRACKCOND uint16
	}, s.N_ITER)

	for i := uint16(0); i < s.N_ITER; i++ {
		d1 := []uint16{15, 15, 4}
		p1 := packets.GetPieces(binSlice[packets.Sum(d)+packets.Sum(d1)*i:], d1)
		s.K[i].D_TRACKCOND, s.K[i].L_TRACKCOND, s.K[i].M_TRACKCOND = p1[0], p1[1], p1[2]
	}
}
