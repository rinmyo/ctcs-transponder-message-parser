package main

import (
	"TransponderMsgParse/packets"
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
		//begin of a packet is from 0
		begin := 0

		//finish with 11111111
		if packets.BINSlice2Uint(bd[begin:begin+8]) == 0b11111111 {
			return
		}

		//end of a packet is by the begin add a packet length
		end := uint16(begin) + packets.BINSlice2Uint(bd[begin+10:begin+23])
		parseBinUserPacket(bd[begin:end]) // begin 和 end 確定一個包
		begin += int(end)                 //解析pk
	}

}

// Judge the type of packet , then generate corresponding packet object, then decode the packet
//only used for single packet
func parseBinUserPacket(pkBinSlice []byte) (rpk packets.IEtcsPack) {
	rpk = packets.GetPacket(string(pkBinSlice[0:8]))
	_ = rpk.Decode(pkBinSlice)
	return
}
