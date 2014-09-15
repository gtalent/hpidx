package main

import (
	"flag"
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
	pages := flag.Int("pages", 0, "Number of pages of House Pets data")
	flag.Parse()
	var srcs []Comic
	for i := 0; i < *pages + 1; i++ {
		url := fmt.Sprint("http://www.housepetscomic.com/category/comic/page/", i, "/")
		src, err := fetchComicUrl(url)
		print(fmt.Sprintf("\r%d/%d", i, *pages))
		if err != nil {
			fmt.Errorf("%e\n", err)
		} else {
			for _, v := range src {
				srcs = append(srcs, v)
			}
		}
	}
	println()
	out, _ := json.MarshalIndent(srcs, "", "\t")
	fmt.Println(string(out))
}
