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

func (msg BinMessage) Decode2FrameMark() (frameMark YML) {
	varLengthMap := make(YML, 0) //fm爲長度map
	frameMark = make(YML, 0)

	b, err := ioutil.ReadFile("./packets/FRAMEMARK.yml")
	if err != nil {
		fmt.Print(err)
	}

	err = yaml.Unmarshal([]byte(b), &varLengthMap)
	if err != nil {
		panic(err)
	}

	i := 0 //counter
	for _, v := range varLengthMap {
		frameMark = append(frameMark, yaml.MapItem{
			Key:   v.Key,
			Value: BIN2DEC(msg.head[i : i+v.Value.(int)]),
		})
		i += v.Value.(int)
	}

	return
}

func Decode2EtcsPacket(binStr string) (result []YML) {
	//要有序
	varLengthMap := make(YML, 0)
	result = make([]YML, 0)

	//視爲循環體之長度唯一
	for end := 0; end < 772; {
		if binStr[end:end+8] == "11111111" {
			fmt.Println("解析結束")
			return
		}

		b, err := ioutil.ReadFile("./packets/ETCS-" + strconv.Itoa(BIN2DEC(binStr[end:end+8])) + ".yml")
		if err != nil {
			fmt.Print(err)
		}

		err = yaml.Unmarshal([]byte(b), &varLengthMap)
		if err != nil {
			panic(err)
		}

		fmt.Println("遇到", "ETCS-"+strconv.Itoa(BIN2DEC(binStr[end:end+8])))
		packetLength := BIN2DEC(binStr[end+10 : end+23]) //包長
		fixedPart := parseUnfixedPart(1, binStr[end:end+packetLength], varLengthMap)[0]
		result = append(result, fixedPart)
		end += packetLength
	}

	fmt.Println("報文有誤")
	return
}

/**
* 僅譯用戶信息包之單個也
length爲其長 夫bin者二進位字符串待解析者也 夫varLengthMap者，模板也，
*/
func parseUnfixedPart(length int, bin string, varLengthMap YML) (result []YML) {
	result = make([]YML, length)

	end := 0 //尾指針

	for i := 0; i < length; i++ {
		result[i] = make(YML, 0)
		// 該循環僅給value爲整數複製
		for k, v := range varLengthMap {
			fmt.Println("本輪", k, v)

			//若v.Value爲數字，則value爲其長
			//於任意包，皆應先取其部之定長
			if value, ok := v.Value.(int); ok {
				val := BIN2DEC(bin[end : end+value])

				//ETCS-68
				if v.Key == "D_TRACKINIT" {
					if result[i][k-1].Value == 0 {
						result[i] = append(result[i], yaml.MapItem{
							Key:   -1,
							Value: val,
						})
						continue
					} else {
						result[i] = append(result[i], yaml.MapItem{
							Key:   v.Key,
							Value: val,
						})
						break
					}
				}

				//ETCS-79
				if v.Key == "NID_C" && len(result[i][len(result[i])-1].Key.(string)) >= 12 && result[i][len(result[i])-1].Key.(string)[:12] == "Q_NEWCOUNTRY" && result[i][len(result[i])-1].Value == 0 {
					result[i] = append(result[i], yaml.MapItem{
						Key:   -1,
						Value: val,
					})
					continue
				}

				result[i] = append(result[i], yaml.MapItem{
					Key:   v.Key,
					Value: val,
				})

				fmt.Println("結果", i, k, result[i][len(result[i])-1], "二進制：", bin[end:end+value], ",從", end, "到", end+value, "\n")
				end += value
			}

			//XXXXXX可得是包爲 etcs44， 同得NID_XUSER乃前位者也
			if v.Key == "XXXXXX" {
				fmt.Println("遇到XXXXXX, 啓用", "./packets/CTCS-"+strconv.Itoa(result[i][k-1].Value.(int)))
				b, err := ioutil.ReadFile("./packets/CTCS-" + strconv.Itoa(result[i][k-1].Value.(int)) + ".yml") //啓yml檔之相對者
				if err != nil {
					fmt.Print(err)
				}
				//譯至varLengthMap
				ctcsVarLengthMap := make(YML, 0) //fm爲長度map

				err = yaml.Unmarshal([]byte(b), &ctcsVarLengthMap)
				if err != nil {
					panic(err)
				}

				packetLength := BIN2DEC(bin[10:23])
				fmt.Println("包長：", bin[10:23], "DEC:", packetLength)
				//每ETCS-44僅單CTCS之得嵌
				result[i] = append(result[i], yaml.MapItem{
					Key:   v.Key,
					Value: parseUnfixedPart(1, bin[23:packetLength], ctcsVarLengthMap)[0],
				})
			}

			//若所遇爲數組，if之首者判斷定通過不得
			if v.Key.(string)[2:] == "ARRAY" {
				num := result[i][k-1].Value.(int) //數所嵌者

				if num == 0 {
					fmt.Println("長度爲零，略過數組 \n")
					result[i] = append(result[i], yaml.MapItem{})
					continue
				}

				fmt.Println("=====遇到", v.Key, "開始遞歸=====\n")
				result[i] = append(result[i], yaml.MapItem{
					Key:   v.Key,
					Value: parseUnfixedPart(num, bin[end:], v.Value.(YML)),
				})
				fmt.Println("=====結束遞歸", v.Key, "=====\n")
			}

		}

	}
	return
}

//template
type YML yaml.MapSlice

func (YML) GetElement(interface{}) {

}
