package main

import (
	database "scrap/DataBase"
	scrap "scrap/ScrapAllData"
)

func main() {

	database.DataMigration()
	scrap.Scrap()

	// c := colly.NewCollector(colly.AllowedDomains("www.amazon.in"))

	// c.OnHTML("div.s-result-list.s-search-results.sg-row", func(h *colly.HTMLElement) {
	// 	h.ForEach("div.a-section.a-spacing-base.a-text-center", func(_ int, h *colly.HTMLElement) {
	// 		var name string
	// 		name = h.ChildText("span.a-size-base-plus.a-color-base")
	// 		var stars string
	// 		stars = h.ChildText("span.a-icon-alt")
	// 		var star = strings.Split(stars, " ")
	// 		rating := star[0]
	// 		var price string
	// 		price = h.ChildText("span.a-price-whole")

	// 		fmt.Println("ProductName: ", name)
	// 		fmt.Println("Ratings: ", rating)
	// 		fmt.Println("Price: ", price)
	// 		fmt.Println("")

	// 	})
	// })
	// // c.OnHTML("a.s-pagination-item.s-pagination-next.s-pagination-button.s-pagination-separator", func(h *colly.HTMLElement) {
	// // 	next_page := h.Request.AbsoluteURL(h.Attr("href"))
	// // 	c.Visit(next_page)
	// // })
	// // c.OnRequest(func(r *colly.Request) {
	// // 	fmt.Println(r.URL.String())
	// // })

	// c.Visit("https://www.amazon.in/s?k=jacket")
}
