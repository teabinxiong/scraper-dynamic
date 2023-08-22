package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type Promotion struct {
	ImageSrc string
	Title    string
	Desc     string
}

func main() {
	// Define the URL to scrape
	scrapeUrl := "https://www.samplewebsite.text/"
	var promotions []Promotion

	// initializing a chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// navigate to the target web page and select the HTML elements of interest
	var nodes []*cdp.Node
	chromedp.Run(ctx,
		chromedp.Navigate(scrapeUrl),
		chromedp.SetValue(`//select`, "54"),
		chromedp.Sleep(3*time.Second),
		chromedp.Nodes(".jet-listing-grid__item .elementor-widget-wrap.elementor-element-populated", &nodes, chromedp.ByQueryAll),
	)

	// scraping data from each node
	var imageSrc, title, desc string
	for _, node := range nodes {
		chromedp.Run(ctx,
			chromedp.AttributeValue("img", "src", &imageSrc, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text("h2.elementor-heading-title>a", &title, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(".jet-listing-dynamic-field__content", &desc, chromedp.ByQuery, chromedp.FromNode(node)),
		)

		promotion := Promotion{}

		promotion.ImageSrc = imageSrc
		promotion.Title = title
		promotion.Desc = desc

		//  if it is not duplicate, add to the array - promotions
		if !isInPromotions(promotion, promotions) {
			promotions = append(promotions, promotion)
		}

	}

	// printing the result - promotion
	for index, promotion := range promotions {
		log.Print("====================================\n")
		log.Printf("Item: %d\n", index+1)
		log.Printf("Image Src: %s\n", promotion.ImageSrc)
		log.Printf("Title: %s\n", promotion.Title)
		log.Printf("Description: %s\n", promotion.Desc)
	}

}

func isInPromotions(item Promotion, promotions []Promotion) bool {
	for _, value := range promotions {
		if value.Title == item.Title {
			return true
		}
	}
	return false
}
