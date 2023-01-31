package scrapalldata

import (
	"fmt"
	Data "scrap/Model/Data_Model"

	database "scrap/DataBase"
	cate "scrap/Model/Category_Model"
	"strings"

	"github.com/gocolly/colly"
)

func Scrap() {

	var name, price, rating, search string

	c := colly.NewCollector(colly.AllowedDomains("www.amazon.in"))

	c.OnHTML("div.nav-search-field", func(r *colly.HTMLElement) {
		searchs := r.Request.URL.Query()
		search = searchs.Get("k")
		var category = []cate.Category{{Category_Name: search}}
		fmt.Println(category)
		database.Database.Create(&category)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	c.OnHTML("div.s-result-list.s-search-results.sg-row", func(h *colly.HTMLElement) {
		h.ForEach("div.sg-col-inner", func(_ int, h *colly.HTMLElement) {
			name = h.ChildText("h2.a-size-mini.a-spacing-none.a-color-base.s-line-clamp-2")
			stars := h.ChildText("span.a-icon-alt")
			star := strings.Split(stars, " ")
			rating = star[0]
			price = h.ChildText("span.a-price-whole")
			fmt.Println("ProductName: ", name)
			fmt.Println("Ratings: ", rating)
			fmt.Println("Price: ", price)
			var cate cate.Category
			database.Database.Where("category_name = ?", search).Find(&cate)
			fmt.Println(search)
			if search == cate.Category_Name {
				var data = []Data.Data{{Name: name, Rating: rating, Price: price, Category_ID: cate.ID}}
				fmt.Println(data)
				database.Database.Create(&data)
			}
			fmt.Println("")

		})
	})
	c.OnHTML("a.s-pagination-item.s-pagination-next.s-pagination-button.s-pagination-separator", func(h *colly.HTMLElement) {
		next_page := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(next_page)
	})

	c.Visit("https://www.amazon.in/s?k=jacket")
}
