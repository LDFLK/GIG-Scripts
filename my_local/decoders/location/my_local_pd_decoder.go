package location

import (
	"GIG-Scripts/my_local/decoders"
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalPDDecoder struct {
	decoders.MyLocalDecoderInterface
	MyLocalLocationDecoder
}

func (d MyLocalPDDecoder) DecodeToEntity(record []string, source string, headers []string) models.Entity {

	// 0-id		1-name	2-country_id	3-province_id	4-district_id	5-ed_id		6-pd_id		7-centroid		8-population
	d.LocationId = record[0]
	d.Name = record[1] + " Polling Division"
	d.Centroid = record[7]
	d.Population = record[8]
	d.GeoSource = "pd"
	d.Category = "Polling Division"
	d.Attribute = "polling_divisions"
	d.Source = source
	d.ParentId = record[5]
	d.ParentAttribute = "location_id"
	d.ParentEntity = d.GetParentEntity()
	entity := d.MapToEntity()
	entity.AddCategory("LOCATION")
	d.AppendToParentEntity(entity)

	return entity
}
