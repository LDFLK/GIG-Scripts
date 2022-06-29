package helpers

import (
	GIG_Scripts "GIG-Scripts"
	"GIG-Scripts/my_local/decoders"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func AddDecodedData(filename string, decoder decoders.MyLocalDecoderInterface, exit chan os.Signal) {
	source := "MyLocal - " + filename
	f, err := os.Open(filename)
	if err != nil {
		log.Panicf("unable to locate file:%s", source)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read tsv values using csv.Reader
	tsvReader := csv.NewReader(f)
	tsvReader.Comma = '\t'

	headers, err := tsvReader.Read()
	if err != nil {
		log.Panicf("headers not found in %s: %s", filename, err)
	}
	log.Println("header found:", headers)

	for {
		select {
		case <-exit:
			log.Println("exiting the decoder. Terminated by user")
			os.Exit(0)
		default:
			rec, err := tsvReader.Read()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatal(err)
			}
			// do something with read line
			fmt.Printf("%+v\n", rec)
			entity := decoder.DecodeToEntity(rec, source, headers)
			//save to db
			_, saveErr := GIG_Scripts.GigClient.CreateEntity(entity)
			if saveErr != nil {
				log.Fatal(saveErr)
			}
			log.Println("saved record:", rec)
		}
	}
}
