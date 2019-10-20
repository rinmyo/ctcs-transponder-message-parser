package packets

//var nid map[string]uint16 = map[string]uint16{
//	"ETCS_5":   0B00000101,
//	"ETCS_21":  0B00010101,
//	"ETCS_27":  0B00011011,
//	"ETCS_41":  0B00101001,
//	"ETCS_44":  0B00101100,
//	"ETCS_68":  0B01000100,
//	"ETCS_132": 0B10000100,
//
//	"CTCS_1": 0B000000001,
//	"CTCS_2": 0B000000010,
//	"CTCS_3": 0B000000011,
//	"CTCS_4": 0B000000100,
//}
//
//func GetNID() map[string]uint16 {
//	return nid
//}

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

type IUserInfoPack interface {
	Encode() ([]byte, error)
	Decode([]byte) error
}
