package decoders

import (
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalGNDDecoder struct {
	MyLocalDecoderInterface
}

func (d MyLocalGNDDecoder) DecodeToEntity(record []string, source string, headers []string) models.Entity {

	//0-id	1-gnd_id	2-gnd_num	3-name	4-province_id	5-district_id	6-dsd_id	7-centroid	8-population
	decoder := MyLocalLocationDecoder{
		LocationId: record[1],
		Name:       record[3] + " Grama Niladhari Division",
		Centroid:   record[7],
		Population: record[8],
		ParentId:   record[6],
		GeoSource:  "gnd",
		Category:   "Grama Niladhari Division",
		Attribute:  "grama_niladhari_divisions",
		Source:     source,
	}
	decoder.ParentEntity = decoder.GetParentEntity()
	entity := decoder.MapToEntity()
	entity.AddCategory("LOCATION")
	decoder.AppendToParentEntity(entity)

	return entity
}
