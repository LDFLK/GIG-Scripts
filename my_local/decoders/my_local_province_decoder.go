package decoders

import (
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalProvinceDecoder struct {
	MyLocalDecoderInterface
}

func (d MyLocalProvinceDecoder) DecodeToEntity(record []string, source string, headers []string) models.Entity {
	// 0-id		1-province_id	2-name	3-centroid	4-population
	decoder := MyLocalLocationDecoder{
		LocationId: record[1],
		Name:       record[2] + " Province",
		Centroid:   record[3],
		Population: record[4],
		ParentId:   "LK",
		GeoSource:  "province",
		Category:   "Province",
		Attribute:  "provinces",
		Source:     source,
	}
	decoder.ParentEntity = decoder.GetParentEntity()
	entity := decoder.MapToEntity()
	entity.AddCategory("LOCATION")
	decoder.AppendToParentEntity(entity)

	return entity
}
