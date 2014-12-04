package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeFunda(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".nvm").Each(func(i int, s *goquery.Selection) {
		address := strings.TrimSpace(s.Find(".object-street").Text())
		link, _ := s.Find(".object-street").Attr("href")
		price := s.Find(".price").Text()

		fmt.Printf("House %d: %s - %s - %s\n", i, address, link, price)
	})
}

func main() {
	for i := 1; 1 <= 10; i++ {
		ScrapeFunda(fmt.Sprintf("http://www.funda.nl/koop/amsterdam/0-200000/3+kamers/p%d", i))
	}
}
