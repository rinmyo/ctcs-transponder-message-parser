package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
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

	b, err := ioutil.ReadFile("./packets/FRAMEMARK.yml")
	if err != nil {
		fmt.Print(err)
	}

	err = yaml.Unmarshal([]byte(b), &varLengthMap)
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

//解析單個包
func Decode2EtcsPacket(binStr string) (result []UserInfoPacket) {
	varLengthMap := make(map[interface{}]interface{})
	result = make([]UserInfoPacket, 0)
	fmt.Println(binStr[0:8])
	b, err := ioutil.ReadFile("./packets/ETCS-" + strconv.Itoa(BIN2DEC(binStr[0:8])) + ".yml")
	if err != nil {
		fmt.Print(err)
	}

	err = yaml.Unmarshal([]byte(b), &varLengthMap)
	if err != nil {
		panic(err)
	}

	//將全部視爲一個長度爲一的循環體
	for end := 0; end < 772; {
		if binStr[end:end+8] == "11111111" {
			fmt.Println("解析結束")
		}
		packetLength := BIN2DEC(binStr[end+10 : end+23])
		fixedPart := parseUnfixedPart(1, binStr[end:end+packetLength], varLengthMap)[0]
		result = append(result, fixedPart)
		end += packetLength
	}

	fmt.Println("報文有誤")
	return
}

/**
* 僅僅解析單個用戶信息包
length 長度 bin 待解析的二進制字符串 format 需要解析至的格式，string爲變量名， int爲該變量的長度
*/
func parseUnfixedPart(length int, bin string, varLengthMap map[interface{}]interface{}) []map[string]interface{} {
	result := make([]map[string]interface{}, length)

	end := 0 //尾指針

	for i := 0; i < length; i++ {
		fmt.Println(i)
		result[i] = make(map[string]interface{})
		// 該循環僅給value爲整數複製
		for k, v := range varLengthMap {
			// value 是當前部分的長度
			if value, ok := v.(int); ok {

				fmt.Println(end, end+value)
				// 如果是數字的話
				result[i][k.(string)] = BIN2DEC(bin[end : end+value])
				end += value

				if len(k.(string)) == 8 && k.(string)[0:6] == "N_ITER" {
					arr := k.(string)[7:] + "_ARRAY"
					result[i][arr] = parseUnfixedPart(value, bin[end:], varLengthMap[arr].(map[interface{}]interface{}))
				}
			}
		}

		//若是etcs44單獨處理
		if result[i]["NID_PACKET"] == 44 {

			b, err := ioutil.ReadFile("./packets/CTCS-" + strconv.Itoa(result[i]["NID_XUSER"].(int)) + ".yml") //打開對應的yml檔案
			if err != nil {
				fmt.Print(err)
			}
			//解析到varLengthMap中
			err = yaml.Unmarshal([]byte(b), &varLengthMap)
			if err != nil {
				panic(err)
			}

			result[i]["XXXXXX"] = parseUnfixedPart(1, bin[23:result[i]["L_PACKET"].(int)], varLengthMap)[0]

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
	data, err := yaml.Marshal(map[string]interface{}(p))
	if err != nil {
		panic(err)
	}
	return string(data)
}
