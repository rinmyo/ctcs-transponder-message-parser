package ETCSPacks

import "TransponderMsgParse/packets"

type Etcs27 struct {
	packets.ETCS_Head

	D_STATIC uint16
	V_STATIC uint16
	Q_FRONT  uint16

	N_ITER_N uint16

	N []struct {
		NC_DIFF uint16
		V_DIFF  uint16
	}

	N_ITER_K uint16

	K []struct {
		D_STATIC uint16
		V_STATIC uint16
		Q_FRONT  uint16

		N_ITER_M uint16

		M []struct {
			NC_DIFF uint16
			V_DIFF  uint16
		}
	}
}

func (s Etcs27) Encode() ([]byte, error) {
	panic("implement me")
}

func (s *Etcs27) Decode(binSlice []byte) {
	// 設置頭
	d := []uint16{8, 2, 13, 2, 15, 7, 1, 5} //擷取定長部分
	p := packets.GetPieces(binSlice[:], d)

	s.NID_PACKET, s.Q_DIR, s.L_PACKET, s.Q_SCALE,
		s.D_STATIC, s.V_STATIC, s.Q_FRONT,
		s.N_ITER_N =
		p[0], p[1], p[2], p[3],
		p[4], p[5], p[6],
		p[7] //定長部分

	d1 := []uint16{4, 7}
	for n := uint16(0); n < s.N_ITER_N; n++ {
		p1 := packets.GetPieces(binSlice[packets.Sum(d)+packets.Sum(d1)*n:], d1)

		s.N[n].NC_DIFF, s.N[n].V_DIFF =
			p1[0], p1[1]
	}

	pointer := packets.Sum(d) + packets.Sum(d1)*s.N_ITER_N + 5 //指針
	d2 := []uint16{15, 7, 1, 5}

	for k := uint16(0); k < s.N_ITER_K; k++ {
		p2 := packets.GetPieces(binSlice[pointer+packets.Sum(d2)*k:], d2)
		s.K[k].D_STATIC, s.K[k].V_STATIC, s.K[k].Q_FRONT,
			s.K[k].N_ITER_M =
			p2[0], p2[1], p2[2],
			p2[3]

		pointer := pointer + packets.Sum(d2)
		d3 := []uint16{4, 7}
		for m := uint16(0); m < s.K[k].N_ITER_M; m++ {
			p3 := packets.GetPieces(binSlice[pointer+packets.Sum(d3):], d3)
			s.K[k].M[m].NC_DIFF, s.K[k].M[m].V_DIFF = p3[0], p3[1]
		}
	}

}

func init() {
	packets.RegisterPacket("00011011", &Etcs27{})
}
