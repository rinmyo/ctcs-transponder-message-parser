package CTCSPacks

import "TransponderMsgParse/packets"

type LargeNumTurnoutPack struct {
	Part1 packets.CTCS_Head

	Part2 struct {
		D_TURNOUT uint16
		V_TURNOUT uint16
	}
}
