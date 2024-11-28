package main

import (
	GIG_Scripts "GIG-Scripts"
	"GIG-Scripts/orgchart/helpers"
	"flag"
	"log"
	"os"

	"github.com/lsflk/gig-sdk/libraries"
	"github.com/lsflk/gig-sdk/models"
)

func main() {

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		log.Println("file path not specified")
		os.Exit(1)
	}

	dataArray, fileName, gazetteDate, err := helpers.LoadDataFromCsv(args[0])

	if err != nil {
		log.Fatal("error loading csv file:", err)
	}

	dataStructure, nameStructure := helpers.GenerateDataStructures(dataArray, fileName)

	for ministry, departments := range dataStructure {
		log.Println(ministry)

		organization := helpers.GenerateOrganizationFromDataRecord(fileName, ministry, departments, gazetteDate, nameStructure)

		var entities []models.Entity
		for _, departmentName := range departments {
			childEntity := helpers.CreateChildOrganization(fileName, departmentName, gazetteDate, ministry)
			entities = append(entities, childEntity.Entity)
		}

		err = GIG_Scripts.GigClient.AddEntitiesAsLinks(&organization.Entity, entities)
		if err != nil {
			panic(err)
		}

		log.Println("P1 Done")

		//save to db
		_, saveErr := GIG_Scripts.GigClient.CreateEntity(organization.Entity)
		log.Println("P2 Done")
		libraries.ReportError(saveErr, ministry)
	}

}
