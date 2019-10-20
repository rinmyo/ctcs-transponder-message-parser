package main

import (
	"TransponderMsgParse/packets"
	"TransponderMsgParse/packets/ETCSPacks"
)

type BinMessage struct {
	head []byte
	body []byte
}

func NewBinMessage(str string) *BinMessage {
	return &BinMessage{
		packets.Invert(str[0:50]),
		packets.Invert(str[50 : 50+772]),
	}
}

type Head struct {
	Q_UPDOWN  uint16 `json:"q_updown"`
	M_VERSION uint16 `json:"m_version"`
	Q_MEDIA   uint16 `json:"q_media"`
	N_PIG     uint16 `json:"n_pig"`
	N_TOTAL   uint16 `json:"n_total"`
	M_DUP     uint16 `json:"m_dup"`
	M_MCOUNT  uint16 `json:"m_mcount"`
	NID_C     uint16 `json:"nid_c"`
	NID_BG    uint16 `json:"nid_bg"`
	Q_LINK    uint16 `json:"q_link"`
}

func NewHead(headStr []byte) *Head {
	get := packets.GetPieces(headStr, []uint16{1, 7, 1, 3, 3, 2, 8, 10, 14, 1})
	return &Head{
		get[0],
		get[1],
		get[2],
		get[3],
		get[4],
		get[5],
		get[6],
		get[7],
		get[8],
		get[9],
	}
}

func (bm BinMessage) ParseBody() {
	bd := bm.body
	for {
		p := 0
		if packets.Bytes2Uint(bd[p:p+8]) == 0b11111111 {
			return
		} //結束
		pk := bd[p : uint16(p)+packets.Bytes2Uint(bd[p+10:p+26])] //拆包，自指針起次個包長的部分擷取下來作爲pk
		parseBinUserPacket(pk)                                    //解析pk
		p += int(packets.Bytes2Uint(bd[10:26]))                   //指針移到下一個包的包頭
	}

}

// Judge the type of packet , then generate corresponding packet object, then decode the packet
//only used for single packet
func parseBinUserPacket(pk []byte) (rpk packets.IUserInfoPack) {
	nid := packets.Bytes2Uint(pk[0:8])
	//NID Assert
	switch nid {
	case ETCSPacks.Etcs5Nid:
		rpk = ETCSPacks.NewTransponderLinkPack()
	case ETCSPacks.Etcs21Nid:
	case ETCSPacks.Etcs27Nid:
	case ETCSPacks.Etcs41Nid:
	case ETCSPacks.Etcs68Nid:
	case ETCSPacks.Etcs132Nid:

	}

	_ = rpk.Decode(pk)
	return
}
