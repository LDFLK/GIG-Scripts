package decoders

import (
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalPDDecoder struct {
	MyLocalDecoderInterface
}

func (d MyLocalPDDecoder) DecodeToEntity(record []string, source string, headers []string) models.Entity {

	// 0-id		1-name	2-country_id	3-province_id	4-district_id	5-ed_id		6-pd_id		7-centroid		8-population
	decoder := MyLocalLocationDecoder{
		LocationId: record[0],
		Name:       record[1] + " Polling Division",
		Centroid:   record[7],
		Population: record[8],
		ParentId:   record[5],
		GeoSource:  "pd",
		Category:   "Polling Division",
		Attribute:  "polling_divisions",
		Source:     source,
	}
	decoder.ParentEntity = decoder.GetParentEntity()
	entity := decoder.MapToEntity()
	entity.AddCategory("LOCATION")
	decoder.AppendToParentEntity(entity)

	return entity
}
