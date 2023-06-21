package helpers

import (
	GIG_Scripts "GIG-Scripts"
	"GIG-Scripts/orgchart/constants"
	"log"
	"time"

	"github.com/lsflk/gig-sdk/models"
)

func TerminateEntities(fileName string, entityName string) {
	/**
	get all orgchart entities
	for each entity add lifeStatus attribute valued terminated by given date
	append 'terminated by date' to all entities
	save entities
	*/

	terminationDate, _ := time.Parse("gazette-2006-1-2.csv", fileName)
	entity := models.Entity{}
	entity.SetSource(constants.SourceName + fileName).SetSourceDate(terminationDate).AddCategory(constants.OrgChartCategory)

	if entityName != "All" {
		entity.Title = entityName
	}
	resp, err := GIG_Scripts.GigClient.PostRequest(GIG_Scripts.GigClient.ApiUrl+"terminate", entity)
	if err != nil {
		log.Println("entity termination error:", err)
	}
	log.Println("Entity Termination:", resp)
}
