package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	_ "github.com/joho/godotenv/autoload"
	"github.com/shomali11/slacker"
)

type APIResponse struct {
	Page     int       `json:"page"`
	Status   string    `json:"status"`
	Articles []Article `json:"articles"`
}

type Article struct {
	ID          string  `json:"_id"`
	Link        string  `json:"link"`
	Rank        int     `json:"rank"`
	Media       string  `json:"media"`
	Title       string  `json:"title"`
	Topic       string  `json:"topic"`
	Score       float64 `json:"_score"`
	Author      string  `json:"author"`
	Rights      string  `json:"rights"`
	Authors     string  `json:"authors"`
	Country     string  `json:"country"`
	Excerpt     string  `json:"excerpt"`
	Summary     string  `json:"summary"`
	Language    string  `json:"language"`
	CleanURL    string  `json:"clean_url"`
	IsOpinion   bool    `json:"is_opinion"`
	PubDate     string  `json:"published_date"`
	TwitterAcc  string  `json:"twitter_account"`
	PubDatePrec string  `json:"published_date_precision"`
}

func main() {
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	queryDefinition := &slacker.CommandDefinition{
		Description: "Query!",
		Examples:    []string{"data {query} {page_size} {countries}"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			query := request.Param("query")
			pageSize := request.Param("page_size")
			countries := request.Param("countries")
			if query == "" || pageSize == "" || countries == "" {
				response.Reply("Please provide all the parameters")
				return
			}
			baseurl := "https://api.newscatcherapi.com/v2/search?q="
			url := fmt.Sprintf("%s%s&page_size=%s&countries=%s", baseurl, query, pageSize, countries)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				panic(err)
			}
			req.Header.Set("x-api-key", os.Getenv("API_KEY"))
			client := http.DefaultClient
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			var data APIResponse
			err = json.NewDecoder(resp.Body).Decode(&data)
			if err != nil {
				panic(err)
			}
			articles := ""
			separator := "\n----------------------------------------\n"
			for _, article := range data.Articles {
				r := regexp.MustCompile(`\n{2,}`)
				summary := r.ReplaceAllString(article.Summary, "\n")
				articles += fmt.Sprintf("*Title:* %s\n*Authors:* %s\n*Country:* %s\n*Excerpt:* %s\n*Summary:* %s\n*Link:* %s\n*PubDate:* %s\n*Language:* %s\n\n",
					article.Title,
					article.Authors,
					article.Country,
					article.Excerpt,
					summary,
					article.Link,
					article.PubDate,
					article.Language)
				articles += separator
			}
			response.Reply(articles)
		},
	}
	bot.Command("query <query> <page_size> <countries>", queryDefinition)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
