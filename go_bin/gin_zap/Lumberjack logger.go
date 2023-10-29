package main

import (
	"net/http"
)

//func main() {
//	InitLogger()
//	defer sugarLogger.Sync()
//	for i := 0; i <= 100000; i++ {
//
//		simpleHttpGet("www.sogo.com")
//		simpleHttpGet("http://www.sogo.com")
//	}
//
//}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
