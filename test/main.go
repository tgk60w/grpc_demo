/*
@Time : 2020/3/12 上午11:58
@Author : gongyikun
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	//var b uint8 = 1
	//
	//fmt.Printf("%v\n",^b)
	//fmt.Printf("%d对应的bit:%v\n",b,biu.ToBinaryString(b))
	//fmt.Printf("^%d=%d对应的bit:%v\n",b,^b,biu.ToBinaryString(^b))
	for i := 0; i < 300; i++ {
		go func() {
			client := &http.Client{}

			req, err := http.NewRequest("POST", "http://127.0.0.1:8001", strings.NewReader("name=cjb"+fmt.Sprintf("%v", i)))

			if err != nil {
				// handle error
			}

			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			resp, err := client.Do(req)

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				// handle error
			}

			resp.Body.Close()
			fmt.Println(string(body))

		}()

	}
	time.Sleep(5 * time.Second)
	fmt.Printf("5秒了\n")
	time.Sleep(5 * time.Second)
}
