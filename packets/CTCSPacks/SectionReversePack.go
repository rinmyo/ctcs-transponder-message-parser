package CTCSPacks

import "TransponderMsgParse/packets"

type SectionReversePack struct {
	Part1 packets.CTCS_Head

	Part2 struct {
		D_STARTREVERSE uint16
		L_REVERSEAREA  uint16
	}
}
