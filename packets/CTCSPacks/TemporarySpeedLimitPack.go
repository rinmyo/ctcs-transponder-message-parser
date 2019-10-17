package CTCSPacks

import "TransponderMsgParse/packets"

type TemporarySpeedLimitPack struct {
	Part1 packets.CTCS_Head

	Part2 struct {
		L_TSRarea uint16
	}

	Part3 struct {
		D_TSR   uint16
		L_TSR   uint16
		Q_FRONT uint16
		V_TSR   uint16
	}

	Part4 struct {
		N_ITER    uint16
		D_TSR_N   uint16
		L_TSR_N   uint16
		Q_FRONT_N uint16
		V_TSR     uint16
	}
}
