package decoders

import (
	GIG_Scripts "GIG-Scripts"
	"GIG-Scripts/extended_models"
	"github.com/lsflk/gig-sdk/models"
	"log"
)

type MyLocalLocationDecoder struct {
	MyLocalDecoderInterface
	MyLocalDecoder
	LocationId   string        `json:"locationId"`
	Name         string        `json:"name"`
	Centroid     string        `json:"centroid"`
	Population   string        `json:"population"`
	GeoSource    string        `json:"geoSource"`
	Category     string        `json:"category"`
	Attribute    string        `json:"attribute"`
	Source       string        `json:"source"`
	ParentEntity models.Entity `json:"parentEntity"`
}

func (d MyLocalLocationDecoder) AppendToParentEntity(entity models.Entity) {
	//update parent entity
	modifiedEntity := models.UpdateEntity{
		SearchAttribute: "attributes.location_id",
		SearchValue:     *new(models.Value).SetValueString(d.ParentId),
		Attribute:       d.Attribute,
		Value:           *new(models.Value).SetSource(d.Source).SetValueString(entity.GetTitle()),
	}
	if _, err := GIG_Scripts.GigClient.AppendToEntity(modifiedEntity); err != nil {
		log.Fatal("error updating parent entity:", err)
	}
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
