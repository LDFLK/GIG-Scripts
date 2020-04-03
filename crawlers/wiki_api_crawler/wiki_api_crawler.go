// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG-SDK/models"
	"GIG-Scripts/crawlers/wiki_api_crawler/decoders"
	"GIG-Scripts/crawlers/wiki_api_crawler/requests"
	"GIG-Scripts/entity_handlers"
	"flag"
	"fmt"
	"os"
	"sync"
)

var visited = make(map[string]bool)

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	if len(args) < 1 {
		fmt.Println("starting title not specified")
		os.Exit(1)
	}
	queue := make(chan string)
	go func() { queue <- args[0] }()

	for title := range queue {
		if title != "" {
			entity := enqueue(title, queue)
			if !entity.IsNil() {
				_, err := entity_handlers.CreateEntity(entity)
				if err != nil {
					fmt.Println(err.Error(), title)
				}
			}
		}
	}
	fmt.Println("end")
}

func enqueue(title string, queue chan string) models.Entity {
	fmt.Println("fetching", title)
	visited[title] = true
	entity := models.Entity{}

	var requestWorkGroup sync.WaitGroup
	for _, propType := range requests.PropTypes() {

		requestWorkGroup.Add(1)
		go func(prop string) {
			defer requestWorkGroup.Done()
			result, err := requests.GetContent(prop, title)
			if err != nil {
				fmt.Println(err)
			} else {
				decoders.Decode(result, &entity)
			}
		}(propType)
	}
	requestWorkGroup.Wait()

	if !entity.IsNil() {

		var (
			linkEntities []models.Entity
			err          error
		)

		for _, link := range entity.Links {
			if link != "" {
				if !visited[link] {
					//fmt.Println("	passed link ->", link.Title)
					go func(title string) {
						queue <- title
						//fmt.Println("	queued link ->", link.Title)
					}(link)
				}
				//add link as an entity
				linkEntities = append(linkEntities, models.Entity{Title:link})
			}
		}

		entity, err = entity_handlers.AddEntitiesAsLinks(entity, linkEntities)

		if err != nil {
			fmt.Println("error creating links:", err)
		}
	}
	return entity
}
