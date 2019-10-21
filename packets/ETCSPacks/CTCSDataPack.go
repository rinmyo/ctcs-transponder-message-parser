package ETCSPacks

import (
	"TransponderMsgParse/packets"
)

type Etcs44 struct {
	Head struct {
		NID_PACKET uint16 `json:"nid_packet"`
		Q_DIR      uint16 `json:"q_dir"`
		L_PACKET   uint16 `json:"l_packet"`
	}

	NID_XUSER uint16
	XXXXXX    packets.ICtcsPack
}

func (C Etcs44) Encode() ([]byte, error) {
	panic("implement me")
}

func (C *Etcs44) Decode(bytes []byte) error {
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
	C.XXXXXX = packets.GetPacket(string(bytes[23 : 23+9]))
	err := C.XXXXXX.Decode(bytes[23:])
	return err
}

func init() {
	packets.RegisterPacket("00101100", &Etcs44{})
}
