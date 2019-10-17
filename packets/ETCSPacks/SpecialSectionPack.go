package ETCSPacks

import "TransponderMsgParse/packets"

type SpecialSectionPack struct {
	Part1 packets.ETCS_Head

	Part2 struct {
		Q_TRACKINIT uint16
		D_TRACKINIT uint16
		D_TRACKCOND uint16
		L_TRACKCOND uint16
		M_TRACKCOND uint16
	}

	Part3 struct {
		N_ITER        uint16
		D_TRACKCOND_K uint16
		L_TRACKCOND_K uint16
		M_TRACKCOND_K uint16
	}
}
