package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"ihtPrivateSDK/share/logging"
	pb "protobuf/projects/go/protocol/basic"
)

// RequestMetadata ...
type RequestMetadata struct {
	Method   string      `json:"method"`
	ObjType  int         `json:"objType"`
	MetaData interface{} `json:"meta"` // 在这里 我对它漠不关心
}

// Response ...
type Response struct {
	Code          int32             `json:"code"`
	Message       string            `json:"message"`
	Payload       interface{}       `json:"payload"`
	FabricPrivate *pb.FabricPrivate `json:"private"`
}

//GetCurrentTime 获取当前时间20170101
func GetCurrentTime() int32 {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	date, _ := strconv.Atoi(tm.Format("20060102"))
	return int32(date)
}

//GetCurrentTimeHM 获取当前时间201701010100
func GetCurrentTimeHM() int {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	ss := tm.Format("200601021504")

	dd, err := strconv.Atoi(ss)
	if err != nil {
		logging.Error("%v", err.Error())
		return 0
	}
	return dd
}

//DateAdd 查询某一日期所在周的周日
func DateAdd(date int32) (time.Time, error) {
	var sat time.Time
	swap := date % 10000
	year := int(date / 10000)
	month := swap / 100
	day := int(swap % 100)

	baseTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	weekday := baseTime.Weekday().String()

	var basedate string
	if strings.EqualFold(weekday, "Monday") {
		basedate = fmt.Sprintf("%d%s", 24*6, "h")

	} else if strings.EqualFold(weekday, "Tuesday") {
		basedate = fmt.Sprintf("%d%s", 24*5, "h")

	} else if strings.EqualFold(weekday, "Wednesday") {
		basedate = fmt.Sprintf("%d%s", 24*4, "h")

	} else if strings.EqualFold(weekday, "Thursday") {
		basedate = fmt.Sprintf("%d%s", 24*3, "h")

	} else if strings.EqualFold(weekday, "Friday") {
		basedate = fmt.Sprintf("%d%s", 24*2, "h")

	} else if strings.EqualFold(weekday, "Saturday") {
		basedate = fmt.Sprintf("%d%s", 24*1, "h")

	} else {
		return sat, errors.New("周日有交易..")
	}

	dd, _ := time.ParseDuration(basedate)
	sat = baseTime.Add(dd) //Saturday（星期日）
	return sat, nil
}
