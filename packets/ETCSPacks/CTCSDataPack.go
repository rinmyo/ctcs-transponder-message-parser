package ETCSPacks

import "TransponderMsgParse/packets"

type CTCSDataPack struct {
	Part1 packets.ETCS_Head

	Part2 struct{
		NID_PACKET uint16
		Q_DIR uint16
		L_PACKET uint16
	}

	Part3 struct{
		NID_XUSER uint16
		XXXXXX interface{}
	}
}