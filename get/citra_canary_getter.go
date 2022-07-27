package get

import (
	"Simple-Emu-Downloader/cdn"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strings"
)

func CitraCanaryVerGetter() string {
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

func CitraCanaryGetter(ver string) string {
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

func CitraCanaryCDN1() string {
	return cdn.GithubCdn1(CitraCanaryGetter(CitraCanaryVerGetter()))
}

func CitraCanaryCDN2() string {
	return cdn.GithubCdn2(CitraCanaryGetter(CitraCanaryVerGetter()))
}
