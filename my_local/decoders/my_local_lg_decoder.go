package decoders

import (
	"GIG-Scripts/extended_models"
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalLGDecoder struct {
	MyLocalDecoderInterface
}

func (d MyLocalLGDecoder) DecodeToEntity(record []string, source string, headers []string) models.Entity {

	// 0-id		1-lg_id		2-name	3-centroid	4-population
	entity := *new(extended_models.Location).
		SetLocationId(record[1], source).
		SetName(record[2]+" Local Government", source).
		SetCentroid(record[3], source).
		SetPopulation(record[4], source).
		SetGeoCoordinates("gig-data-master/geo/lg/"+record[0]+".json", source).
		AddCategory("Local Government").AddCategory("LOCATION")
	return entity
}
