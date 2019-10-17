package ETCSPacks

import "TransponderMsgParse/packets"

type LineSpeed struct {
	Part1 packets.ETCS_Head

	Part2 struct{
		D_STATIC  uint16
		V_STATIC  uint16
		Q_FRONT   uint16
		N_ITER    uint16
		NC_DIFF_N uint16
		V_DIFF_N  uint16
	}

	Part3 struct{
		N_ITER      uint16
		D_STATIC_K  uint16
		V_STATIC_K  uint16
		Q_FRONT_K   uint16
		N_ITER_K    uint16
		NC_DIFF_K_M uint16
		V_DIFF_K_M  uint16
	}
}

