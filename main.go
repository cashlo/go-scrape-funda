package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type House struct {
	address string
	price   string
	link    string
}

func ScrapeFunda(url string, result *[]House) {
	fmt.Printf("Scraping %s\n", url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".nvm").Each(func(i int, s *goquery.Selection) {
		var h House
		h.address = strings.TrimSpace(s.Find(".object-street").Text())
		h.link, _ = s.Find(".object-street").Attr("href")
		h.price = s.Find(".price").Text()

		*result = append(*result, h)
	})
}

func main() {
	var results []House
	for i := 1; i <= 10; i++ {
		ScrapeFunda(fmt.Sprintf("http://www.funda.nl/koop/amsterdam/0-200000/3+kamers/p%d", i), &results)
	}
	fmt.Println(results)
}
