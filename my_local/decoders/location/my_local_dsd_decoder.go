package location

import (
	"GIG-Scripts/my_local/decoders"
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalDSDDecoder struct {
	decoders.MyLocalDecoderInterface
	MyLocalLocationDecoder
}

func (d MyLocalDSDDecoder) DecodeToEntity(record []string, source string, headers []string) models.Entity {

	// 0-id	1-dsd_id	2-name	3-province_id	4-district_id	5-centroid	6-population
	d.LocationId = record[1]
	d.Name = record[2] + " Divisional Secretariats Division"
	d.Centroid = record[5]
	d.Population = record[6]
	d.GeoSource = "dsd"
	d.ParentId = record[4]
	d.Category = "Divisional Secretariats Division"
	d.Attribute = "divisional_secretariats_divisions"
	d.Source = source
	d.ParentAttribute = "location_id"
	d.ParentEntity = d.GetParentEntity()
	entity := d.MapToEntity()
	entity.AddCategory("LOCATION")
	d.AppendToParentEntity(entity)

	return entity
}
