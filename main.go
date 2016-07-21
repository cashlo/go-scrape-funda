package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/atotto/encoding/csv"
)

type House struct {
	Price   int
	Address string
	Link    string
	Area    int
	Volume  int
	Year    int
}

func ScrapeFunda(url string, result *[]House) {
	fmt.Printf("Scraping %s\n", url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return
	}

	numberRegex, _ := regexp.Compile("[0-9\\.]+")
	notNumberRegex, _ := regexp.Compile("[^0-9]")

	doc.Find(".search-result").Each(func(i int, s *goquery.Selection) {

		var h House
		h.Address = strings.TrimSpace(s.Find(".search-result-title").Text())
		h.Link, _ = s.Find(".search-result-header a").Attr("href")

		firstPrice := numberRegex.FindString(s.Find(".search-result-price").Text())
		priceString := notNumberRegex.ReplaceAllString(firstPrice, "")
		h.Price, _ = strconv.Atoi(priceString)

		GetHouseDetail(&h)

		*result = append(*result, h)
	})
}

func GetHouseDetail(h *House) {
	url := "http://www.funda.nl" + h.Link + "kenmerken/"

	fmt.Printf("Scraping %s\n", url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return
	}
	yearRegex, _ := regexp.Compile("[0-9]{4}")
	numberRegex, _ := regexp.Compile("[0-9]+")

	h.Year, _ = strconv.Atoi(yearRegex.FindString(doc.Find("#boja12").Text()))
	h.Area, _ = strconv.Atoi(numberRegex.FindString(doc.Find("#twwo13").Text()))
	h.Volume, _ = strconv.Atoi(numberRegex.FindString(doc.Find("#twih12").Text()))
}

func main() {
	var results []House
	for i := 1; i <= 5; i++ {
		ScrapeFunda(fmt.Sprintf("http://www.funda.nl/koop/amsterdam/p%d", i), &results)
	}

	f, _ := os.Create("house.txt")
	defer f.Close()

	w := csv.NewWriter(f)
	w.WriteStructAll(results)
	fmt.Println(results)
}
