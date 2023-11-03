package main

import (
	"encoding/json"
	"fmt"

	"github.com/kirinlabs/HttpRequest"
)

type ResData struct {
	Data int `json:"data"`
}

func Add(a, b int) int {
	req := HttpRequest.NewRequest()
	r, _ := req.Get(fmt.Sprintf("http://127.0.0.1:9090/%s?a=%d&b=%d", "add", a, b))
	res, _ := r.Body()
	//fmt.Println(string(res))
	//反序列化
	rspData := new(ResData)
	_ = json.Unmarshal(res, rspData)
	return rspData.Data
}
func main() {
	fmt.Println(Add(1, 9))
	fmt.Println(Add(2, 5))
}
