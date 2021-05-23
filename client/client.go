package client

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dstotijn/go-notion"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load .env file - %v", err)
	}
}

func NewClient() *notion.Client {
	LoadEnv()
	apiKey := os.Getenv("APIKEY")
	c := notion.NewClient(apiKey)
	return c
}

func createContent() string {
	day := time.Now()
	content := fmt.Sprintln(day.Format("2006/01/02"))
	return content
}

func CreatePage(c *notion.Client, pageId string) {
	ctx := context.Background()
	content := createContent()
	params := notion.CreatePageParams{
		ParentType: notion.ParentTypePage,
		ParentID:   pageId,
		Title: []notion.RichText{
			{
				Text: &notion.Text{
					Content: content,
				},
			},
		},
		Children: []notion.Block{
			{
				Type: notion.BlockTypeHeading1,
				Heading1: &notion.Heading{
					Text: []notion.RichText{
						{
							Text: &notion.Text{
								Content: "議題",
							},
						},
					},
				},
			},
			{
				Type: notion.BlockTypeHeading1,
				Heading1: &notion.Heading{
					Text: []notion.RichText{
						{
							Text: &notion.Text{
								Content: "決定事項",
							},
						},
					},
				},
			},
		},
	}

	_, err := c.CreatePage(ctx, params)
	if err != nil {
		log.Printf("Could not create page - %v", err)
	} else {
		fmt.Println("Success create minutes")
	}
}
