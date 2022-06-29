package helpers

import (
	"GIG-Scripts/extended_models"
	"GIG-Scripts/orgchart/constants"
	"encoding/json"
	"log"
	"time"
)

var categories = []string{constants.OrgChartCategory, constants.Ministry, constants.Organization}

func GenerateOrganizationFromDataRecord(fileName string, ministry string, departments []string, gazetteDate time.Time,
	nameStructure map[string]string) extended_models.Organization {
	var filteredDepartments []string
	for _, department := range departments {
		log.Println("	", department)
		if department != "" {
			filteredDepartments = append(filteredDepartments, department)
		}
	}

	//decode to entity
	jsonDepartments, err := json.Marshal(filteredDepartments)
	if err != nil {
		panic("Error converting to json,")
	}

	ministryName := ministry
	if oldName, ok := nameStructure[ministry]; ok {
		ministryName = oldName
	}

	organization := extended_models.Organization{}
	organization.SetOrganizationTitle(ministryName, constants.SourceName+fileName, gazetteDate).
		SetChildOrganizations(string(jsonDepartments), constants.SourceName+fileName, gazetteDate).
		SetSource(constants.SourceName + fileName).
		SetSourceDate(gazetteDate).
		SetSourceSignature("trusted").
		AddCategories(categories)

	// detect entity name changes and include it in attributes
	if _, newTitleFound := nameStructure[ministry]; newTitleFound {
		organization.SetNewTitle(ministry, constants.SourceName+fileName, gazetteDate)
	}
	return organization
}
