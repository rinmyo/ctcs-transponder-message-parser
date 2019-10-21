package CTCSPacks

import "TransponderMsgParse/packets"

type Ctcs1 struct {
	packets.CTCS_Head

	D_SIGNAL uint16

	NID_SIGNAL    uint16
	NID_FREQUENCY uint16
	L_SECTION     uint16

	N_ITER uint16

	K []struct {
		NID_SIGNAL    uint16
		NID_FREQUENCY uint16
		L_SECTION     uint16
	}
}

func init() {
	packets.RegisterPacket("000000001", &Ctcs1{})
}

func (s Ctcs1) Encode() ([]byte, error) {
	panic("implement me")
}

func (s Ctcs1) Decode(binSlice []byte) error {
	d := []uint16{9, 2, 13, 2, 15, 4, 5, 15, 5}
	p := packets.GetPieces(binSlice[:], d)

	//定長部分
	s.NID_XUSER, s.Q_DIR, s.L_PACKET, s.Q_SCALE,
		s.D_SIGNAL, s.NID_SIGNAL, s.NID_FREQUENCY, s.L_SECTION,
		s.N_ITER =
		p[0], p[1], p[2], p[3],
		p[4], p[5], p[6], p[7],
		p[8]

	//變長部分
	for k := uint16(0); k < s.N_ITER; k++ {
		d1 := []uint16{4, 5, 15}
		p1 := packets.GetPieces(binSlice[packets.Sum(d)+packets.Sum(d1)*k:], d1)
		s.K[k].NID_SIGNAL, s.K[k].NID_FREQUENCY, s.K[k].L_SECTION = p1[0], p1[1], p1[2]
	}

	return nil

}
