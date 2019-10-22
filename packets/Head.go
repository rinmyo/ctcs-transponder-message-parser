package packets

type ETCS_Head struct {
	NID_PACKET uint16 `json:"nid_packet"`
	Q_DIR      uint16 `json:"q_dir"`
	L_PACKET   uint16 `json:"l_packet"`
	Q_SCALE    uint16 `json:"q_scale"`
}

type CTCS_Head struct {
	NID_XUSER uint16 `json:"nid_xuser"`
	Q_DIR     uint16 `json:"q_dir"`
	L_PACKET  uint16 `json:"l_packet"`
	Q_SCALE   uint16 `json:"q_scale"`
}

var packetMap = make(map[string]IEtcsPack)

func RegisterPacket(nidStr string, sample IEtcsPack) {
	packetMap[nidStr] = sample
}

func GetPacket(nidStr string) IEtcsPack {
	return packetMap[nidStr]
}

type UserInfoPacket struct {
	Length uint16
}

func (u UserInfoPacket) Encode() ([]byte, error) {
	panic("implement me")
}

func (u UserInfoPacket) Decode(binSlice []byte) []byte {
	panic("implement me")
}

type IEtcsPack interface {
	Encode() ([]byte, error)
	Decode(binSlice []byte) []byte
}

type ICtcsPack interface {
	IEtcsPack
}
