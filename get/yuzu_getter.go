package get

import (
	"Simple-Emu-Downloader/downloader"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gookit/color"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

//获取Yuzu下载链接

func YuzuGetter() string {
	if modeChooser() == 1 {
		color.Notice.Prompt("已选择代理模式")
		res, err := proxy().Get("https://github.com/pineappleEA/pineapple-src/tags")
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
		sel := doc.Find("div[data-test-selector='commit-message-expander']>pre").First().Text()
		re := regexp.MustCompile("[0-9]")
		num := fmt.Sprintf(strings.Join(re.FindAllString(sel, -1), ""))
		return "https://github.com/pineappleEA/pineapple-src/releases/download/EA-" + num + "/Windows-Yuzu-EA-" + num + ".7z"
	} else {
		color.Notice.Prompt("已选择直连模式")
		res, err := http.Get("https://github.com/pineappleEA/pineapple-src/tags")
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
		sel := doc.Find("div[data-test-selector='commit-message-expander']>pre").First().Text()
		re := regexp.MustCompile("[0-9]")
		num := fmt.Sprintf(strings.Join(re.FindAllString(sel, -1), ""))
		return "https://github.com/pineappleEA/pineapple-src/releases/download/EA-" + num + "/Windows-Yuzu-EA-" + num + ".7z"
	}
}

func YuzuCDN1() string {
	return downloader.GithubCdn1(YuzuGetter())
}

func YuzuCDN2() string {
	return downloader.GithubCdn2(YuzuGetter())
}
