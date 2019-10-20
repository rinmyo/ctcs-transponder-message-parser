package CTCSPacks

import "TransponderMsgParse/packets"

const Ctcs3Nid = 0b000000011

type SectionReversePack struct {
	packets.CTCS_Head

	D_STARTREVERSE uint16
	L_REVERSEAREA  uint16
}
