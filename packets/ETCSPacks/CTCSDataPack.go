package ETCSPacks

import (
	"TransponderMsgParse/packets"
	"TransponderMsgParse/packets/ETCSPacks/CTCSPacks"
)

const Etcs44Nid = 0b00101100

type CTCSDataPack struct {
	Head struct {
		NID_PACKET uint16 `json:"nid_packet"`
		Q_DIR      uint16 `json:"q_dir"`
		L_PACKET   uint16 `json:"l_packet"`
	}

	NID_XUSER uint16
	XXXXXX    packets.ICtcsPack
}

func (C CTCSDataPack) Encode() ([]byte, error) {
	panic("implement me")
}

func (C *CTCSDataPack) Decode(bytes []byte) error {
	// 設置頭
	h := packets.GetPieces(bytes, []uint16{8, 2, 13})
	C.Head = struct {
		NID_PACKET uint16 `json:"nid_packet"`
		Q_DIR      uint16 `json:"q_dir"`
		L_PACKET   uint16 `json:"l_packet"`
	}{
		h[0], h[1], h[2],
	}

	//Route the CTCS packet inside the ETCS-44
	switch nid := packets.Bytes2Uint(bytes[23 : 23+9]); nid {
	case CTCSPacks.Ctcs1Nid:
		C.XXXXXX = &CTCSPacks.TrackSectionPack{}
	case CTCSPacks.Ctcs2Nid:
		C.XXXXXX = &CTCSPacks.TemporarySpeedLimitPack{}
	case CTCSPacks.Ctcs3Nid:
		C.XXXXXX = &CTCSPacks.SectionReversePack{}
	case CTCSPacks.Ctcs4Nid:
		C.XXXXXX = &CTCSPacks.LargeNumTurnoutPack{}
	}
	err := C.XXXXXX.Decode(bytes[23:])

	return err
}
