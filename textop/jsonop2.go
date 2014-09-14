package main

import (
	"encoding/json"
	"fmt"
)

type Msg struct {
	Head MsgHead     `json:"head"`
	Body interface{} `json:"body"`
}

type MsgHead struct {
	Mid  string `json:"mid,omitempty"`
	Type string `json:"type,omitempty"`
	Code int    `json:"code,omitempty"`
	Desc string `json:"msg,omitempty"`
	Cts  int64  `json:"cts"`
}

type StationStatus struct {
	printerStatus  int64 `json:"printerStatus"`  //打印机状态
	cameraStatus   int64 `json:"cameraStatus"`   //摄像头状态
	keyboardStatus int64 `json:"keyboardStatus"` //键盘状态
	sweepgunStatus int64 `json:"sweepgunStatus"` //扫描头状态
}

func test1() {
	msg := &Msg{}
	str := `{"head":{"mid":"c12345678","type":"connect"},"body":{"vmCode":"00001","token":"KYMKYVOKAxLmIvelC7HqVRMKOeyYtikhcS6X7K8CX3M=","cts":1341234223}}`
	err := json.Unmarshal([]byte(str), msg)
	if err != nil {
		fmt.Println(err)
	}

	bytes, _ := json.Marshal(msg.Body)
	fmt.Println("body:", string(bytes))

	type Parma struct {
		VmCode string  `json:"vmCode,omitempty"`
		Token  string  `json:"token,omitempty"`
		Cts    float64 `json:"cts"`
	}
	parma := &Parma{}
	err = json.Unmarshal(bytes, parma)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(int64(parma.Cts))
}

func test2() {
	msg := &Msg{}
	str := `{"head":{"mid":"c12345678","type":"boxStatus"},"body":[{"boxSeq":5,"boxCode":"A5","deviceId":1,"status":1},{"boxSeq":6,"boxCode":"A6","deviceId":1,"status":1}]}`
	err := json.Unmarshal([]byte(str), msg)
	if err != nil {
		fmt.Println(err)
	}

	bytes, _ := json.Marshal(msg.Body)
	fmt.Println("body:", string(bytes))

	type Parma struct {
		BoxSeq   int64  `json:"boxSeq,omitempty"`
		BoxCode  string `json:"boxCode,omitempty"`
		DeviceId int64  `json:"deviceId,omitempty"`
		Status   int    `json:"status,omitempty"`
	}
	parmas := []*Parma{}
	err = json.Unmarshal(bytes, &parmas)
	if err != nil {
		fmt.Println(err)
	}
	for _, parma := range parmas {
		fmt.Println(parma.BoxCode)
	}
}

func test3() {
	stationStatus := &StationStatus{}
	str := `{"printerStatus":2,"cameraStatus":2,"keyboardStatus":2,"sweepgunStatus":2}`
	err := json.Unmarshal([]byte(str), stationStatus)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stationStatus)
}

func test4() {
	msg := &Msg{}
	str := `{"head":{"mid":"c12345678","type":"boxStatus","cts":139943085681542411}}`
	err := json.Unmarshal([]byte(str), msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msg.Head.Cts)
}

func test5() {
	msg := &Msg{}
	err := json.Unmarshal([]byte(nil), msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msg.Head.Cts)
}

func test6() {
	msg := &Msg{}

	str := `{"head":{"mid":"c12345678","type":"boxStatus"},"body":[{"boxSeq":5,"boxCode":"A5","deviceId":1,"status":1},{"boxSeq":6,"boxCode":"A6","deviceId":1,"status":1}]}`

	type Parma struct {
		BoxSeq   int64  `json:"boxSeq,omitempty"`
		BoxCode  string `json:"boxCode,omitempty"`
		DeviceId int64  `json:"deviceId,omitempty"`
		Status   int    `json:"status,omitempty"`
	}

	parmas := make([]*Parma, 0)
	msg.Body = parmas
	err := json.Unmarshal([]byte(str), msg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(msg.Body)
}

func test7() {
	a := map[int]string{3: "aaa", 4: "bbb", 5: "ccc"}
	bytes, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(bytes))
}

func test8() {
	a := make([]interface{}, 2)
	b := map[string]interface{}{"aaa": "aaa", "bbb": 12}
	a[0] = b
	a[1] = b
	bytes, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(bytes))
}

func test9() {
	type Parma struct {
		VmCode string `json:"vmCode"`
		Sn     int64  `json:"sn"`
		CTime  string `json:"ctime"`
	}

	parmas := make([]*Parma, 0)
	parma := &Parma{
		VmCode: "0001407",
		Sn:     1,
		CTime:  "2014-09-10 11:25:30",
	}
	parmas = append(parmas, parma)
	parma = &Parma{
		VmCode: "0001407",
		Sn:     1,
		CTime:  "2014-09-10 11:25:30",
	}
	parmas = append(parmas, parma)
	bytes, err := json.Marshal(parmas)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(bytes))

}

func test10() {
	type ProSellNum struct {
		ProId int64 `json:"proId"`
		Num   int64 `json:num`
	}

	type SellResult struct {
		VmCode  string        `json:"vmCode"`
		Sn      int64         `json:"sn"`
		SellNum []*ProSellNum `json:"sellNum"`
	}

	jsonStr := `[{"vmCode": "0001407", "sn": 2, "sellNum": [{"proId": 54, "num": 100}, {"proId": 12, "num": 100}, {"proId": 33, "num": 100}]},{"vmCode": "0001407", "sn": 2, "sellNum": [{"proId": 54, "num": 100}, {"proId": 12, "num": 100}, {"proId": 33, "num": 100}]},{"vmCode": "0001407", "sn": 2, "sellNum": [{"proId": 54, "num": 100}, {"proId": 12, "num": 100}, {"proId": 33, "num": 100}]}]`

	var srs []*SellResult
	err := json.Unmarshal([]byte(jsonStr), &srs)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf(srs[0].VmCode)
	}

}

/**
 * 出货量接口数据
 */
type ProSellNum struct {
	ProId int64 `json:"proId"`
	Num   int64 `json:"num"`
}

type SellResult struct {
	VmCode  string        `json:"vmCode"`
	Sn      int64         `json:"sn"`
	SellNum []*ProSellNum `json:"sellNum"`
}

func main() {
	test10()
}
