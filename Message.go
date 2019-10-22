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

func (bm BinMessage) ParseBody() (mmap map[int]packets.IEtcsPack) {
	mmap = map[int]packets.IEtcsPack{}
	fpk := packets.GetPacket(packets.GetStr(bm.body[0:8]))
	for i := 0; ; i++ {
		bm.body = fpk.Decode(bm.body)
		mmap[i] = fpk

		if packets.GetStr(bm.body[0:8]) == "11111111" { //結束則返回
			return
		}
		fpk = packets.GetPacket(packets.GetStr(bm.body[0:8])) //否則繼續解析
	}
}
