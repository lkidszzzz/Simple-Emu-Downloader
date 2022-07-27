package downloader

import (
	"strings"
)

func Namer(link string) string {
	return link[strings.LastIndex(link, "/")+1:]
}
