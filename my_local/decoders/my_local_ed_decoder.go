package decoders

import (
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalEDDecoder struct {
	MyLocalDecoderInterface
}

func (d MyLocalEDDecoder) DecodeToEntity(record []string, source string, headers []string) models.Entity {

	// 0-id	1-name	2-country_id	3-province_id	4-ed_id		5-centroid	6-population
	decoder := MyLocalLocationDecoder{
		LocationId: record[0],
		Name:       record[1] + " Electoral District",
		Centroid:   record[5],
		Population: record[6],
		ParentId:   record[3],
		GeoSource:  "ed",
		Category:   "Electoral District",
		Attribute:  "electoral_districts",
		Source:     source,
	}
	decoder.ParentEntity = decoder.GetParentEntity()
	entity := decoder.MapToEntity()
	entity.AddCategory("LOCATION")
	decoder.AppendToParentEntity(entity)

	return entity
}
