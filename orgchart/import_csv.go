package main

import (
	GIG_Scripts "GIG-Scripts"
	"GIG-Scripts/orgchart/helpers"
	"log"

	"github.com/lsflk/gig-sdk/libraries"
	"github.com/lsflk/gig-sdk/models"
)

func import_csv(csvFilePath string) {

	dataArray, fileName, gazetteDate, err := helpers.LoadDataFromCsv(csvFilePath)

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

		//save to db
		_, saveErr := GIG_Scripts.GigClient.CreateEntity(organization.Entity)
		libraries.ReportError(saveErr, ministry)
	}
}
