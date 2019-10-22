package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type BinMessage struct {
	head string
	body string
}

func NewBinMessage(str string) *BinMessage {
	return &BinMessage{
		str[0:50],
		str[50 : 50+772],
	}
}

func (msg BinMessage) Decode2FrameMark() (result FrameMark) {
	varLengthMap := make(map[string]int) //fm爲長度map
	result = make(FrameMark)

	b, err := ioutil.ReadFile("./packets/FRAMEMARK.json")
	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal([]byte(b), &varLengthMap)
	if err != nil {
		panic(err)
	}

	i := 0 //counter
	for k, v := range varLengthMap {
		result[k] = BIN2DEC(msg.head[i : i+v])
		i = v
	}

	return
}

func Decode2EtcsPacket(binStr string) (result UserInfoPacket) {
	varLengthMap := make(map[string]interface{})
	result = make(UserInfoPacket)

	b, err := ioutil.ReadFile("./packets/ETCS-" + strconv.Itoa(BIN2DEC(binStr[0:8])) + ".json")
	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal([]byte(b), &varLengthMap)
	if err != nil {
		panic(err)
	}

	//將全部視爲一個長度爲一的循環體
	parseUnfixedPart(1, binStr, varLengthMap)

	return
}

/**
length 長度 bin 待解析的二進制字符串 format 需要解析至的格式，string爲變量名， int爲該變量的長度
*/
func parseUnfixedPart(length int, bin string, varLengthMap map[string]interface{}) []map[string]interface{} {
	result := make([]map[string]interface{}, length)

	end := 0 //尾指針

	for i := 0; i < length; i++ {
		result[i] = make(map[string]interface{})
		for k, v := range varLengthMap {
			// value 是當前部分的長度 或 嵌套的部分
			if value, ok := v.(float64); ok {
				// 如果是數字的話
				result[i][k] = BIN2DEC(bin[end : end+int(value)])
				end += int(value)
				//遇到iter跳過下一個，因爲下一個是數組
				if k[0:6] == "N_ITER" {
					arr := k[7:] + "_ARRAY"
					result[i][arr] = parseUnfixedPart(int(value), bin[end:], varLengthMap[arr].(map[string]interface{}))
					continue

				}
			}

		}
	}
	return result
}

//幀標誌
type FrameMark map[string]int

//用戶信息包
type UserInfoPacket map[string]interface{}

//解包到json字符串
func (p UserInfoPacket) Parse2json() (result string) {
	data, err := json.Marshal(map[string]interface{}(p))
	if err != nil {
		panic(err)
	}
	return string(data)
}
