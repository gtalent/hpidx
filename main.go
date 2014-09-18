/*
   Copyright 2014 gtalent2@gmail.com

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
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
	for i := 1; i < *pages + 1; i++ {
		url := fmt.Sprint("http://www.housepetscomic.com/category/comic/page/", i, "/")
		src, err := fetchComicUrl(url)
		print(fmt.Sprint("\r", i, "/", *pages))
		if err != nil {
			fmt.Errorf("%e\n", err)
		} else {
			srcs = append(srcs, src...)
		}
	}
	println()
	out, _ := json.MarshalIndent(srcs, "", "\t")
	fmt.Println(string(out))
}
