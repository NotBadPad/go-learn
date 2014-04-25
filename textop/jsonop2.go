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
}

type StationStatus struct {
	printerStatus  int64 `json:"printerStatus"`  //打印机状态
	cameraStatus   int64 `json:"cameraStatus"`   //摄像头状态
	keyboardStatus int64 `json:"keyboardStatus"` //键盘状态
	sweepgunStatus int64 `json:"sweepgunStatus"` //扫描头状态
}

func test1() {
	msg := &Msg{}
	str := `{"head":{"mid":"c12345678","type":"connect"},"body":{"vmCode":"00001","token":"KYMKYVOKAxLmIvelC7HqVRMKOeyYtikhcS6X7K8CX3M="}}`
	err := json.Unmarshal([]byte(str), msg)
	if err != nil {
		fmt.Println(err)
	}

	bytes, _ := json.Marshal(msg.Body)
	fmt.Println("body:", string(bytes))

	type Parma struct {
		VmCode string `json:"vmCode,omitempty"`
		Token  string `json:"token,omitempty"`
	}
	parma := &Parma{}
	err = json.Unmarshal(bytes, parma)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parma.Token)
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
func main() {
	test3()
}
