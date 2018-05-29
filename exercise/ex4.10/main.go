// ex4.10 prints t table of Github issues matching the search terms, organized
// by the past day, month, and year.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/iostiny/gopl/example/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	format := "#%-5d %12.11s %.99s\n"
	now := time.Now()

	pastDay := make([]*github.Issue, 0)
	pastMonth := make([]*github.Issue, 0)
	pastYear := make([]*github.Issue, 0)

	day := now.AddDate(0, 0, -1)
	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)

	for _, item := range result.Items {
		switch {
		case item.CreateAt.After(day):
			pastDay = append(pastDay, item)
		case item.CreateAt.After(month) && item.CreateAt.Before(day):
			pastMonth = append(pastMonth, item)
		case item.CreateAt.After(year) && item.CreateAt.Before(month):
			pastYear = append(pastYear, item)
		}

		if len(pastDay) > 0 {
			fmt.Printf("\nPast day:\n")
			for _, item := range pastDay {
				fmt.Printf(format, item.Number, item.User.Login, item.Title)
			}
		}
		if len(pastMonth) > 0 {
			fmt.Printf("\nPast Month:\n")
			for _, item := range pastMonth {
				fmt.Printf(format, item.Number, item.User.Login, item.Title)
			}
		}
		if len(pastYear) > 0 {
			fmt.Printf("\nPast Year:\n")
			for _, item := range pastYear {
				fmt.Printf(format, item.Number, item.User.Login, item.Title)
			}
		}
	}
}
