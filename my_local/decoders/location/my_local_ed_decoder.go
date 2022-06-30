package location

import (
	"GIG-Scripts/my_local/decoders"
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalEDDecoder struct {
	decoders.MyLocalDecoderInterface
	MyLocalLocationDecoder
}

func (d MyLocalEDDecoder) DecodeToEntity(record []string, source string, headers []string) models.Entity {

	// 0-id	1-name	2-country_id	3-province_id	4-ed_id		5-centroid	6-population
	d.LocationId = record[0]
	d.Name = record[1] + " Electoral District"
	d.Centroid = record[5]
	d.Population = record[6]
	d.GeoSource = "ed"
	d.Category = "Electoral District"
	d.Attribute = "electoral_districts"
	d.Source = source
	d.ParentId = record[3]
	d.ParentAttribute = "location_id"
	d.ParentEntity = d.GetParentEntity()
	entity := d.MapToEntity()
	entity.AddCategory("LOCATION")
	d.AppendToParentEntity(entity)

	return entity
}
