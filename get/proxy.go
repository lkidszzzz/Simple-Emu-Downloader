package get

import (
	"fmt"
	"github.com/gookit/color"
	"net/http"
	"net/url"
	"strconv"
)

var port string

func proxy() *http.Client {
	if port == "" {
		for {
			fmt.Println()
			color.Notice.Print("	请输入代理端口号(Clash默认为7890)：")
			fmt.Scan(&port)
			_, err := strconv.Atoi(port)
			if err != nil {
				color.Error.Prompt("	Warn:请按要求正确输入数字。输入错误的端口号可能导致问题。")
			} else {
				break
			}
		}
	}
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:" + port)
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}
	return client
}
