package location

import (
	"GIG-Scripts/my_local/decoders"
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalDistrictDecoder struct {
	decoders.MyLocalDecoderInterface
}

func (d MyLocalDistrictDecoder) DecodeToEntity(record []string, source string, headers []string) models.Entity {
	// 0-id		1-district_id	2-province_id	3-name		4-centroid		5-population
	decoder := MyLocalLocationDecoder{
		LocationId: record[1],
		Name:       record[3] + " District",
		Centroid:   record[4],
		Population: record[5],
		GeoSource:  "district",
	}
	decoder.ParentId = record[2]
	decoder.Category = "District"
	decoder.Attribute = "districts"
	decoder.Source = source
	decoder.ParentEntity = decoder.GetParentEntity()
	entity := decoder.MapToEntity()
	entity.AddCategory("LOCATION")
	decoder.AppendToParentEntity("location_id", entity)

	return entity
}
