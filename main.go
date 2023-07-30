package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector()
	c.AllowedDomains = []string{"golangcode.com"}

	c.OnHTML("a", func(e *colly.HTMLElement) {
		fmt.Println("-", e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Print(r.URL.String(), " is being visited ")

		for i := 1; i <= 3; i++ {
			time.Sleep(1 * time.Second)
			fmt.Print(".")
		}
		fmt.Println("")
	})

	Err := c.Visit("https://golangcode.com/")
	if Err != nil {
		fmt.Println(Err)
	}

}
