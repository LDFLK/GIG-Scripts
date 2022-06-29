package decoders

import (
	"GIG-Scripts/extended_models"
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalMOHDecoder struct {
	MyLocalDecoderInterface
}

func (d MyLocalMOHDecoder) DecodeToEntity(record []string, source string, headers []string) models.Entity {

	// 0-id		1-moh_id	2-name	3-centroid	4-population
	entity := *new(extended_models.Location).
		SetLocationId(record[1], source).
		SetName(record[2]+" MOH", source).
		SetCentroid(record[3], source).
		SetPopulation(record[4], source).
		SetGeoCoordinates("gig-data-master/geo/moh/"+record[0]+".json", source).
		AddCategory("MOH").AddCategory("LOCATION")
	return entity
}
