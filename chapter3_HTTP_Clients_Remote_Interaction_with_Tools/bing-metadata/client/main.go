package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/tralahm/blackhat-go/chapter3_HTTP_Clients_Remote_Interaction_with_Tools/bing-metadata/metadata"
	"github.com/PuerkitoBio/goquery"
)

func handler(i int, s *goquery.Selection) {
	url, ok := s.Find("a").Attr("href")
	if !ok {
		return
	}

	fmt.Printf("%d: %s\n", i, url)
	res, err := http.Get(url)
	if err != nil {
		return
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()

	r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return
	}

	cp, ap, err := metadata.NewProperties(r)
	if err != nil {
		return
	}

	log.Printf(
		"%21s %s - %s %s\n",
		cp.Creator,
		cp.LastModifiedBy,
		ap.Application,
		ap.GetMajorVersion())
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Missing required argument. Usage: main.go <domain> <ext>")
	}
	domain := os.Args[1]
	filetype := os.Args[2]

	q := fmt.Sprintf(
		"site:%s && filetype:%s && instreamset:(url title):%s",
		domain,
		filetype,
		filetype)
	search := fmt.Sprintf("http://www.bing.com/search?q=%s", url.QueryEscape(q))
	doc, err := goquery.NewDocument(search)
	if err != nil {
		log.Panicln(err)
	}

	s := "html body div#b_content ol#b_results li.b_algo div.b_title h2"
	doc.Find(s).Each(handler)
}
