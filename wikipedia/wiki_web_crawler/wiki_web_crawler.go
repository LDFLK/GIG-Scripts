// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG-SDK/libraries"
	"GIG-SDK/libraries/clean_html"
	"GIG-SDK/models"
	"GIG-SDK/request_handlers"
	"GIG-Scripts/wikipedia/wiki_web_crawler/parsers"
	"errors"
	"flag"
	"golang.org/x/net/html"
	"io"
	"log"
	"os"
)

var visited = make(map[string]bool)

const bufferFile = "wiki.log"

func main() {
	flag.Parse()
	args := flag.Args()
	log.Println(args)
	if len(args) < 1 {
		log.Println("starting url not specified")
		os.Exit(1)
	}
	queue := make(chan string)
	go func() { queue <- args[0] }()

	var f *os.File

	if _, err := os.Stat(bufferFile); errors.Is(err, os.ErrNotExist) {
		log.Println("create buffer file")
		f, err = os.Create(bufferFile)

		if err != nil {
			log.Fatal(err)
		}
		//defer f.Close()
	} else {
		log.Println("open buffer file")
		f, err = os.OpenFile(bufferFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		//defer f.Close()
	}

	// open buffer if file exist
	// else create new file

	for uri := range queue {
		entity, err := enqueue(uri, queue, f)
		if err != nil {
			log.Println("enqueue error:", err.Error(), uri)
			continue
		}
		log.Println(entity.ImageURL)
		_, err = request_handlers.CreateEntity(entity)
		log.Println("entity added", entity.Title)
		if err != nil {
			log.Println(err.Error(), uri)
		}
		//TODO: remove added entity from buffer
	}

	defer f.Close()
}

func enqueue(uri string, queue chan string, f *os.File) (models.Entity, error) {
	log.Println("fetching", uri)
	visited[uri] = true

	var (
		entity models.Entity
		err    error
		body   *html.Node
	)

	entity = models.Entity{}.SetSource(uri).SetSourceSignature("trusted")

	resp, err := request_handlers.GetRequest(uri)
	if err != nil {
		return entity, err
	}

	doc, err := libraries.HTMLStringToDoc(resp)
	if err != nil {
		return entity, err
	}

	entity.Title, body, err = parsers.ParseHTMLContent(doc)
	if err != nil {
		return entity, err
	}

	//clean html code by removing unwanted information
	htmlCleaner := clean_html.HtmlCleaner{Config: clean_html.Config{
		LineBreakers:   []string{"div", "caption"},
		IgnoreElements: []string{"noscript", "script", "style", "input"},
		IgnoreStrings:  []string{"[", "]", "edit", "Jump to search", "Jump to navigation"},
		IgnoreTitles:   []string{"(page does not exist)", ":"},
		IgnoreClasses:  []string{"box-Multiple_issues"},
	}}
	result, linkedEntities, imageList, defaultImageSource := htmlCleaner.CleanHTML(uri, body)
	entity.ImageURL = defaultImageSource

	// queue new links for crawling
	for _, linkedEntity := range linkedEntities {
		if !visited[linkedEntity.GetSource()] {
			go func(url string, file *os.File) {
				queue <- url
				if _, err := io.WriteString(file, url+"\n"); err != nil {
					log.Println("error writing to buffer: ", err)
				}
			}(linkedEntity.GetSource(), f)
		}
		entity = entity.AddLink(models.Link{}.SetTitle(linkedEntity.GetTitle()).AddDate(entity.GetSourceDate()))
	}

	for _, image := range imageList {
		go func(payload models.Upload) {
			request_handlers.UploadImage(payload)
		}(image)
	}

	// save linkedEntities (create empty if not exist)
	entity, err = request_handlers.AddEntitiesAsLinks(entity, linkedEntities)
	entity = entity.SetAttribute("content", models.Value{
		ValueType:   "html",
		ValueString: result,
	}).AddCategory("Wikipedia")
	return entity, nil
}
