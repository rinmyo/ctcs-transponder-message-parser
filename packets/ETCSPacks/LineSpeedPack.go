package ETCSPacks

import "TransponderMsgParse/packets"

type Etcs27 struct {
	packets.ETCS_Head

	D_STATIC  uint16
	V_STATIC  uint16
	Q_FRONT   uint16
	N_ITER    uint16
	NC_DIFF_N uint16
	V_DIFF_N  uint16

	K []struct {
		D_STATIC_K  uint16
		V_STATIC_K  uint16
		Q_FRONT_K   uint16
		N_ITER_K    uint16
		NC_DIFF_K_M uint16
		V_DIFF_K_M  uint16
	}
}

func (l Etcs27) Encode() ([]byte, error) {
	panic("implement me")
}

func (l Etcs27) Decode([]byte) error {
	panic("implement me")
}

func init() {
	packets.RegisterPacket("00011011", &Etcs27{})
}
