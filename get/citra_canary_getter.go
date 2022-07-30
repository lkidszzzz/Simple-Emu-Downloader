package get

import (
	"Simple-Emu-Downloader/downloader"
	"github.com/PuerkitoBio/goquery"
	"github.com/gookit/color"
	"io"
	"log"
	"net/http"
	"strings"
)

var key int

func CitraCanaryVerGetter() string {
	if modeChooser() == 1 {
		key = 1
		color.Notice.Prompt("已选择代理模式")
		res, err := proxy().Get("https://github.com/citra-emu/citra-canary/tags")
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
		ver := strings.Replace(strings.Replace(doc.Find("h4[data-test-selector=\"tag-title\"]>a").First().Text(), "\n", "", -1), " ", "", -1)
		return ver
	} else {
		color.Notice.Prompt("已选择直连模式")
		res, err := http.Get("https://github.com/citra-emu/citra-canary/tags")
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
		ver := strings.Replace(strings.Replace(doc.Find("h4[data-test-selector=\"tag-title\"]>a").First().Text(), "\n", "", -1), " ", "", -1)
		return ver
	}
}

func CitraCanaryGetter(ver string) string {
	if key == 1 {
		res, err := http.Get("https://github.com/citra-emu/citra-canary/releases")
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
		source := doc.Find("div[class=\"d-flex flex-justify-start col-12 col-lg-9\"]>a:contains(citra-windows-mingw)").First().Text()
		return "https://github.com/citra-emu/citra-canary/releases/download/" + ver + "/" + strings.Replace(strings.Replace(source, "\n", "", -1), " ", "", -1)
	} else {
		res, err := http.Get("https://github.com/citra-emu/citra-canary/releases")
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
		source := doc.Find("div[class=\"d-flex flex-justify-start col-12 col-lg-9\"]>a:contains(citra-windows-mingw)").First().Text()
		return "https://github.com/citra-emu/citra-canary/releases/download/" + ver + "/" + strings.Replace(strings.Replace(source, "\n", "", -1), " ", "", -1)
	}
}

func CitraCanaryCDN1() string {
	return downloader.GithubCdn1(CitraCanaryGetter(CitraCanaryVerGetter()))
}

func CitraCanaryCDN2() string {
	return downloader.GithubCdn2(CitraCanaryGetter(CitraCanaryVerGetter()))
}
