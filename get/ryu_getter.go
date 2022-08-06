package get

import (
	"Simple-Emu-Downloader"
	"Simple-Emu-Downloader/downloader"
	"github.com/PuerkitoBio/goquery"
	"github.com/gookit/color"
	"io"
	"log"
	"net/http"
	"strings"
)

//获取Ryu下载链接

func RyuGetter() string {
	if sed.ModeChooser() == 1 {
		color.Light.Prompt("已选择代理模式")
		res, err := sed.Proxy().Get("https://github.com/Ryujinx/release-channel-master/tags")
		if err != nil {
			log.Fatal(err)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(res.Body)
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		sel := doc.Find("h4[data-test-selector=\"tag-title\"]>a").First().Text()
		return "https://github.com/Ryujinx/release-channel-master/releases/download/" +
			strings.Replace(strings.Replace(sel, "\n", "", -1), " ", "", -1) + "/ryujinx-" +
			strings.Replace(strings.Replace(sel, "\n", "", -1), " ", "", -1) + "-win_x64.zip"
	} else {
		color.Light.Prompt("已选择直连模式")
		res, err := http.Get("https://github.com/Ryujinx/release-channel-master/tags")
		if err != nil {
			log.Fatal(err)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(res.Body)
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		sel := doc.Find("h4[data-test-selector=\"tag-title\"]>a").First().Text()
		return "https://github.com/Ryujinx/release-channel-master/releases/download/" +
			strings.Replace(strings.Replace(sel, "\n", "", -1), " ", "", -1) + "/ryujinx-" +
			strings.Replace(strings.Replace(sel, "\n", "", -1), " ", "", -1) + "-win_x64.zip"
	}
}

func RyuCDN1() string {
	return downloader.GithubCdn1(RyuGetter())

}

func RyuCDN2() string {
	return downloader.GithubCdn2(RyuGetter())
}
