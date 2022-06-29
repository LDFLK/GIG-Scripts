package decoders

import (
	"GIG-Scripts/extended_models"
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalCountryDecoder struct {
	MyLocalDecoderInterface
}

func (d MyLocalCountryDecoder) DecodeToEntity(record []string, source string, headers []string) models.Entity {
	// 0-id	1-country_id	2-name	3-population
	entity := *new(extended_models.Location).
		SetLocationId(record[1], source).
		SetName(record[2], source).
		SetPopulation(record[3], source).
		AddCategory("Country").AddCategory("LOCATION")
	return entity
}
