package services

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

func AnalysisHtml(url string) {
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	fmt.Println(string(content))
	if err != nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	demo := doc.Find("h2.title").Text()
	fmt.Println(demo)
}
