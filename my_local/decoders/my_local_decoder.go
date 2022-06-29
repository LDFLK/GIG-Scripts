package decoders

import (
	GIG_Scripts "GIG-Scripts"
	"GIG-Scripts/extended_models"
	"github.com/lsflk/gig-sdk/models"
	"log"
)

var entityMemo = map[string]models.Entity{}

type MyLocalDecoderInterface interface {
	DecodeToEntity(record []string, source string, headers []string) models.Entity
}

type MyLocalLocationDecoder struct {
	LocationId   string        `json:"locationId"`
	Name         string        `json:"name"`
	Centroid     string        `json:"centroid"`
	Population   string        `json:"population"`
	ParentId     string        `json:"parentId"`
	GeoSource    string        `json:"geoSource"`
	Category     string        `json:"category"`
	Attribute    string        `json:"attribute"`
	Source       string        `json:"source"`
	ParentEntity models.Entity `json:"parentEntity"`
}

func (d MyLocalLocationDecoder) GetParentEntity() models.Entity {

	// check if entity is already loaded in memory - this helps to avoid multiple request to fetch same entity
	if entity, found := entityMemo[d.ParentId]; found {
		log.Println("Parent entity found in memo.")
		return entity
	}

	// get the entity from server
	log.Println("Requesting parent entity from server")
	parentEntity, err := GIG_Scripts.GigClient.GetEntityByAttribute("attributes.location_id", d.ParentId)

	if err != nil {
		log.Fatal("error getting parent entity:", d.ParentId)
	}
	// save the entity to memory
	entityMemo[d.ParentId] = parentEntity
	return parentEntity
}

func (d MyLocalLocationDecoder) AppendToParentEntity(entity models.Entity) {
	//update parent entity
	payload := models.UpdateEntity{
		SearchAttribute: "attributes.location_id",
		SearchValue:     *new(models.Value).SetValueString(d.ParentId),
		Attribute:       d.Attribute,
		Value:           *new(models.Value).SetSource(d.Source).SetValueString(entity.GetTitle()),
	}
	if _, err := GIG_Scripts.GigClient.AppendToEntity(payload); err != nil {
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
