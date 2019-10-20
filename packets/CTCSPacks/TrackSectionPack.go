package CTCSPacks

import "TransponderMsgParse/packets"

const Ctcs1Nid = 0b000000001

type TrackSectionPack struct {
	packets.CTCS_Head

	D_SIGNAL uint16

	NID_SIGNAL    uint16
	NID_FREQUENCY uint16
	L_SECTION     uint16

	K []struct {
		NID_SIGNAL_K    uint16
		NID_FREQUENCY_K uint16
		L_SECTION_K     uint16
	}
}
