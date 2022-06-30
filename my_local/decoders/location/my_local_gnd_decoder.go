package location

import (
	"GIG-Scripts/my_local/decoders"
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalGNDDecoder struct {
	decoders.MyLocalDecoderInterface
	MyLocalLocationDecoder
}

func (d MyLocalGNDDecoder) DecodeToEntity(record []string, source string, headers []string) models.Entity {

	//0-id	1-gnd_id	2-gnd_num	3-name	4-province_id	5-district_id	6-dsd_id	7-centroid	8-population
	d.LocationId = record[1]
	d.Name = record[3] + " Grama Niladhari Division"
	d.Centroid = record[7]
	d.Population = record[8]
	d.GeoSource = "gnd"
	d.Category = "Grama Niladhari Division"
	d.Attribute = "grama_niladhari_divisions"
	d.Source = source
	d.ParentId = record[6]
	d.ParentAttribute = "location_id"
	d.ParentEntity = d.GetParentEntity()
	entity := d.MapToEntity()
	entity.AddCategory("LOCATION")
	d.AppendToParentEntity(entity)

	return entity
}
