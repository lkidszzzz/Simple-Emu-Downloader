package get

import (
	"fmt"
	"github.com/gookit/color"
	"net/http"
	"net/url"
)

var port string

func proxy() *http.Client {
	if port == "" {
		fmt.Println()
		color.Notice.Print("	请输入代理端口号(Clash默认为7890)：")
		fmt.Scan(&port)
	}
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:" + port)
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}
	return client
}
