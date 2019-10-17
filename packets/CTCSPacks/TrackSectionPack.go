package CTCSPacks

import "TransponderMsgParse/packets"

type TrackSectionPack struct {
	Part1 packets.CTCS_Head

	Part2 struct {
		D_SIGNAL uint16
	}

	Part3 struct {
		NID_SIGNAL    uint16
		NID_FREQUENCY uint16
		L_SECTION     uint16
	}

	Part4 struct {
		N_ITER          uint16
		NID_SIGNAL_K    uint16
		NID_FREQUENCY_K uint16
		L_SECTION_K     uint16
	}
}
