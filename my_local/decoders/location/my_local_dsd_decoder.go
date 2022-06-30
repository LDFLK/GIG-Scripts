package location

import (
	"GIG-Scripts/my_local/decoders"
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalDSDDecoder struct {
	decoders.MyLocalDecoderInterface
}

func (d MyLocalDSDDecoder) DecodeToEntity(record []string, source string, headers []string) models.Entity {

	// 0-id	1-dsd_id	2-name	3-province_id	4-district_id	5-centroid	6-population
	decoder := MyLocalLocationDecoder{
		LocationId: record[1],
		Name:       record[2] + " Divisional Secretariats Division",
		Centroid:   record[5],
		Population: record[6],
		GeoSource:  "dsd",
		Category:   "Divisional Secretariats Division",
		Attribute:  "divisional_secretariats_divisions",
		Source:     source,
	}
	decoder.ParentId = record[4]
	decoder.Category = "Divisional Secretariats Division"
	decoder.Attribute = "divisional_secretariats_divisions"
	decoder.Source = source
	decoder.ParentEntity = decoder.GetParentEntity()
	entity := decoder.MapToEntity()
	entity.AddCategory("LOCATION")
	decoder.AppendToParentEntity(entity)

	return entity
}
