package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Running ÖSK..")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/shop", ShopIndex)

	log.Fatal(http.ListenAndServe(":8080", router))
}

type ShopIndexResponse struct {
	Image string `json:"image"`
	Url   string `json:"url"`
	Title string `json:"title"`
	Price string `json:"price"`
}

func ShopIndex(w http.ResponseWriter, r *http.Request) {
	res := []ShopIndexResponse{}

	gq, err := goquery.NewDocument("http://oskshop.se")
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	big := gq.Find(".frontpage_big1")
	big_link, _ := big.Find("a").First().Attr("href")
	big_image, _ := big.Find("img").First().Attr("src")
	big_text := big.Find("h3").Text()
	big_price := big.Find(".frontbox_big1-price").Text()
	//big_teaser := big.Find(".frontbox_big1-teaser").Text()

	bigProduct := ShopIndexResponse{
		Image: big_image,
		Url:   "http://oskshop.se/" + big_link,
		Title: big_text,
		Price: big_price,
	}

	res = append(res, bigProduct)

	// Find products in list
	gq.Find(".product_list_box").Each(func(i int, s *goquery.Selection) {
		a, _ := s.Find("a").Attr("href")
		img, _ := s.Find("img").Attr("src")
		title := s.Find("h3").Text()
		price := s.Find(".product_list_price").Text()

		row := ShopIndexResponse{
			Image: img,
			Url:   "http://oskshop.se/" + a,
			Title: title,
			Price: price,
		}

		res = append(res, row)
	})

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	json.NewEncoder(w).Encode(res)
}
