package main

type Msg struct {
	*Head
}

type Head struct {
	Q_UPDOWN uint16 `json:"q_updown"`
	M_VERSION uint16 `json:"m_version"`
	Q_MEDIA uint16 `json:"q_media"`
	N_PIG uint16 `json:"n_pig"`
	N_TOTAL uint16 `json:"n_total"`
	M_DUP uint16 `json:"m_dup"`
	M_MCOUNT uint16 `json:"m_mcount"`
	NID_C uint16 `json:"nid_c"`
	NID_BG uint16 `json:"nid_bg"`
	Q_LINK uint16 `json:"q_link"`
}






func NewHead(headStr []byte) *Head {
	return &Head{
		Q_UPDOWN:  Bytes2Uint(headStr[0:1]),
		M_VERSION: Bytes2Uint(headStr[1:8]),
		Q_MEDIA:   Bytes2Uint(headStr[8:9]),
		N_PIG:     Bytes2Uint(headStr[9:12]),
		N_TOTAL:   Bytes2Uint(headStr[12:15]),
		M_DUP:     Bytes2Uint(headStr[15:17]),
		M_MCOUNT:  Bytes2Uint(headStr[17:25]),
		NID_C:     Bytes2Uint(headStr[25:35]),
		NID_BG:    Bytes2Uint(headStr[35:49]),
		Q_LINK:    Bytes2Uint(headStr[49:50]),
	}
}
