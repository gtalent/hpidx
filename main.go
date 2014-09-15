package main

import (
	"fmt"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
)

type Comic struct {
	Image  string
	AltTxt string
}

func fetchComicUrl(url string) ([]Comic, error) {
	var srcs []Comic
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return srcs, err
	}

	doc.Find(".comicarchiveframe").Each(func(i int, s *goquery.Selection) {
		img := s.Find("img")
		src, _ := img.Attr("src")
		alt, _ := img.Attr("alt")
		srcs = append(srcs, Comic{Image: src, AltTxt: alt})
	})

	return srcs, nil
}

func main() {
	src, err := fetchComicUrl("http://www.housepetscomic.com/2008/06/02/")
	if err != nil {
		fmt.Println(err)
	} else {
		out, _ := json.Marshal(src)
		fmt.Println(string(out))
	}
}
