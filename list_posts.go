package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
)

type Quiz struct {
	Name     string
	Dialogue string
}

func postScrape(w http.ResponseWriter, r *http.Request) {
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
		q := Quiz{Name: (movieNames)[index], Dialogue: element}
		quizes = append(quizes, q)
	}
	json.NewEncoder(w).Encode(quizes)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", postScrape).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

	// postScrape()
}
