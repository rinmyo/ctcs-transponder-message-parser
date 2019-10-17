package ETCSPacks

import "TransponderMsgParse/packets"

type TransponderLinkPack struct {
	Part1 packets.ETCS_Head

	Part2 struct {
		D_LINK            uint16
		Q_NEWCOUNTRY      uint16
		NID_C             uint16
		NID_BG            uint16
		Q_LINKORIENTATION uint16
		Q_LINKREACTION    uint16
		Q_LINKACC         uint16
	}

	Part3 struct {
		N_ITER              uint16
		D_LINK_K            uint16
		Q_NEWCOUNTRY_K      uint16
		NID_C_K             uint16
		NID_BG_K            uint16
		Q_LINKORIENTATION_K uint16
		Q_LINKREACTION_K    uint16
		Q_LINKACC_K         uint16
	}
}
