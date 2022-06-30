package location

import (
	"GIG-Scripts/my_local/decoders"
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalDistrictDecoder struct {
	decoders.MyLocalDecoderInterface
	MyLocalLocationDecoder
}

func (d MyLocalDistrictDecoder) DecodeToEntity(record []string, source string, headers []string) models.Entity {
	// 0-id		1-district_id	2-province_id	3-name		4-centroid		5-population
	d.LocationId = record[1]
	d.Name = record[3] + " District"
	d.Centroid = record[4]
	d.Population = record[5]
	d.GeoSource = "district"
	d.ParentId = record[2]
	d.Category = "District"
	d.Attribute = "districts"
	d.Source = source
	d.ParentAttribute = "location_id"
	d.ParentEntity = d.GetParentEntity()
	entity := d.MapToEntity()
	entity.AddCategory("LOCATION")
	d.AppendToParentEntity(entity)

	return entity
}
