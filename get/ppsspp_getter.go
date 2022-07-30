package get

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gookit/color"
	"io"
	"log"
	"net/http"
)

func PpssppGetter() string {
	if modeChooser() == 1 {
		color.Light.Prompt("已选择代理模式")
		res, err := proxy().Get("https://www.ppsspp.org/downloads_all.html")
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
		href, _ := doc.Find("p:contains(zip)").Find("a").First().Attr("href")
		return "https://www.ppsspp.org/" + href
	} else {
		color.Light.Prompt("已选择直连模式")
		res, err := http.Get("https://www.ppsspp.org/downloads_all.html")
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
		href, _ := doc.Find("p:contains(zip)").Find("a").First().Attr("href")
		return "https://www.ppsspp.org/" + href
	}
}
