package decoders

import (
	GIG_Scripts "GIG-Scripts"
	"github.com/lsflk/gig-sdk/models"
	"log"
)

var entityMemo = map[string]models.Entity{}

type MyLocalDecoderInterface interface {
	DecodeToEntity(record []string, source string, headers []string) models.Entity
}

type MyLocalDecoder struct {
	MyLocalDecoderInterface
	ParentId        string
	Category        string
	Attribute       string
	Source          string
	ParentEntity    models.Entity
	ParentAttribute string
}

func (d MyLocalDecoder) GetParentEntity() models.Entity {

	// check if entity is already loaded in memory - this helps to avoid multiple request to fetch same entity
	if entity, found := entityMemo[d.ParentId]; found {
		log.Println("Parent entity found in memo.")
		return entity
	}

	// get the entity from server
	log.Println("Requesting parent entity from server")
	parentEntity, err := GIG_Scripts.GigClient.GetEntityByAttribute("attributes."+d.ParentAttribute, d.ParentId)

	if err != nil {
		log.Fatal("error getting parent entity:", d.ParentId)
	}
	// save the entity to memory
	entityMemo[d.ParentId] = parentEntity
	return parentEntity
}

func (d MyLocalDecoder) AppendToParentEntity(entity models.Entity) {
	//update parent entity
	modifiedEntity := models.UpdateEntity{
		SearchAttribute: "attributes." + d.ParentAttribute,
		SearchValue:     *new(models.Value).SetValueString(d.ParentId),
		Attribute:       d.Attribute,
		Value:           *new(models.Value).SetSource(d.Source).SetValueString(entity.GetTitle()),
	}
	if _, err := GIG_Scripts.GigClient.AppendToEntity(modifiedEntity); err != nil {
		log.Fatal("error updating parent entity:", err)
	}
}
