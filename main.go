package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type House struct {
	address string
	price   string
	link    string
	area    int
	volume  int
	year    int
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

		GetHouseDetail(&h)

		*result = append(*result, h)
	})
}

func GetHouseDetail(h *House) {
	url := "http://www.funda.nl" + h.link + "kenmerken/"

	fmt.Printf("Scraping %s\n", url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	yearRegex, _ := regexp.Compile("[0-9]{4}")
	numberRegex, _ := regexp.Compile("[0-9]+")
	h.year, _ = strconv.Atoi(yearRegex.FindString(doc.Find("#boja12").Text()))
	h.area, _ = strconv.Atoi(numberRegex.FindString(doc.Find("#twwo13").Text()))
	h.volume, _ = strconv.Atoi(numberRegex.FindString(doc.Find("#twih12").Text()))
}

func main() {
	var results []House
	for i := 1; i <= 1; i++ {
		ScrapeFunda(fmt.Sprintf("http://www.funda.nl/koop/amsterdam/0-200000/3+kamers/p%d", i), &results)
	}
	fmt.Println(results)
}
