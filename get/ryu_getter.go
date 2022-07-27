package get

import (
	"Simple-Emu-Downloader/cdn"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strings"
)

//获取Ryu下载链接

func RyuGetter() string {
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

func RyuCDN1() string {
	return cdn.GithubCdn1(RyuGetter())

}

func RyuCDN2() string {
	return cdn.GithubCdn2(RyuGetter())
}
