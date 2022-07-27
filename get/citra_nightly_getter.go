package get

import (
	"Simple-Emu-Downloader/cdn"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strings"
)

func CitraNightlyVerGetter() string {
	res, err := http.Get("https://github.com/citra-emu/citra-nightly/tags")
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

func CitraNightlyGetter(ver string) string {
	res, err := http.Get("https://github.com/citra-emu/citra-nightly/releases")
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
	return "https://github.com/citra-emu/citra-nightly/releases/download/" + ver + "/" + strings.Replace(strings.Replace(source, "\n", "", -1), " ", "", -1)
}

func CitraNightlyCDN1() string {
	return cdn.GithubCdn1(CitraNightlyGetter(CitraNightlyVerGetter()))
}

func CitraNightlyCDN2() string {
	return cdn.GithubCdn2(CitraNightlyGetter(CitraNightlyVerGetter()))
}
