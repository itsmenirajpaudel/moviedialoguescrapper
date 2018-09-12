package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type Quiz struct {
	name     string
	dialogue string
}

func postScrape() {
	doc, err := goquery.NewDocument("http://www.afi.com/100years/quotes.aspx")

	if err != nil {
		log.Fatal(err)
	}

	var dialogues []string
	var movieNames []string

	doc.Find("#subcontent table").Each(func(index int, items *goquery.Selection) {
		if index == 0 {
			items.Find("tr").Each(func(index2 int, tr *goquery.Selection) {
				if index2 != 0 {
					tr.Find("td p span").Each(func(index3 int, span *goquery.Selection) {
						if index3 == 1 {
							dialogues = append(dialogues, span.Text())
							// fmt.Println(span.Text())
						}

						if index3 == 2 {
							movieNames = append(movieNames, span.Text())
							// fmt.Println(span.Text())
						}
					})
				}
			})
		}
	})

	quizes := []Quiz{}
	for index, element := range dialogues {
		q := Quiz{name: (movieNames)[index], dialogue: element}
		quizes = append(quizes, q)
	}

	fmt.Printf("%+v\n", quizes)

}

func main() {
	postScrape()
}
