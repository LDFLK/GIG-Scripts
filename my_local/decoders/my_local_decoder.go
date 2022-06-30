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
	ParentId string `json:"parentId"`
}

func (d MyLocalDecoder) GetParentEntity() models.Entity {

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
