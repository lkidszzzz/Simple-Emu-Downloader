package sed

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

//获取Yuzu下载链接

func YuzuGetter() string {
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
	link := "https://github.com/pineappleEA/pineapple-src/releases/download/EA-" + num + "/Windows-Yuzu-EA-" + num + ".7z"
	return link
}

func YuzuCDN1() string {
	str := YuzuGetter()
	return "https://ghproxy.com/" + str
}

func YuzuCDN2() string {
	str := YuzuGetter()
	return "https://gh.api.99988866.xyz/" + str
}
