package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	//http://127.0.0.1:8080/add?a=1&b=3
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			println("请求参数有错！")
			return
		}
		fmt.Println("path:", r.URL.Path)
		fmt.Println(r.Form["a"][0])
		fmt.Println(r.Form["b"][0])
		a, _ := strconv.Atoi(r.Form["a"][0])
		b, _ := strconv.Atoi(r.Form["b"][0])
		w.Header().Set("Content-Type", "application/json")
		returnData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		_, _ = w.Write(returnData)
	})
	_ = http.ListenAndServe(":9090", nil)
}
