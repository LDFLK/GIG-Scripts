package location

import (
	"GIG-Scripts/my_local/decoders"
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalProvinceDecoder struct {
	decoders.MyLocalDecoderInterface
	MyLocalLocationDecoder
}

func (d MyLocalProvinceDecoder) DecodeToEntity(record []string, source string, headers []string) models.Entity {
	// 0-id		1-province_id	2-name	3-centroid	4-population
	d.LocationId = record[1]
	d.Name = record[2] + " Province"
	d.Centroid = record[3]
	d.Population = record[4]
	d.GeoSource = "province"
	d.Category = "Province"
	d.Attribute = "provinces"
	d.Source = source
	d.ParentId = "LK"
	d.ParentAttribute = "location_id"
	d.ParentEntity = d.GetParentEntity()
	entity := d.MapToEntity()
	entity.AddCategory("LOCATION")
	d.AppendToParentEntity(entity)

	return entity
}
