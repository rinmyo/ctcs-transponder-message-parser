package packets

type ETCS_Head struct {
	NID_PACKET uint16 `json:"nid_packet"`
	Q_DIR      uint16 `json:"q_dir"`
	L_PACKET   uint16 `json:"l_packet"`
	Q_SCALE    uint16 `json:"q_scale"`
}

func NewETCS_Head(bytes []byte) *ETCS_Head {
	h := GetPieces(bytes, []uint16{8, 2, 13, 2})
	return &ETCS_Head{
		h[0],
		h[1],
		h[2],
		h[3],
	}
}

type CTCS_Head struct {
	NID_XUSER uint16 `json:"nid_xuser"`
	Q_DIR     uint16 `json:"q_dir"`
	L_PACKET  uint16 `json:"l_packet"`
	Q_SCALE   uint16 `json:"q_scale"`
}

func NewCTCS_Head(bytes []byte) *ETCS_Head {
	h := GetPieces(bytes, []uint16{9, 2, 13, 2})
	return &ETCS_Head{
		h[0],
		h[1],
		h[2],
		h[3],
	}
}

var packetMap = make(map[string]IEtcsPack)

func RegisterPacket(nidStr string, sample IEtcsPack) {
	packetMap[nidStr] = sample
}

func GetPacket(nidStr string) IEtcsPack {
	return packetMap[nidStr]
}

type IEtcsPack interface {
	Encode() ([]byte, error)
	Decode([]byte) error
}

type ICtcsPack interface {
	IEtcsPack
}
