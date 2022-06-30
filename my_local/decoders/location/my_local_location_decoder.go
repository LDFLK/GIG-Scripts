package location

import (
	"GIG-Scripts/extended_models"
	"GIG-Scripts/my_local/decoders"
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalLocationDecoder struct {
	decoders.MyLocalDecoderInterface
	decoders.MyLocalDecoder
	LocationId string `json:"locationId"`
	Name       string `json:"name"`
	Centroid   string `json:"centroid"`
	Population string `json:"population"`
	GeoSource  string `json:"geoSource"`
}

func (d MyLocalLocationDecoder) MapToEntity() models.Entity {

	return *new(extended_models.Location).
		SetLocationId(d.LocationId, d.Source).
		SetName(d.Name, d.Source).
		SetCentroid(d.Centroid, d.Source).
		SetPopulation(d.Population, d.Source).
		SetParent(d.ParentEntity.GetTitle(), d.Source).
		SetGeoCoordinates("gig-data-master/geo/"+d.GeoSource+"/"+d.LocationId+".json", d.Source).
		AddCategory(d.Category).
		AddLink(models.Link{Title: d.ParentEntity.GetTitle()})

}
