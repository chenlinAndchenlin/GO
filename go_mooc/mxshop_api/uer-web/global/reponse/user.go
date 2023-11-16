package reponse

import (
	"fmt"
	"time"
)

type JsonTiem time.Time

func (j JsonTiem) MarshalJSON() ([]byte, error) {
	var stmp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stmp), nil
}

type UserResponse struct {
	Id       int32    `json:"id"`
	NickName string   `json:"name"`
	BIrthDay JsonTiem `json:"birthday"`
	//BIrthDay string `json:"birthday"`
	Mobile string `json:"mobile"`
	Gender string `json:"gender"`
}
