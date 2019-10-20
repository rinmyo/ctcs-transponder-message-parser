package CTCSPacks

import "TransponderMsgParse/packets"

const Ctcs2Nid = 0b000000010

type TemporarySpeedLimitPack struct {
	packets.CTCS_Head

	L_TSRarea uint16

	D_TSR   uint16
	L_TSR   uint16
	Q_FRONT uint16
	V_TSR   uint16

	K []struct {
		D_TSR_N   uint16
		L_TSR_N   uint16
		Q_FRONT_N uint16
		V_TSR     uint16
	}
}
