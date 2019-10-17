package ETCSPacks

import "TransponderMsgParse/packets"

type LineSlopePack struct {
	Part1 packets.ETCS_Head

	Part2 struct {
		D_GRADIENT uint16
		Q_GDIR     uint16
		G_A        uint16
	}

	Part3 struct {
		N_ITER       uint16
		D_GRADIENT_K uint16
		Q_GDIR_K     uint16
		G_A_K        uint16
	}
}
