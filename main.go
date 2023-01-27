package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(colly.AllowedDomains("www.amazon.in"))

	c.OnHTML("div.s-main-slot.s-result-list.s-search-results.sg-row", func(h *colly.HTMLElement) {
		h.ForEach("div.a-section.a-spacing-small.a-spacing-top-small", func(_ int, h *colly.HTMLElement) {
			var name string
			name = h.ChildText("span.a-size-medium.a-color-base")
			var stars string
			stars = h.ChildText("span.a-icon-alt")
			var price string
			price = h.ChildText("span.a-price-whole")

			fmt.Println("ProductName: ", name)
			fmt.Println("Ratings: ", stars)
			fmt.Println("Price: ", price)
			fmt.Println("")

		})
	})
	c.Visit("https://www.amazon.in/s?k=mobile")
}
