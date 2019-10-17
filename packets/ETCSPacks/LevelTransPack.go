package ETCSPacks

import "TransponderMsgParse/packets"

type LevelTransPack struct {
	Part1 packets.ETCS_Head

	Part2 struct {
		D_LEVELTR    uint16
		M_LEVELTR    uint16
		NID_STM      uint16
		L_ACKLEVELTR uint16
	}

	Part3 struct {
		N_ITER         uint16
		M_LEVELTR_K    uint16
		NID_STM_K      uint16
		L_ACKLEVELTR_K uint16
	}
}
